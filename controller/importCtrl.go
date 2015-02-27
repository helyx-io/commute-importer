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
    "encoding/json"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
    "github.com/helyx-io/gtfs-playground/models"
    "github.com/helyx-io/gtfs-playground/config"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/service"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/data"
    "github.com/jiecao-fm/ssdb"
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

type StopTime struct {
    Arrival_time string     `json:"arrival_time"`
    Departure_time string   `json:"departure_time"`
    Stop_sequence int       `json:"stop_sequence"`
    Stop_name string        `json:"stop_name"`
}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper functions
////////////////////////////////////////////////////////////////////////////////////////////////

func initRepositoryMap() {
	repositoryByFilenameMap = make(map[string]database.GTFSCreatedModelRepository)

	repositoryByFilenameMap["agency.txt"] = config.GTFS.Agencies()
	repositoryByFilenameMap["calendar_dates.txt"] = config.GTFS.CalendarDates()
	repositoryByFilenameMap["calendar.txt"] = config.GTFS.Calendars()
	repositoryByFilenameMap["routes.txt"] = config.GTFS.Routes()
	repositoryByFilenameMap["stops.txt"] = config.GTFS.Stops()
	repositoryByFilenameMap["stop_times.txt"] = config.GTFS.StopTimes()
	repositoryByFilenameMap["transfers.txt"] = config.GTFS.Transfers()
	repositoryByFilenameMap["trips.txt"] = config.GTFS.Trips()

}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Import Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {

	// Init Router
	r.HandleFunc("/{key}", importController.Import)
    r.HandleFunc("/{key}/rewrite", importController.RewriteCsvFiles)
	r.HandleFunc("/{key}/caches/trips", importController.BuildTripsCache)

	// Init Repository Map
	initRepositoryMap()
}

