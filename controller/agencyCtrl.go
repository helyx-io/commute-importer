package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "fmt"
	"log"
    "sync"
    "time"
    "net/http"
    "encoding/json"
    "gopkg.in/redis.v2"
    "github.com/gorilla/mux"
    "github.com/fatih/stopwatch"
    "github.com/helyx-io/gtfs-playground/config"
    "github.com/helyx-io/gtfs-playground/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type Stop struct {
    Id int
    Name string
    Distance float64
    Lines []Line
}

type Line struct {
    Name string
    StopTimesFull []StopTimeFull
}

type StopTimeFull struct {
    StopId int
    StopName string
    StopDesc string
    StopLat string
    StopLon string
    LocationType int
    ArrivalTime string
    DepartureTime string
    StopSequence int
    DirectionId int
    RouteShortName string
    RouteType int
    RouteColor string
    RouteTextColor string
    TripId int
}

type FirstLastStopNamesByTripId struct {
    TripId string
    FirstStopName string
    LastStopName string
}

type JsonStop struct {
    Name string `json:"name"`
    Desc string `json:"desc"`
    Distance string `json:"distance"`
    LocationType int `json:"location_type"`
    GeoLocation JsonGeoLocation `json:"geo_location"`
    StopIds []string `json:"stop_ids"`
    Lines []JsonLine `json:"lines"`
}

type JsonGeoLocation struct {
    Lat string `json:"lat"`
    Lon string `json:"lon"`
}

type JsonLine struct {
    Name string `json:"lat"`
    StopTimes []string `json:"stop_times"`
    TripId int `json:"trip_id"`
    RouteColor string `json:"route_color"`
    RouteTextColor string `json:"route_text_color"`
    RouteType string `json:"route_type"`
    FirstStopName string `json:"first_stop_name"`
    LastStopName string `json:"last_stop_name"`
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
    client redis.Client
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Auth Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type AgencyController struct { }

func (ac *AgencyController) Init(r *mux.Router) {
    client = redis.NewTCPClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d", config.RedisInfos.Host, config.RedisInfos.Port),
        Password: "", // no password set
        DB:       0,  // use default DB
        PoolSize: 16,
    })

    // Init Router
	r.HandleFunc("/{agencyKey}/stops/{date}/nearest", ac.NearestStops)
}


