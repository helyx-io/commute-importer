package service

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "log"
    "fmt"
    "github.com/helyx-io/gtfs-importer/data"
    "github.com/helyx-io/gtfs-importer/utils"
    "github.com/helyx-io/gtfs-importer/database"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structure
////////////////////////////////////////////////////////////////////////////////////////////////

type AgenciesMetadataUpdater struct {
    driver *database.Driver
}

func NewAgenciesMetadataUpdater(driver *database.Driver) *AgenciesMetadataUpdater {
    return &AgenciesMetadataUpdater{driver}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func (amu *AgenciesMetadataUpdater) UpdateAgenciesMetaData(agencyKey, schema string) error {
    dbSql, err := amu.driver.Open()

    if err != nil {
        panic(err.Error())
    }

    defer dbSql.Close()


    log.Printf("Updating agency zone for schema: %s", schema)

    selectFilePath := fmt.Sprintf("resources/ddl/%s/select-agency-zone.sql", amu.driver.ConnectInfos.Dialect)
    ddlSelect, err := data.Asset(selectFilePath)
    utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for fetching agency zone: '%s'", selectFilePath, schema))

    selectStmt := fmt.Sprintf(string(ddlSelect), schema)

    log.Printf("Fetch agency zone infos for schema: '%s': '%s'", schema, selectStmt)

    row := amu.driver.Raw(selectStmt).Row()

    var min_stop_lat float64;
    var max_stop_lat float64;
    var min_stop_lon float64;
    var max_stop_lon float64;

    row.Scan(&min_stop_lat, &max_stop_lat, &min_stop_lon, &max_stop_lon)

    log.Printf(" - Min stop lat: %f", min_stop_lat)
    log.Printf(" - Max stop lat: %f", max_stop_lat)
    log.Printf(" - Min stop lon: %f", min_stop_lon)
    log.Printf(" - Max stop lon: %f", max_stop_lon)


    log.Printf("Updating agency zone for schema: %s", schema)


    updateFilePath := fmt.Sprintf("resources/ddl/%s/update-agency-zone.sql", amu.driver.ConnectInfos.Dialect)
    ddlUpdate, err := data.Asset(updateFilePath)
    utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for updating agency zone: '%s'", updateFilePath, schema))

    updateStmt := fmt.Sprintf(string(ddlUpdate), schema)

    updateValueArgs := []interface{}{ min_stop_lat, max_stop_lat, min_stop_lon, max_stop_lon }

    log.Printf("Fetch agency zone infos for schema: '%s': '%s' - Args: %v", schema, updateStmt, updateValueArgs)

    _, err = dbSql.Exec(updateStmt, updateValueArgs...)

    if err != nil {
        log.Println(fmt.Println("Failed on Error: '%v'", err))
        return err
    }


    updateGtfsFilePath := fmt.Sprintf("resources/ddl/%s/update-gtfs-agency-zone.sql", amu.driver.ConnectInfos.Dialect)
    ddlUpdateGtfs, err := data.Asset(updateGtfsFilePath)
    utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for updating agency zone: '%s'", updateGtfsFilePath, agencyKey))

    updateGtfsStmt := string(ddlUpdateGtfs)

    updateGtfsValueArgs := []interface{}{ min_stop_lat, max_stop_lat, min_stop_lon, max_stop_lon, agencyKey }

    log.Printf("Fetch agency zone infos for schema: '%s': '%s' - Args: %v", schema, updateGtfsStmt, updateGtfsValueArgs)

    _, err = dbSql.Exec(updateGtfsStmt, updateGtfsValueArgs...)

    if err != nil {
        log.Printf("Failed on Error: '%v'", err)
    }

    return err
}
