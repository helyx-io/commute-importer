package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "os"
	"fmt"
	"log"
	"sort"
    "path"
	"time"
	"regexp"
    "encoding/csv"
	"net/http"
	"database/sql"

    "gopkg.in/redis.v2"
    "github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
    "github.com/helyx-io/gtfs-importer/models"
    "github.com/helyx-io/gtfs-importer/config"
	"github.com/helyx-io/gtfs-importer/utils"
	"github.com/helyx-io/gtfs-importer/service"
	"github.com/helyx-io/gtfs-importer/database"
	"github.com/helyx-io/gtfs-importer/data"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	repositoryByFilenameMap map[string]database.GTFSCreatedModelRepository
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

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
/// Helper functions
////////////////////////////////////////////////////////////////////////////////////////////////

func initRepositoryMap(gtfs database.GTFSRepository) {
	repositoryByFilenameMap = make(map[string]database.GTFSCreatedModelRepository)

	repositoryByFilenameMap["agency.txt"] = gtfs.Agencies()
	repositoryByFilenameMap["calendar_dates.txt"] = gtfs.CalendarDates()
	repositoryByFilenameMap["calendar.txt"] = gtfs.Calendars()
	repositoryByFilenameMap["routes.txt"] = gtfs.Routes()
	repositoryByFilenameMap["stops.txt"] = gtfs.Stops()
	repositoryByFilenameMap["stop_times.txt"] = gtfs.StopTimes()
	repositoryByFilenameMap["transfers.txt"] = gtfs.Transfers()
	repositoryByFilenameMap["trips.txt"] = gtfs.Trips()

}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Import Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type ImportController struct {
    db *gorm.DB
    gtfs database.GTFSRepository
    redis *redis.Client
    connectInfos *config.DBConnectInfos
    dataResources map[string]string
    tmpDir string
}

func (ic *ImportController) Init(r *mux.Router, dataResources map[string]string, tmpDir string, redis *redis.Client, db *gorm.DB, connectInfos *config.DBConnectInfos, gtfs database.GTFSRepository) {

    ic.gtfs = gtfs
    ic.db = db
    ic.connectInfos = connectInfos
    ic.redis = redis
    ic.tmpDir = tmpDir
    ic.dataResources = dataResources


	// Init Router
    r.HandleFunc("/{key}", ic.Import)
    r.HandleFunc("/{key}/caches/trips", ic.BuildTripCache)

	// Init Repository Map
	initRepositoryMap(gtfs)
}