func (ac *AgencyController) NearestStops(w http.ResponseWriter, r *http.Request) {

    defer utils.RecoverFromError(w)

    sw := stopwatch.Start(0)

    params := mux.Vars(r)

    agencyKey := params["agencyKey"]
    date := params["date"]

    lat := r.URL.Query().Get("lat")
    lon := r.URL.Query().Get("lon")
    distance := r.URL.Query().Get("distance")

    if len(distance) <= 0 {
        distance = "1000"
    }

    log.Printf("Agency Key: %s", agencyKey)
    log.Printf("Lat: %s", lat)
    log.Printf("Lon: %s", lon)
    log.Printf("Distance: %s", distance)
    log.Printf("Date: %s", date)


    stops := fetchStopsByDate(agencyKey, date, distance)

    tripIds := extractTripIds(stops)

    flStopNamesByTripId := fetchFirstAndLastStopNamesByTripIds(agencyKey, tripIds)

    // mergeFlStopNamesByTripIdWithStopLines(stops, flStopNamesByTripId)

    log.Printf("Resulting flStopNamesByTripId: %v", flStopNamesByTripId)

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- NearestStops. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")


    w.Header().Set("X-Response-Time", sw.ElapsedTime().String())
    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func fetchFirstAndLastStopNamesByTripIds(agencyKey string, tripIds []string) map[string][]string {

    flStopNamesByTripIdChan := make(chan FirstLastStopNamesByTripId)

    sem := make(chan bool, 64)

    go func() {
        for tripId := range tripIds {

            sem <- true

            go func(tripIds string) {
                defer func() { <-sem }()

                tripPayload := client.Get(fmt.Sprintf("/%s/s/t/fl/%s", agencyKey, tripId))

                tripFirstLast := make([]string, 2)

                err := json.Unmarshal([]byte(tripPayload), tripFirstLast)
                if err != nil {
                    log.Printf("Error: '%s' ...", err.Error())
                }

                flStopNamesByTripIdChan <- FirstLastStopNamesByTripId{tripId}
            }(tripId)
        }

        close(flStopNamesByTripIdChan)
    }()

    flStopNamesByTripIds := make(map[string][]string)

    for flStopNamesByTripId := range flStopNamesByTripIdChan {
        flStopNamesByTripIds[flStopNamesByTripId.TripId] = []string{flStopNamesByTripId.FirstStopName, flStopNamesByTripId.LastStopName}
    }

    return flStopNamesByTripIds
}

func extractTripIds(stops []Stop) []string {
    tripIdMap := make(map[string]bool)

    for stop := range stops {
       for line := range stop.Lines {
           if len(line) > 0 {
               tripIdMap[line.StopTimesFull.TripId] = true
           }
       }
    }

    tripIds := make([]int, 0, len(tripIdMap))
    for k := range tripIdMap {
        tripIds = append(tripIds, k)
    }

    return tripIds
}

func fetchStopsByDate(agencyKey, date, distance string) []Stop {

    query := fmt.Sprintf("select s.stop_id, s.stop_name, 111195 * st_distance(point(%s, %s), s.stop_geo) as stop_distance from gtfs_%s.stops s where 111195 * st_distance(point(%s, %s), s.stop_geo) < %s order by stop_distance asc", lat, lon, agencyKey, lat, lon, distance)

    //    log.Printf("Query: %s", query)
    rows, err := config.DB.Raw(query).Rows()
    defer rows.Close()

    utils.FailOnError(err, "Failed to execute query")

    sem := make(chan bool, 512)

    stopChan := make(chan StopTimeFull)

    go func() {
        for rows.Next() {
            id := 0
            name := ""
            distance := 0.0
            rows.Scan(&id, &name, &distance)

            stop := Stop{id, name, distance}
            //        log.Printf("Stop: %v", stop)

            sem <- true

            go func(stop Stop) {
                defer func() { <-sem }()

                stop.Lines = fetchLinesForDateAndStop(agencyKey, date, stop)

                stopChan <- stop
            }(stop)
        }

        close(stopChan)
    }()

    for i := 0; i < cap(sem); i++ {
        sem <- true
    }

    stops := make(chan Stop)

    for _, stop := range stopChan {
        stops = append(stops, stop)
    }

    return stops
}

func fetchLinesForDateAndStop(agencyKey, date string, stop Stop) []Line {
    stfs := fetchStopTimesFullForDateAndStop(agencyKey, date, stop)

    groupStopTimesFullByLine := func (stfs []StopTimeFull) []Line {

        stfByRouteShortName := make(map[string]StopTimeFull, 0)

        for _, stf := range stfs {
            if _, ok := stfByRouteShortName[stf.RouteShortName]; !ok {
                stfByRouteShortName[stf.RouteShortName] = make([]StopTimeFull, 0)
            }

            stfByRouteShortName[stf.RouteShortName] = append(stfByRouteShortName[stf.RouteShortName], stf)
        }

        lines := make([]Lines, len(stfByRouteShortName))

        for rsn, stf := range stfByRouteShortName {
            lines = append(lines, Line{rsn, stf})
        }

        return lines
    }

    return groupStopTimesFullByLine(stfs)
}

func fetchStopTimesFullForDateAndStop(agencyKey, date string, stop Stop) []StopTimeFull {
    day, _ := time.Parse("2006-01-02", date)
    dayOfWeek := day.Weekday().String()

    sem := make(chan bool, 512)

    stfChan := make(chan StopTimeFull)

    go func() {
        var wg sync.WaitGroup
        wg.Add(2)

        go fetchStopTimesFullForCalendar(agencyKey, stop, date, dayOfWeek, stfChan, wg)
        go fetchStopTimesFullForCalendarDates(agencyKey, stop, date, stfChan, wg)

        wg.Wait()

        close(stfChan)
    }()

    for i := 0; i < cap(sem); i++ {
        sem <- true
    }

    stfs := make([]StopTimeFull, 0)

    for stf := range stfChan {
        stfs = append(stfs, stf)
    }

    return stfs
}

func fetchStopTimesFullForCalendar(agencyKey string, stop Stop, date, dayOfWeek string, stfChan chan StopTimeFull, wg sync.WaitGroup) {
    queryCalendar := fmt.Sprintf(
    "select " +
    "stf.stop_id, " +
    "stf.stop_name, " +
    "stf.stop_desc, " +
    "stf.stop_lat, " +
    "stf.stop_lon, " +
    "stf.location_type, " +
    "stf.arrival_time, " +
    "stf.departure_time, " +
    "stf.stop_sequence, " +
    "stf.direction_id, " +
    "stf.route_short_name, " +
    "stf.route_type, " +
    "stf.route_color, " +
    "stf.route_text_color, " +
    "stf.trip_id " +
    "from gtfs_%s.stop_times_full stf inner join gtfs_%s.calendars c on stf.service_id=c.service_id where stf.stop_id=%d and c.start_date <= '%s' and c.end_date >= '%s' and %s=1", agencyKey, agencyKey, stop.Id, date, date, dayOfWeek)
    //                log.Printf("Query calendar: %s", queryCalendar)
    calendarRows, err := config.DB.Raw(queryCalendar).Rows()

    utils.FailOnError(err, "Calendar row fetch error")

    defer func() {
        calendarRows.Close()
        wg.Done()
    }()

    var stopId, locationType, stopSequence, directionId, routeType, tripId int
    var stopName, stopDesc, stopLat, stopLon, arrivalTime, departureTime, routeShortName, routeColor, routeTextColor string

    for calendarRows.Next() {
        calendarRows.Scan(
        &stopId, &stopName, &stopDesc, &stopLat, &stopLon, &locationType, &arrivalTime, &departureTime,
        &stopSequence, &directionId, &routeShortName, &routeType, &routeColor, &routeTextColor, &tripId,
        )

        log.Printf("StopId: %s", stopId)
        log.Printf("StopName: %s", stopName)

        stfChan <- StopTimeFull{stopId, stopName, stopDesc, stopLat, stopLon, locationType, arrivalTime, departureTime, stopSequence, directionId, routeShortName, routeType, routeColor, routeTextColor, tripId}
    }
}

func fetchStopTimesFullForCalendarDates(agencyKey string, stop Stop, date string, stfChan chan StopTimeFull, wg sync.WaitGroup) {
    queryCalendarDate := fmt.Sprintf(
        "select " +
        "stf.stop_id, " +
        "stf.stop_name, " +
        "stf.stop_desc, " +
        "stf.stop_lat, " +
        "stf.stop_lon, " +
        "stf.location_type, " +
        "stf.arrival_time, " +
        "stf.departure_time, " +
        "stf.stop_sequence, " +
        "stf.direction_id, " +
        "stf.route_short_name, " +
        "stf.route_type, " +
        "stf.route_color, " +
        "stf.route_text_color, " +
        "stf.trip_id " +
        "from gtfs_%s.stop_times_full stf inner join gtfs_%s.calendar_dates cd on stf.service_id=cd.service_id where stf.stop_id=%d and cd.date = '%s'", agencyKey, agencyKey, stop.Id, date)
    //                log.Printf("Query calendar dates : %s", queryCalendarDate)

    calendarDateRows, _ := config.DB.Raw(queryCalendarDate).Rows()

    defer func() {
        calendarDateRows.Close()
        wg.Done()
    }()

    var stopId, locationType, stopSequence, directionId, routeType, tripId int
    var stopName, stopDesc, stopLat, stopLon, arrivalTime, departureTime, routeShortName, routeColor, routeTextColor string

    for calendarDateRows.Next() {
        calendarDateRows.Scan(
        &stopId, &stopName, &stopDesc, &stopLat, &stopLon, &locationType, &arrivalTime, &departureTime,
        &stopSequence, &directionId, &routeShortName, &routeType, &routeColor, &routeTextColor, &tripId,
        )

        stfChan <- StopTimeFull{stopId, stopName, stopDesc, stopLat, stopLon, locationType, arrivalTime, departureTime, stopSequence, directionId, routeShortName, routeType, routeColor, routeTextColor, tripId}
    }

}