func (ac *ImportController) RewriteCsvFiles(w http.ResponseWriter, r *http.Request) {

    defer utils.RecoverFromError(w)

    sw := stopwatch.Start(0)

    params := mux.Vars(r)
    keyParam := params["key"]

    rewriteCsvFiles(keyParam, "out")

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")

    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func (ac *ImportController) BuildTripsCache(w http.ResponseWriter, r *http.Request) {

    log.Printf("Building trips cache ...")

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
	keyParam := params["key"]

    schema := keyParam

    poolconf := ssdb.PoolConfig{Host: config.RedisInfos.Host, Port: config.RedisInfos.Port, Initial_conn_count: 16, Max_idle_count: 64, Max_conn_count: 256}
    pool, err := ssdb.NewPool(poolconf)
    if err != nil {
        return
    }
    defer pool.Close()

    tripIds := make([]string, 0)

    tripId := ""
    tripIdsQuery := fmt.Sprintf("select trip_id from `gtfs_%s`.`trips` order by trip_id", schema)
    log.Printf("Query: %s", tripIdsQuery)
    rows, err := config.DB.Raw(tripIdsQuery).Rows()
    defer rows.Close()

    for rows.Next() {
        rows.Scan(&tripId)
        tripIds = append(tripIds, tripId);
    }

    log.Printf("TripIds: %d", len(tripIds))

   /* sem := make(chan bool, 64)*/

    for _, tripId := range tripIds {

        /* sem <- true */
        /* go */ func() {

           /* defer func() { <-sem }() */

           db, err := pool.GetDB()

            if err != nil {
                log.Printf("Error: %s", err.Error())
            }

            defer func() {
                if db != nil {
                    pool.ReturnDB(db)
                } else {
                    fmt.Printf("Pool idle count: %d\n", pool.IdleCount())
                }
            }()

            if db == nil {
                log.Printf("db is nil for tripId: %s", tripId)
                return
            }

            stopTimesQuery := fmt.Sprintf("select st.arrival_time, st.departure_time, st.stop_sequence, s.stop_name from `gtfs_%s`.`stop_times` st inner join `gtfs_%s`.`stops` s on st.stop_id=s.stop_id where st.trip_id='%s' order by st.stop_sequence", schema, schema, tripId)

//            log.Printf("Query: %s", stopTimesQuery)

            stopTimeRows, err := config.DB.Raw(stopTimesQuery).Rows()
            defer rows.Close()

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

            bytes, err := json.Marshal(stopTimes)
            if err != nil {
                log.Printf("Error: '%s' ...", err.Error())
            }

//            log.Printf("Selecting stop-times (%d) for tripId: '%s' ...", len(stopTimes), tripId)

            cacheKey := fmt.Sprintf("/agencies/%s/trips/%s/stop-times", keyParam, tripId)
            stopTimesStr := string(bytes)
            err = db.Set(cacheKey, stopTimesStr);
            if err != nil {
                log.Printf("Error: '%s' ...", err.Error())
            }
        }()
    }

    /*for i := 0; i < cap(sem); i++ {
        sem <- true
    }*/


	utils.FailOnError(err, fmt.Sprintf("Build trips cache for agency key: %s", keyParam))

	log.Printf("-----------------------------------------------------------------------------------")
	log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
	log.Printf("-----------------------------------------------------------------------------------")

	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func rewriteCsvFiles(schema, outFolderName string) error {

    agencyIndexes, err := getIndexes(schema, "agency.txt", 0)
    serviceIndexes, err := getIndexes(schema, "trips.txt", 1)
    tripIndexes, err := getIndexes(schema, "trips.txt", 2)
    stopIndexes, err := getIndexes(schema, "stops.txt", 0)
    stopTimesIndexes, err := getIndexes(schema, "stop_times.txt", 3)
    routeIndexes, err := getIndexes(schema, "routes.txt", 0)


    indexes := map[int](map[string]string){ 0: stopIndexes }
    rewriteCsvFile(schema, "stops.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: tripIndexes, 3: stopTimesIndexes }
    rewriteCsvFile(schema, "stop_times.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: routeIndexes, 1: agencyIndexes }
    rewriteCsvFile(schema, "routes.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: agencyIndexes }
    rewriteCsvFile(schema, "agency.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: routeIndexes, 1: serviceIndexes, 2: tripIndexes }
    rewriteCsvFile(schema, "trips.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: serviceIndexes }
    rewriteCsvFile(schema, "calendar.txt", outFolderName, indexes)

    indexes = map[int](map[string]string){ 0: serviceIndexes }
    rewriteCsvFile(schema, "calendar_dates.txt", outFolderName, indexes)

    return err
}

func rewriteCsvFile(schema, filename, outFolderName string, indexes map[int](map[string]string)) error {

    folderName := path.Join(config.TmpDir, schema)
    filePath := path.Join(folderName, filename)

    gtfsFile := models.GTFSFile{filePath}
    
    outFile, err := os.Create(path.Join(folderName, path.Join(outFolderName, filename)))
    if err != nil {
        return err
    }

    writer := csv.NewWriter(outFile)
    headers, err := utils.ReadCsvFileHeader(gtfsFile.Filename, ",")
    log.Printf("headers: %v", headers)

    if err != nil {
        return err
    }

    writer.Write(headers)

    resultChan := make(chan [][]string)

    go func() {

        sem := make(chan bool, 8)

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            sem <- true
            go func() {
                defer func() { <-sem }()
                records, _ := models.ParseCsvAsStringArrays(lines)

                for _, record := range *records {
                    for i, _ := range indexes {
                        record[i] = indexes[i][record[i]]
                    }
                }

                resultChan <- *records
            }()
        }

        for i := 0; i < cap(sem); i++ {
            sem <- true
        }

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

func getIndexes(schema, filename string, index int) (map[string]string, error) {

    folderName := path.Join(config.TmpDir, schema)
    filePath := path.Join(folderName, filename)

    gtfsFile := models.GTFSFile{filePath}

    resultChan := make(chan []string)

    go func() {

        sem := make(chan bool, 8)

        for lines := range gtfsFile.LinesIterator(1024 * 1024) {
            sem <- true
            go func() {
                defer func() { <-sem }()
                records, _ := models.ParseCsvAsStringArrays(lines)

                keys := []string{}
                for _, record := range *records {
                    keys = append(keys, record[index])
                }

                resultChan <- keys
            }()
        }

        for i := 0; i < cap(sem); i++ {
            sem <- true
        }

        close(resultChan)

    }()

    offset := 0
    result := []string{}
    indexes := map[string]string{}
    increment := 0

    for results := range resultChan {
        offset++
        log.Printf("[%s][%d] Records read", filePath, offset)
        for _, val := range results {
            if _, ok := indexes[val]; !ok {
                increment += 1
                result = append(result, val)
                indexes[val] = fmt.Sprintf("%d", increment)
            }
        }
    }

    return indexes, nil
}


func importPostProcess(schema string) {

	createTable(schema, "lines")
	createTable(schema, "line_stops")
	createTable(schema, "stations")
	createTable(schema, "station_stops")
	createTable(schema, "station_lines")
	createTable(schema, "route_stops")

	populateTable(schema, "lines")
	populateTable(schema, "line_stops")
	populateTable(schema, "stations")
	populateTable(schema, "station_stops")
	populateTable(schema, "station_lines")
	populateTable(schema, "route_stops")
}

func importStopTimesFull(schema string) {
//	config.DB.LogMode(true)

	tableName := "stop_times_full"
	createTable(schema, tableName)

	filePath := fmt.Sprintf("resources/ddl/insert-%s.sql", tableName)

	log.Printf("Inserting data into table with name: `gtfs_%s`.`%s` with query from file path: '%s'", schema, tableName, filePath)

	ddlBytes, err := data.Asset(filePath)
	ddl := string(ddlBytes)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table `gtfs_%s`.`%s`", filePath, schema, tableName))

	var lines Lines

	fullTableName := fmt.Sprintf("`gtfs_%s`.`%s`", schema, "lines")
	err = config.DB.Table(fullTableName).Find(&lines).Error
	utils.FailOnError(err, fmt.Sprintf("Could get lines from table %s", fullTableName))

	insertLineDoneChan := make(chan InsertLineResult, 8)

    go func() {
        for _, line := range lines {
            insertForLine(schema, tableName, ddl, line, insertLineDoneChan)
            //		log.Printf("--- Inserting data for line: %s [%s, %s, %s]", line.Name, schema, tableName, ddl)
        }
    }()

	doneCount := 0
	for insertLineResult := range insertLineDoneChan {
		if insertLineResult.Error != nil {
			log.Printf("Received event on done chan with error: %s", insertLineResult.Error)
		} else {
			doneCount += 1
			if len(lines) == doneCount {
				log.Printf("Closing done chan.")
				close(insertLineDoneChan)
			} else {
				log.Printf("Received event on done chan for line %s.", insertLineResult.Line.Name)
			}
		}
	}


	createIndexDoneChan := make(chan CreateIndexResult, 8)
	indexes := []string{"service_id", "stop_id", "trip_id", "route_id"}

    go func() {
        for _, index := range indexes {
            createIndex(schema, tableName, index, createIndexDoneChan)
        }
    }()

	doneCount = 0
	for createIndexResult := range createIndexDoneChan {
		if createIndexResult.Error != nil {
			log.Printf("[CREATE_INDEX] Received event on done chan for index :%s with error: %s", createIndexResult.Index, createIndexResult.Error)
		} else {
			doneCount += 1
			if len(indexes) == doneCount {
				log.Printf("[CREATE_INDEX] Closing done chan.")
				close(createIndexDoneChan)
			} else {
				log.Printf("[CREATE_INDEX] Received event on done chan for index %s.", createIndexResult.Index)
			}
		}
	}
}

func updateAgenciesMetaData(schema string) error {
	dbSql, err := sql.Open(config.ConnectInfos.Dialect, config.ConnectInfos.URL)

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()


	log.Printf("Updating agency zone for schema: %s", schema)

	selectFilePath := "resources/ddl/select-agency-zone.sql"
	ddlSelect, err := data.Asset(selectFilePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for fetching agency zone: `%s`", selectFilePath, schema))

	selectStmt := fmt.Sprintf(string(ddlSelect), schema)

	log.Printf("Fetch agency zone infos for schema: '%s': '%s'", schema, selectStmt)

	row := config.DB.Raw(selectStmt).Row()

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


	updateFilePath := "resources/ddl/update-agency-zone.sql"
	ddlUpdate, err := data.Asset(updateFilePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for updating agency zone: `%s`", updateFilePath, schema))

	updateStmt := fmt.Sprintf(string(ddlUpdate), schema)

	log.Printf("Fetch agency zone infos for schema: '%s': '%s'", schema, updateStmt)

	updateValueArgs := []interface{}{ min_stop_lat, max_stop_lat, min_stop_lon, max_stop_lon }


	_, err = dbSql.Exec(updateStmt, updateValueArgs...)

	if err != nil {
		log.Println(fmt.Println("Failed on Error: %v", err))
		return err
	}


	updateGtfsFilePath := "resources/ddl/update-gtfs-agency-zone.sql"
	ddlUpdateGtfs, err := data.Asset(updateGtfsFilePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for updating agency zone: `%s`", updateGtfsFilePath, schema))

	updateGtfsStmt := string(ddlUpdateGtfs)

	log.Printf("Fetch agency zone infos for schema: '%s': '%s'", schema, updateGtfsStmt)

	updateGtfsValueArgs := []interface{}{ min_stop_lat, max_stop_lat, min_stop_lon, max_stop_lon, schema }

	_, err = dbSql.Exec(updateGtfsStmt, updateGtfsValueArgs...)

	if err != nil {
		log.Printf("Failed on Error: %v", err)
	}

	return err
}

func insertForLine(schema string, tableName string, ddl string, line Line, doneChan chan InsertLineResult) {
	log.Printf("--- Inserting data for line: %s", line.Name)

	insertStmt := regexp.MustCompile("%s").ReplaceAllString(ddl, schema)
	insertStmt = regexp.MustCompile("%v").ReplaceAllString(insertStmt, line.Name)

	//		log.Printf("Insert statement: %s", insertStmt)
	err := config.DB.Exec(insertStmt).Error
	doneChan <- InsertLineResult{line, err}
}

func (ac *ImportController) Import(w http.ResponseWriter, r *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
	keyParam := params["key"]

	log.Printf("Importing agencies for Key: %s ...", keyParam)

	w.Header().Set("Content-Type", "text/html")

	folderFilename := config.TmpDir + "/" + keyParam
	url := config.DataResources[keyParam]

	zipFilename := config.TmpDir + "/" + keyParam + "-" + time.Now().Format("20060102-150405") + ".zip"

	utils.DownloadFile(url, zipFilename)
	utils.UnzipArchive(zipFilename, folderFilename)

    outFolderFilename := path.Join(folderFilename, "out")
    
    if os.MkdirAll(outFolderFilename, 0755) != nil {
        panic("Unable to create directory for tagfile!")
    }
    
    rewriteCsvFiles(keyParam, "out")

	fis := utils.ReadDirectoryFileInfos(outFolderFilename)
	sort.Sort(utils.FileInfosBySize(fis))

	err := config.GTFS.CreateSchema(keyParam)
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

				gtfsAgencyModelRepository := config.GTFS.GtfsAgencies()
				gaf := service.NewGTFSArchiveFile(fi)

				err:= gaf.ImportGTFSArchiveFileWithoutTableCreation(keyParam, outFolderFilename, gtfsAgencyModelRepository, 512 * 1000)
				utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive without table creation for key: '%s'", keyParam, fi.Name()))
			}

		}
	}

	importPostProcess(keyParam)
	importStopTimesFull(keyParam)
	updateAgenciesMetaData(keyParam)

	log.Printf("-----------------------------------------------------------------------------------")
	log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
	log.Printf("-----------------------------------------------------------------------------------")

	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func (ac *ImportController) ImportMetaData(w http.ResponseWriter, r *http.Request) {

    defer utils.RecoverFromError(w)

    sw := stopwatch.Start(0)

    params := mux.Vars(r)
    keyParam := params["key"]

    importPostProcess(keyParam)
    importStopTimesFull(keyParam)
    updateAgenciesMetaData(keyParam)

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")

    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func createTable(schema string, tableName string) {

	log.Printf("Drop table with name: `gtfs_%s`.`%s`", schema, tableName)

	dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS `gtfs_%s`.`%s`", schema, tableName)
	log.Printf("Create statement: %s", dropStmt)
	err := config.DB.Exec(dropStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not drop '%s' table", tableName))


	filePath := fmt.Sprintf("resources/ddl/create-table-%s.sql", tableName)
	log.Printf("Creating table with name: `gtfs_%s`.`%s` with query from file path: '%s'", schema, tableName, filePath)

	dml, err := data.Asset(filePath)
	utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for create of table `gtfs_%s`.`%s`", filePath, schema, tableName))
	createStmt := fmt.Sprintf(string(dml), schema)
	log.Printf("Create statement: %s", createStmt)
	err = config.DB.Exec(createStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not create '%s' table", tableName))
}

func createIndex(schema string, tableName string, indexName string, doneChan chan CreateIndexResult) {
	log.Printf("Creating index with name: `%s_idx` on field `%s` for table with name: `gtfs_%s`.`%s`", indexName, indexName, schema, tableName)

	createIndexStmt := fmt.Sprintf("ALTER TABLE `gtfs_%s`.`%s` ADD INDEX `%s_idx` (`%s` ASC);", schema, tableName, indexName, indexName)
	log.Printf("Create statement: %s", createIndexStmt)
	err := config.DB.Exec(createIndexStmt).Error
	doneChan <- CreateIndexResult{indexName, err}
}

func populateTable(schema string, tableName string) {

	filePath := fmt.Sprintf("resources/ddl/insert-%s.sql", tableName)

	log.Printf("Inserting data into table with name: `gtfs_%s`.`%s` with query from file path: '%s'", schema, tableName, filePath)

	ddl, err := data.Asset(filePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table `gtfs_%s`.`%s`", filePath, schema, tableName))

	re := regexp.MustCompile("%s")
	insertStmt := re.ReplaceAllString(string(ddl), schema)

	log.Printf("Insert statement: %s", insertStmt)
	err = config.DB.Exec(insertStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not insert into '%s' table", tableName))
}