func (ic *ImportController) rewriteCsvFiles(schema, outFolderName string) error {

    agencyIndexes, err := ic.getIndexes(schema, "agency.txt", 0)
    serviceIndexes, err := ic.getIndexes(schema, "trips.txt", 1)
    tripIndexes, err := ic.getIndexes(schema, "trips.txt", 2)
    stopIndexes, err := ic.getIndexes(schema, "stops.txt", 0)
    routeIndexes, err := ic.getIndexes(schema, "routes.txt", 0)

    ic.writeIndexes(schema, "routes.indexes.txt", outFolderName, routeIndexes);
    ic.writeIndexes(schema, "trip.indexes.txt", outFolderName, tripIndexes);
    ic.writeIndexes(schema, "stop.indexes.txt", outFolderName, stopIndexes);

    folderFilename := ic.tmpDir + "/" + schema
    outFolderFilename := path.Join(folderFilename, outFolderName)

    if os.MkdirAll(outFolderFilename, 0755) != nil {
        panic("Unable to create directory for tagfile!")
    }

    indexes := map[int](map[string]string){ 0: stopIndexes }
    ic.rewriteCsvFile(schema, "stops.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: tripIndexes, 3: stopIndexes }
    ic.rewriteCsvFile(schema, "stop_times.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: routeIndexes, 1: agencyIndexes }
    ic.rewriteCsvFile(schema, "routes.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: agencyIndexes }
    ic.rewriteCsvFile(schema, "agency.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: routeIndexes, 1: serviceIndexes, 2: tripIndexes }
    ic.rewriteCsvFile(schema, "trips.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: serviceIndexes }
    ic.rewriteCsvFile(schema, "calendar.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: serviceIndexes }
    ic.rewriteCsvFile(schema, "calendar_dates.txt", outFolderName, indexes)

    return err
}


func (ic *ImportController) writeIndexes(schema, filename, outFolderName string, indexes map[string]string) error {

    folderName := path.Join(ic.tmpDir, schema)

    outFile, err := os.Create(path.Join(folderName, path.Join(outFolderName, filename)))
    if err != nil {
        log.Printf("Error: '%v'", err.Error())
        return err
    }

    log.Printf("Writing to file: '%v'", outFile.Name())

    writer := csv.NewWriter(outFile)

    for key, value := range indexes {
        writer.Write([]string{key, value})
    }

    writer.Flush()
    err = outFile.Close()

    return err
}


func (ic *ImportController) rewriteCsvFile(schema, filename, outFolderName string, indexes map[int](map[string]string)) error {

    folderName := path.Join(ic.tmpDir, schema)
    filePath := path.Join(folderName, filename)

    gtfsFile := models.GTFSFile{filePath}
    
    outFile, err := os.Create(path.Join(folderName, path.Join(outFolderName, filename)))
    if err != nil {
        log.Printf("Error: %v", err.Error())
        return err
    }

    log.Printf("Writing to file: '%v'", outFile.Name())

    writer := csv.NewWriter(outFile)
    headers, err := utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
    log.Printf("headers: '%v'", headers)

    if err != nil {
        log.Printf("Error: '%v'", err.Error())
        return err
    }

    writer.Write(headers)

    resultChan := make(chan [][]string)

    go func() {

//        sem := make(chan bool, 8)

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
//            sem <- true
//            go func() {
//                defer func() { <-sem }()
                records, _ := models.ParseCsvAsStringArrays(lines)

                for _, record := range *records {
                    for i, _ := range indexes {
                        record[i] = indexes[i][record[i]]
                    }
                }

                resultChan <- *records
//            }()
        }

//        for i := 0; i < cap(sem); i++ {
//            sem <- true
//        }

        close(resultChan)

    }()

    offset := 0

    for results := range resultChan {
        offset++
        log.Printf("[%s][%d] Records write", filePath, offset)

        writer.WriteAll(results)
    }

    writer.Flush()
    err = outFile.Close()

    return err
}

func (ic *ImportController) getIndexes(schema, filename string, index int) (map[string]string, error) {

    folderName := path.Join(ic.tmpDir, schema)
    filePath := path.Join(folderName, filename)

    gtfsFile := models.GTFSFile{filePath}

    resultChan := make(chan []string)

    go func() {

//        sem := make(chan bool, 8)

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
//            sem <- true
//            go func() {
//                defer func() { <-sem }()
                records, _ := models.ParseCsvAsStringArrays(lines)

                keys := []string{}
                for _, record := range *records {
                    keys = append(keys, record[index])
                }

                resultChan <- keys
//            }()
        }

//        for i := 0; i < cap(sem); i++ {
//            sem <- true
//        }

        close(resultChan)

    }()

    offset := 0
    result := []string{}
    indexes := map[string]string{}
    increment := 0

    for results := range resultChan {
        offset++
        log.Printf("[%s][%d] Records read", filePath, offset)
        for _, key := range results {
            if _, ok := indexes[key]; !ok {
                increment++
                result = append(result, key)
                index :=  fmt.Sprintf("%d", increment)
                indexes[key] = index
            }
        }
    }

    return indexes, nil
}


func (ic *ImportController) importPostProcess(schema string) {

    log.Printf("ic.db: %v", ic.db)
    log.Printf("ic.connectInfos: %v", ic.connectInfos)

	err := database.CreateTable(ic.db, ic.connectInfos, schema, "lines", true)
    utils.FailOnError(err, "Could not create 'lines' table")

    err = database.CreateTable(ic.db, ic.connectInfos, schema, "line_stops", true)
    utils.FailOnError(err, "Could not create 'line_stops' table")

    err = database.CreateTable(ic.db, ic.connectInfos, schema, "stations", true)
    utils.FailOnError(err, "Could not create 'stations' table")

    err = database.CreateTable(ic.db, ic.connectInfos, schema, "station_stops", true)
    utils.FailOnError(err, "Could not create 'station_stops' table")

    err = database.CreateTable(ic.db, ic.connectInfos, schema, "station_lines", true)
    utils.FailOnError(err, "Could not create 'station_lines' table")

    err = database.CreateTable(ic.db, ic.connectInfos, schema, "route_stops", true)
    utils.FailOnError(err, "Could not create 'route_stops' table")


    ic.populateTable(schema, "lines")
    ic.populateTable(schema, "line_stops")
    ic.populateTable(schema, "stations")
    ic.populateTable(schema, "station_stops")
    ic.populateTable(schema, "station_lines")
    ic.populateTable(schema, "route_stops")


    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "lines", "line_name")
    utils.FailOnError(err, "Could not create index 'line_name' for table 'lines'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "line_stops", "line_id")
    utils.FailOnError(err, "Could not create index 'line_id' for table 'line_stops'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "line_stops", "stop_id")
    utils.FailOnError(err, "Could not create index 'stop_id' for table 'line_stops'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "stations", "station_name")
    utils.FailOnError(err, "Could not create index 'station_name' for table 'stations'")

    err = database.CreateSpatialIndex(ic.db, ic.connectInfos, schema, "stations", "station_geo")
    utils.FailOnError(err, "Could not create index 'station_geo' for table 'stations'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "station_stops", "station_id")
    utils.FailOnError(err, "Could not create index 'station_id' for table 'station_stops'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "station_stops", "stop_id")
    utils.FailOnError(err, "Could not create index 'stop_id' for table 'station_stops'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "station_lines", "station_id")
    utils.FailOnError(err, "Could not create index 'station_id' for table 'station_lines'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "station_lines", "line_id")
    utils.FailOnError(err, "Could not create index 'line_id' for table 'station_lines'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "route_stops", "route_id")
    utils.FailOnError(err, "Could not create index 'route_id' for table 'route_stops'")

    err = database.CreateIndex(ic.db, ic.connectInfos, schema, "route_stops", "stop_id")
    utils.FailOnError(err, "Could not create index 'stop_id' for table 'route_stops'")
}


func (ic *ImportController) importStopTimesFull(schema string) {
//	config.DB.LogMode(true)

	tableName := "stop_times_full"
    database.CreateTable(ic.db, ic.connectInfos, schema, tableName, true)

	filePath := fmt.Sprintf("resources/ddl/%s/insert-%s.sql", ic.connectInfos.Dialect, tableName)

	log.Printf("Inserting data into table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

	ddlBytes, err := data.Asset(filePath)
	ddl := string(ddlBytes)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table '%s.%s'", filePath, schema, tableName))

	var lines Lines

	fullTableName := fmt.Sprintf("%s.%s", schema, "lines")
	err = ic.db.Table(fullTableName).Find(&lines).Error
	utils.FailOnError(err, fmt.Sprintf("Could get lines from table '%s'", fullTableName))

	insertLineDoneChan := make(chan InsertLineResult, 8)

    go func() {
        for _, line := range lines {
            ic.insertForLine(schema, tableName, ddl, line, insertLineDoneChan)
            //		log.Printf("--- Inserting data for line: %s [%s, %s, %s]", line.Name, schema, tableName, ddl)
        }
    }()

	doneCount := 0
	for insertLineResult := range insertLineDoneChan {
		if insertLineResult.Error != nil {
			log.Printf("Received event on done chan with error: '%s'", insertLineResult.Error)
		} else {
			doneCount += 1
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
            ic.createIndex(schema, tableName, index, createIndexDoneChan)
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


func (ic *ImportController) updateAgenciesMetaData(agencyKey, schema string) error {
	dbSql, err := sql.Open(ic.connectInfos.Dialect, ic.connectInfos.URL)

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()


	log.Printf("Updating agency zone for schema: %s", schema)

	selectFilePath := fmt.Sprintf("resources/ddl/%s/select-agency-zone.sql", ic.connectInfos.Dialect)
	ddlSelect, err := data.Asset(selectFilePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for fetching agency zone: '%s'", selectFilePath, schema))

	selectStmt := fmt.Sprintf(string(ddlSelect), schema)

	log.Printf("Fetch agency zone infos for schema: '%s': '%s'", schema, selectStmt)

	row := ic.db.Raw(selectStmt).Row()

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


	updateFilePath := fmt.Sprintf("resources/ddl/%s/update-agency-zone.sql", ic.connectInfos.Dialect)
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


	updateGtfsFilePath := fmt.Sprintf("resources/ddl/%s/update-gtfs-agency-zone.sql", ic.connectInfos.Dialect)
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


func (ic *ImportController) insertForLine(schema string, tableName string, ddl string, line Line, doneChan chan InsertLineResult) {
	log.Printf("--- Inserting data for line: '%s'", line.Name)

	insertStmt := regexp.MustCompile("%s").ReplaceAllString(ddl, schema)
	insertStmt = regexp.MustCompile("%v").ReplaceAllString(insertStmt, line.Name)

	//		log.Printf("Insert statement: %s", insertStmt)
	err := ic.db.Exec(insertStmt).Error
	doneChan <- InsertLineResult{line, err}
}


func (ic *ImportController) Import(w http.ResponseWriter, r *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
	keyParam := params["key"]

	log.Printf("Importing agencies for Key: '%s' ...", keyParam)

	w.Header().Set("Content-Type", "text/html")

	folderFilename := ic.tmpDir + "/" + keyParam
	url := ic.dataResources[keyParam]

	zipFilename := ic.tmpDir + "/" + keyParam + "-" + time.Now().Format("20060102-150405") + ".zip"

	utils.DownloadFile(url, zipFilename)
	utils.UnzipArchive(zipFilename, folderFilename)

    outFolderFilename := path.Join(folderFilename, "out")

    if os.MkdirAll(outFolderFilename, 0755) != nil {
        panic("Unable to create directory for tagfile!")
    }

    ic.rewriteCsvFiles(keyParam, "out")

	fis := utils.ReadDirectoryFileInfos(outFolderFilename)
	sort.Sort(utils.FileInfosBySize(fis))

	err := ic.gtfs.CreateSchema(keyParam)
	utils.FailOnError(err, fmt.Sprintf("Could not create schema for key: '%s'", keyParam))

	for _, fi := range fis {
		if fi.Mode().IsRegular() {
			gtfsModelRepository := repositoryByFilenameMap[fi.Name()]

			if gtfsModelRepository == nil {
				log.Printf("Filename '%v' is not available in map", fi.Name())
				continue;
			}

			gaf := service.NewGTFSArchiveFile(fi)

			err := gaf.ImportGTFSArchiveFileWithTableCreation(keyParam, outFolderFilename, gtfsModelRepository, 512 * 1000)
			utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive with table creation for key: '%s'", keyParam, fi.Name()))

			if fi.Name() == "agency.txt" {
				log.Println("Importing agencies in GTFS agencies table ...")

				gtfsAgencyModelRepository := ic.gtfs.GtfsAgencies()
				gaf := service.NewGTFSArchiveFile(fi)

				err:= gaf.ImportGTFSArchiveFileWithoutTableCreation(keyParam, outFolderFilename, gtfsAgencyModelRepository, 512 * 1000)
				utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive without table creation for key: '%s'", keyParam, fi.Name()))
			}

		}
	}

    agencyKey := keyParam
    schema := fmt.Sprintf("gtfs_%s", agencyKey)

	ic.importPostProcess(schema)
	ic.importStopTimesFull(schema)
	ic.updateAgenciesMetaData(agencyKey, schema)
    service.BuildTripCache(ic.db, ic.connectInfos, ic.redis, agencyKey, schema)

	log.Printf("-----------------------------------------------------------------------------------")
	log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
	log.Printf("-----------------------------------------------------------------------------------")

	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}


func (ic *ImportController) ImportMetaData(w http.ResponseWriter, r *http.Request) {

    defer utils.RecoverFromError(w)

    sw := stopwatch.Start(0)

    params := mux.Vars(r)
    keyParam := params["key"]
    agencyKey := keyParam
    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    ic.importPostProcess(schema)
    ic.importStopTimesFull(schema)
    ic.updateAgenciesMetaData(agencyKey, keyParam)

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")

    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}


func (ic *ImportController) BuildTripCache(w http.ResponseWriter, r *http.Request) {

    defer utils.RecoverFromError(w)

    sw := stopwatch.Start(0)

    params := mux.Vars(r)
    keyParam := params["key"]

    agencyKey := keyParam

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    service.BuildTripCache(ic.db, ic.connectInfos, ic.redis, agencyKey, schema)

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")

    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}


func (ic *ImportController) createIndex(schema string, tableName string, indexName string, doneChan chan CreateIndexResult) {
	log.Printf("Creating index with name: '%s_idx' on field '%s' for table with name: '%s.%s'", indexName, indexName, schema, tableName)

    err := database.CreateIndex(ic.db, ic.connectInfos, schema, tableName, indexName)

	doneChan <- CreateIndexResult{indexName, err}
}


func (ic *ImportController) populateTable(schema string, tableName string) {

	filePath := fmt.Sprintf("resources/ddl/%s/insert-%s.sql", ic.connectInfos.Dialect, tableName)

	log.Printf("Inserting data into table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

	ddl, err := data.Asset(filePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table '%s.%s'", filePath, schema, tableName))

	re := regexp.MustCompile("%s")
	insertStmt := re.ReplaceAllString(string(ddl), schema)

	log.Printf("Insert statement: %s", insertStmt)
	err = ic.db.Exec(insertStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not insert into '%s' table", tableName))
}
