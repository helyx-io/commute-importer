package service

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "log"
    "fmt"
    "regexp"
    "strings"
    "github.com/helyx-io/commute-importer/data"
    "github.com/helyx-io/commute-importer/database"
    "github.com/helyx-io/commute-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structure
////////////////////////////////////////////////////////////////////////////////////////////////

type StopTimesFullImporter struct {
    driver *database.Driver
}

func NewStopTimesFullImporter(driver *database.Driver) *StopTimesFullImporter {
    return &StopTimesFullImporter{driver}
}

type Lines []Line

type Line struct {
    Id int `gorm:"column:line_id"`
    Name  string `gorm:"column:line_name"`
}

type InsertLineResult struct {
    Line Line
    Error error
}

type CreateIndexResult struct {
    Index string
    Error error
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////


func (stfi *StopTimesFullImporter) ImportStopTimesFull(schema string, columnLengthsByFiles map[string]map[string]int) {
    //	config.DB.LogMode(true)

    columnLengthsByFilesTmp := make(map[string]interface{})
    for key, value := range columnLengthsByFiles {
        columnLengthsByFilesTmp[key] = value
    }


    tableName := "stop_times_full"
    stfi.driver.CreateTable(schema, tableName, columnLengthsByFilesTmp, true)

    filePath := fmt.Sprintf("resources/ddl/%s/insert-%s.sql", stfi.driver.ConnectInfos.Dialect, tableName)

    log.Printf("Inserting data into table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    ddlBytes, err := data.Asset(filePath)
    ddl := string(ddlBytes)
    utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table '%s.%s'", filePath, schema, tableName))

    var lines Lines

    fullTableName := strings.ToLower(fmt.Sprintf("%s.%s", schema, "lines"))
    err = stfi.driver.DB.Table(fullTableName).Find(&lines).Error
    utils.FailOnError(err, fmt.Sprintf("Could get lines from table '%s'", fullTableName))

    insertLineDoneChan := make(chan InsertLineResult, 8)

    go func() {
        for _, line := range lines {
            stfi.insertForLine(schema, tableName, ddl, line, insertLineDoneChan)
            //		log.Printf("--- Inserting data for line: %s [%s, %s, %s]", line.Name, schema, tableName, ddl)
        }
    }()

    doneCount := 0
    for insertLineResult := range insertLineDoneChan {
        doneCount += 1
        if insertLineResult.Error != nil {
            log.Printf("Received event on done chan with error: '%s'", insertLineResult.Error)
        } else {
            log.Printf("Line count: %v - done count: %v", len(lines), doneCount)
            if len(lines) == doneCount {
                log.Printf("Closing done chan.")
                close(insertLineDoneChan)
            } else {
                log.Printf("Received event on done chan for line '%s'", insertLineResult.Line.Name)
            }
        }
    }


    createIndexDoneChan := make(chan CreateIndexResult, 8)
    indexes := []string{"service_id", "stop_id", "trip_id"/*, "route_id"*/}

    go func() {
        for _, index := range indexes {
            err := stfi.driver.CreateIndex(schema, tableName, index)
            createIndexDoneChan <- CreateIndexResult{index, err}
        }
    }()

    doneCount = 0
    for createIndexResult := range createIndexDoneChan {
        if createIndexResult.Error != nil {
            log.Printf("[CREATE_INDEX] Received event on done chan for index :%s with error: '%s'", createIndexResult.Index, createIndexResult.Error)
        } else {
            doneCount += 1
            if len(indexes) == doneCount {
                log.Printf("[CREATE_INDEX] Closing done chan")
                close(createIndexDoneChan)
            } else {
                log.Printf("[CREATE_INDEX] Received event on done chan for index '%s'", createIndexResult.Index)
            }
        }
    }
}


func (stfi *StopTimesFullImporter) insertForLine(schema string, tableName string, ddl string, line Line, doneChan chan InsertLineResult) {
    log.Printf("--- Inserting data for line: '%s'", line.Name)

    insertStmt := regexp.MustCompile("%s").ReplaceAllString(ddl, schema)
    insertStmt = regexp.MustCompile("%v").ReplaceAllString(insertStmt, line.Name)

    //		log.Printf("Insert statement: %s", insertStmt)
    err := stfi.driver.ExecQuery(insertStmt)
    doneChan <- InsertLineResult{line, err}
}
