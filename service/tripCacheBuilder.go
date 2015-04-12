package service


////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "fmt"
    "log"
    "encoding/json"
    "gopkg.in/redis.v2"
    "github.com/jinzhu/gorm"
    "github.com/helyx-io/gtfs-importer/utils"
    "github.com/helyx-io/gtfs-importer/data"
    "github.com/helyx-io/gtfs-importer/config"
    "github.com/fatih/stopwatch"
    "sync"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type Task string

type StringKeyValue struct {
    Key string
    Value string
}

type StopTime struct {
    Arrival_time string     `json:"a"`
    Departure_time string   `json:"d"`
    Stop_sequence int       `json:"s"`
    Stop_name string        `json:"n"`
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Trip Cache Builder
////////////////////////////////////////////////////////////////////////////////////////////////

func BuildTripCache(db *gorm.DB, connectInfos *config.DBConnectInfos, redis *redis.Client, schema string) {

    sw := stopwatch.Start(0)

    selectTripStopTimesDdl, _ := data.Asset(fmt.Sprintf("resources/ddl/%s/select_trip_stop_times.sql", connectInfos.Dialect))
    selectTripStopTimesQuery := string(selectTripStopTimesDdl)

    keyValues := make(chan StringKeyValue, 32)

    go func() {

        ddl, err := data.Asset(fmt.Sprintf("resources/ddl/%s/select_trips.sql", connectInfos.Dialect))
        utils.FailOnError(err, fmt.Sprintf("Build trips cache for agency key: %s", schema))

        tripIdsQuery := fmt.Sprintf(string(ddl), schema)
        log.Printf("Query: %s", tripIdsQuery)
        rows, _ := db.Raw(tripIdsQuery).Rows()
        defer rows.Close()

        tasks := make(chan Task, 128)
        quit := make(chan bool)
        var wg sync.WaitGroup

        // spawn 8 workers
        for i := 0; i < 8; i++ {
            wg.Add(1)
            go worker(db, selectTripStopTimesQuery, schema, keyValues, tasks, quit, &wg)
        }

        tripId := ""
        for rows.Next() {
            rows.Scan(&tripId)
            tasks <- Task(tripId)
        }

        // end of tasks. the workers should quit afterwards
        close(tasks)
        // use "close(quit)", if you do not want to wait for the remaining tasks

        // wait for all workers to shut down properly
        wg.Wait()

        close(keyValues)
    }()

    flushCount := 0

    entries := make([]string, 0)

    for keyValue := range keyValues {

        entries = append(entries, keyValue.Key, keyValue.Value)

        if len(entries) > 2048 { // 2 entries ... => 1024
            flushCount += 1

            statusCmd := redis.MSet(entries...);
            if statusCmd.Err() != nil {
                log.Printf("Error: '%s' ...", statusCmd.Err().Error())
            }

            log.Printf("Flush Count: %d", flushCount)

            entries = make([]string, 0)
        }
    }

    if len(entries) > 0 {
        flushCount += 1
        log.Printf("Flush Count: %d", flushCount)

        statusCmd := redis.MSet(entries...);
        if statusCmd.Err() != nil {
            log.Printf("Error: '%s' ...", statusCmd.Err().Error())
        }

        entries = make([]string, 0)
    }

    log.Printf(fmt.Sprintf("Query trip ids - Elapsed Time: %v", sw.ElapsedTime()))
}


func worker(db *gorm.DB, query, schema string, keyValues chan StringKeyValue, tripIds <-chan Task, quit <-chan bool, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case tripId, ok := <-tripIds:
            if !ok {
                return
            }

            processTripId(db, query, schema, string(tripId), keyValues)
        case <-quit:
            return
        }
    }
}


func processTripId(db *gorm.DB, query, schema, tripId string, keyValues chan StringKeyValue) {

    stopTimesQuery := fmt.Sprintf(query, schema, schema, tripId)

    stopTimeRows, err := db.Raw(stopTimesQuery).Rows()
    defer stopTimeRows.Close()

    if err != nil {
        panic(err.Error())
    }

    stopTimes := make([]StopTime, 0)

    for stopTimeRows.Next() {
        var arrival_time, departure_time, stop_name string
        var stop_sequence int
        stopTimeRows.Scan(&arrival_time, &departure_time, &stop_sequence, &stop_name)
        stopTimes = append(stopTimes, StopTime{arrival_time, departure_time, stop_sequence, stop_name});
    }

    stopTimesLength := len(stopTimes)

    if stopTimesLength >= 2 {
        tripFirstLast := []string{ stopTimes[0].Stop_name, stopTimes[len(stopTimes) - 1].Stop_name }

        bytes, err := json.Marshal(tripFirstLast)
        if err != nil {
            log.Printf("Error: '%s' ...", err.Error())
        }

        cacheKey := fmt.Sprintf("/%s/t/st/fl/%s", schema, tripId)
        tripFirstLastStr := string(bytes)

        keyValues <- StringKeyValue{cacheKey, tripFirstLastStr}
    }

}


