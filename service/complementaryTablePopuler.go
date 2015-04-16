package service

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "github.com/helyx-io/gtfs-importer/utils"
    "github.com/helyx-io/gtfs-importer/database"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structure
////////////////////////////////////////////////////////////////////////////////////////////////

type ComplementaryTablePopuler struct {
    driver *database.Driver
}

func NewComplementaryTablePopuler(driver *database.Driver) *ComplementaryTablePopuler {
    return &ComplementaryTablePopuler{driver}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func (ctp *ComplementaryTablePopuler) Populate(schema string, columnLengthsByFiles map[string]map[string]int) {

    columnLengthsByFilesTmp := make(map[string]interface{})
    for key, value := range columnLengthsByFiles {
        columnLengthsByFilesTmp[key] = value
    }

    err := ctp.driver.CreateTable(schema, "lines", columnLengthsByFilesTmp, true)
    utils.FailOnError(err, "Could not create 'lines' table")

    err = ctp.driver.CreateTable(schema, "line_stops", columnLengthsByFilesTmp, true)
    utils.FailOnError(err, "Could not create 'line_stops' table")

    err = ctp.driver.CreateTable(schema, "stations", columnLengthsByFilesTmp, true)
    utils.FailOnError(err, "Could not create 'stations' table")

    err = ctp.driver.CreateTable(schema, "station_stops", columnLengthsByFilesTmp, true)
    utils.FailOnError(err, "Could not create 'station_stops' table")

    err = ctp.driver.CreateTable(schema, "station_lines", columnLengthsByFilesTmp, true)
    utils.FailOnError(err, "Could not create 'station_lines' table")

    err = ctp.driver.CreateTable(schema, "route_stops", columnLengthsByFilesTmp, true)
    utils.FailOnError(err, "Could not create 'route_stops' table")


    ctp.driver.PopulateTable(schema, "lines")
    ctp.driver.PopulateTable(schema, "line_stops")
    ctp.driver.PopulateTable(schema, "stations")
    ctp.driver.PopulateTable(schema, "station_stops")
    ctp.driver.PopulateTable(schema, "station_lines")
    ctp.driver.PopulateTable(schema, "route_stops")


    err = ctp.driver.CreateIndex(schema, "lines", "line_name")
    utils.FailOnError(err, "Could not create index 'line_name' for table 'lines'")

    err = ctp.driver.CreateIndex(schema, "line_stops", "line_id")
    utils.FailOnError(err, "Could not create index 'line_id' for table 'line_stops'")

    err = ctp.driver.CreateIndex(schema, "line_stops", "stop_id")
    utils.FailOnError(err, "Could not create index 'stop_id' for table 'line_stops'")

    err = ctp.driver.CreateIndex(schema, "stations", "station_name")
    utils.FailOnError(err, "Could not create index 'station_name' for table 'stations'")

    err = ctp.driver.CreateSpatialIndex(schema, "stations", "station_geo")
    utils.FailOnError(err, "Could not create index 'station_geo' for table 'stations'")

    err = ctp.driver.CreateIndex(schema, "station_stops", "station_id")
    utils.FailOnError(err, "Could not create index 'station_id' for table 'station_stops'")

    err = ctp.driver.CreateIndex(schema, "station_stops", "stop_id")
    utils.FailOnError(err, "Could not create index 'stop_id' for table 'station_stops'")

    err = ctp.driver.CreateIndex(schema, "station_lines", "station_id")
    utils.FailOnError(err, "Could not create index 'station_id' for table 'station_lines'")

    err = ctp.driver.CreateIndex(schema, "station_lines", "line_id")
    utils.FailOnError(err, "Could not create index 'line_id' for table 'station_lines'")

    err = ctp.driver.CreateIndex(schema, "route_stops", "route_id")
    utils.FailOnError(err, "Could not create index 'route_id' for table 'route_stops'")

    err = ctp.driver.CreateIndex(schema, "route_stops", "stop_id")
    utils.FailOnError(err, "Could not create index 'stop_id' for table 'route_stops'")
}
