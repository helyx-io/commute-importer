package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "fmt"
	"log"
    "sync"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/fatih/stopwatch"
    "github.com/helyx-io/gtfs-playground/config"
    "github.com/helyx-io/gtfs-playground/utils"
    "time"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type Stop struct {
    Id int
    Name string
    Distance float64
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


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (

)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Auth Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type AgencyController struct { }

func (ac *AgencyController) Init(r *mux.Router) {

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

    day, _ := time.Parse("2006-01-02", date)
    dayOfWeek := day.Weekday().String()
//    dayOfWeek := daysOfWeek(moment(date, 'YYYY-MM-DD').format('E'));


    if len(distance) <= 0 {
        distance = "1000"
    }

    log.Printf("Agency Key: %s", agencyKey)
    log.Printf("Lat: %s", lat)
    log.Printf("Lon: %s", lon)
    log.Printf("Distance: %s", distance)
    log.Printf("Date: %s", date)
    log.Printf("Day of Week: %s", dayOfWeek)


    query := fmt.Sprintf("select s.stop_id, s.stop_name, 111195 * st_distance(point(%s, %s), s.stop_geo) as stop_distance from gtfs_%s.stops s where 111195 * st_distance(point(%s, %s), s.stop_geo) < %s order by stop_distance asc", lat, lon, agencyKey, lat, lon, distance)

//    log.Printf("Query: %s", query)
    rows, err := config.DB.Raw(query).Rows()
    defer rows.Close()

    utils.FailOnError(err, "Failed to execute query")

    sem := make(chan bool, 512)

    for rows.Next() {
        id := 0
        name := ""
        distance := 0.0
        rows.Scan(&id, &name, &distance)

        stop := Stop{id, name, distance}
        log.Printf("Stop: %v", stop)

        sem <- true

        go func(stop Stop) {

            defer func() { <-sem }()

            var wg sync.WaitGroup
            wg.Add(2)

            go func() {
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

                stopTimesCalendars := make([]StopTimeFull, 0)

                for calendarRows.Next() {
                    calendarRows.Scan(
                        &stopId, &stopName, &stopDesc, &stopLat, &stopLon, &locationType, &arrivalTime, &departureTime,
                        &stopSequence, &directionId, &routeShortName, &routeType, &routeColor, &routeTextColor, &tripId,
                    )


                    log.Printf("StopId: %s", stopId)
                    log.Printf("StopName: %s", stopName)


                    stopTimeFull := StopTimeFull{stopId, stopName, stopDesc, stopLat, stopLon, locationType, arrivalTime, departureTime, stopSequence, directionId, routeShortName, routeType, routeColor, routeTextColor, tripId}

                    stopTimesCalendars = append(stopTimesCalendars, stopTimeFull)
                }

//                log.Printf("StopTimesFull for calendars: %d", len(stopTimesCalendars))
            }()

            go func() {
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

                utils.FailOnError(err, "Calendar dates row fetch error")

                defer func() {
                    calendarDateRows.Close()
                    wg.Done()
                }()

                var stopId, locationType, stopSequence, directionId, routeType, tripId int
                var stopName, stopDesc, stopLat, stopLon, arrivalTime, departureTime, routeShortName, routeColor, routeTextColor string

                stopTimesCalendarDates := make([]StopTimeFull, 0)

                for calendarDateRows.Next() {
                    calendarDateRows.Scan(
                        &stopId, &stopName, &stopDesc, &stopLat, &stopLon, &locationType, &arrivalTime, &departureTime,
                        &stopSequence, &directionId, &routeShortName, &routeType, &routeColor, &routeTextColor, &tripId,
                    )

                    stopTimeFull := StopTimeFull{stopId, stopName, stopDesc, stopLat, stopLon, locationType, arrivalTime, departureTime, stopSequence, directionId, routeShortName, routeType, routeColor, routeTextColor, tripId}

                    stopTimesCalendarDates = append(stopTimesCalendarDates, stopTimeFull)
                }

//                log.Printf("StopTimesFull for calednarDates: %d", len(stopTimesCalendarDates))
            }()

            wg.Wait()

        }(stop)

    }

    for i := 0; i < cap(sem); i++ {
        sem <- true
    }

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- NearestStops. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")


    w.Header().Set("X-Response-Time", sw.ElapsedTime().String())
    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}
