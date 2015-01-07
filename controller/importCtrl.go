package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"sort"
	"regexp"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/helyx-io/gtfs-playground/config"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/service"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/data"
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
	r.HandleFunc("/{key}/stop_times_full", importController.ImportStopTimesFull)
	r.HandleFunc("/{key}/post_process", importController.ImportPostProcess)
	r.HandleFunc("/{key}/{file}", importController.Import)

	// Init Repository Map
	initRepositoryMap()
}

func (ac *ImportController) ImportPostProcess(w http.ResponseWriter, r *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
	keyParam := params["key"]

	importPostProcess(keyParam)

	log.Println("-----------------------------------------------------------------------------------")
	log.Println(fmt.Sprintf("--- All Done. ElapsedTime: %v", sw.ElapsedTime()))
	log.Println("-----------------------------------------------------------------------------------")
	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func (ac *ImportController) ImportStopTimesFull(w http.ResponseWriter, r *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
	keyParam := params["key"]

	importStopTimesFull(keyParam)

	log.Println("-----------------------------------------------------------------------------------")
	log.Println(fmt.Sprintf("--- All Done. ElapsedTime: %v", sw.ElapsedTime()))
	log.Println("-----------------------------------------------------------------------------------")
	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
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

	log.Println(fmt.Sprintf("Inserting data into table with name: `gtfs_%s`.`%s` with query from file path: '%s'", schema, tableName, filePath))

	ddlBytes, err := data.Asset(filePath)
	ddl := string(ddlBytes)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table `gtfs_%s`.`%s`", filePath, schema, tableName))

	var lines Lines

	fullTableName := fmt.Sprintf("`gtfs_%s`.`%s`", schema, "lines")
	err = config.DB.Table(fullTableName).Find(&lines).Error
	utils.FailOnError(err, fmt.Sprintf("Could get lines from table %s", fullTableName))

	insertLineDoneChan := make(chan InsertLineResult, 8)

	for _, line := range lines {
		go insertForLine(schema, tableName, ddl, line, insertLineDoneChan)
//		log.Println(fmt.Sprintf("--- Inserting data for line: %s [%s, %s, %s]", line.Name, schema, tableName, ddl))
	}

	doneCount := 0
	for insertLineResult := range insertLineDoneChan {
		if insertLineResult.Error != nil {
			log.Println(fmt.Sprintf("Received event on done chan with error: %s", insertLineResult.Error))
		} else {
			doneCount += 1
			if len(lines) == doneCount {
				log.Println(fmt.Sprintf("Closing done chan."))
				close(insertLineDoneChan)
			} else {
				log.Println(fmt.Sprintf("Received event on done chan for line %s.", insertLineResult.Line.Name))
			}
		}
	}


	createIndexDoneChan := make(chan CreateIndexResult, 8)
	indexes := []string{"service_id", "stop_id", "trip_id", "route_id"}
	for _, index := range indexes {
		go createIndex(schema, tableName, index, createIndexDoneChan)
	}

	doneCount = 0
	for createIndexResult := range createIndexDoneChan {
		if createIndexResult.Error != nil {
			log.Println(fmt.Sprintf("[CREATE_INDEX] Received event on done chan for index :%s with error: %s", createIndexResult.Index, createIndexResult.Error))
		} else {
			doneCount += 1
			if len(indexes) == doneCount {
				log.Println(fmt.Sprintf("[CREATE_INDEX] Closing done chan."))
				close(createIndexDoneChan)
			} else {
				log.Println(fmt.Sprintf("[CREATE_INDEX] Received event on done chan for index %s.", createIndexResult.Index))
			}
		}
	}
}

func insertForLine(schema string, tableName string, ddl string, line Line, doneChan chan InsertLineResult) {
	log.Println(fmt.Sprintf("--- Inserting data for line: %s", line.Name))

	insertStmt := regexp.MustCompile("%s").ReplaceAllString(ddl, schema)
	insertStmt = regexp.MustCompile("%v").ReplaceAllString(insertStmt, line.Name)

	//		log.Println(fmt.Sprintf("Insert statement: %s", insertStmt))
	err := config.DB.Exec(insertStmt).Error
	doneChan <- InsertLineResult{line, err}
}

func (ac *ImportController) Import(w http.ResponseWriter, r *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
	keyParam := params["key"]
	var fileParam string = params["file"]

	if _, ok := config.DataResources[keyParam]; !ok {
		log.Println(fmt.Sprintf("Cannot import agencies for Key: '%s'. key does not exist", keyParam))
		w.WriteHeader(404)
		return
	}

	log.Println(fmt.Sprintf("Importing agencies for Key: %s ...", keyParam))

	if fileParam != "" {
		log.Println(fmt.Sprintf("Processing on for file: %s ...", fileParam))
	}

	w.Header().Set("Content-Type", "text/html")

	folderFilename := config.TmpDir + "/" + keyParam
	url := config.DataResources[keyParam]
	zipFilename := config.TmpDir + "/" + keyParam + ".zip"

	utils.DownloadFile(url, zipFilename)
	utils.UnzipArchive(zipFilename, folderFilename)
	fis := utils.ReadDirectoryFileInfos(folderFilename)
	sort.Sort(utils.FileInfosBySize(fis))

	err := config.GTFS.CreateSchema(keyParam)
	utils.FailOnError(err, fmt.Sprintf("Could not create schema for key: '%s'", keyParam))

	for _, fi := range fis {
		if fi.Mode().IsRegular() {
			gtfsModelRepository := repositoryByFilenameMap[fi.Name()]

			if gtfsModelRepository == nil {
				log.Println(fmt.Sprintf("Filename '%v' is not available in map", fi.Name()))
				continue;
			}

			if fileParam != "" && fileParam + ".txt" != fi.Name() {
				log.Println(fmt.Sprintf("Filename '%v' is not filtered - Does not match with: '%v'", fi.Name(), fileParam))
				continue;
			}

			gaf := service.NewGTFSArchiveFile(fi)

			err := gaf.ImportGTFSArchiveFileWithTableCreation(keyParam, folderFilename, gtfsModelRepository, 512 * 1000)
			utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive with table creation for key: '%s'", keyParam, fi.Name()))

			if fi.Name() == "agency.txt" {
				log.Println("Importing agencies in GTFS agencies table ...")

				gtfsAgencyModelRepository := config.GTFS.GtfsAgencies()
				gaf := service.NewGTFSArchiveFile(fi)

				err:= gaf.ImportGTFSArchiveFileWithoutTableCreation(keyParam, folderFilename, gtfsAgencyModelRepository, 512 * 1000)
				utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive without table creation for key: '%s'", keyParam, fi.Name()))
			}

		}
	}

	importPostProcess(keyParam)
	importStopTimesFull(keyParam)

	log.Println("-----------------------------------------------------------------------------------")
	log.Println(fmt.Sprintf("--- All Done. ElapsedTime: %v", sw.ElapsedTime()))
	log.Println("-----------------------------------------------------------------------------------")
	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}

func createTable(schema string, tableName string) {

	log.Println(fmt.Sprintf("Drop table with name: `gtfs_%s`.`%s`", schema, tableName))

	dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS `gtfs_%s`.`%s`", schema, tableName)
	log.Println(fmt.Sprintf("Create statement: %s", dropStmt))
	err := config.DB.Exec(dropStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not drop '%s' table", tableName))


	filePath := fmt.Sprintf("resources/ddl/create-table-%s.sql", tableName)
	log.Println(fmt.Sprintf("Creating table with name: `gtfs_%s`.`%s` with query from file path: '%s'", schema, tableName, filePath))

	dml, err := data.Asset(filePath)
	utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for create of table `gtfs_%s`.`%s`", filePath, schema, tableName))
	createStmt := fmt.Sprintf(string(dml), schema)
	log.Println(fmt.Sprintf("Create statement: %s", createStmt))
	err = config.DB.Exec(createStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not create '%s' table", tableName))
}

func createIndex(schema string, tableName string, indexName string, doneChan chan CreateIndexResult) {
	log.Println(fmt.Sprintf("Creating index with name: `%s_idx` on field `%s` for table with name: `gtfs_%s`.`%s`", indexName, indexName, schema, tableName))

	createIndexStmt := fmt.Sprintf("ALTER TABLE `gtfs_%s`.`%s` ADD INDEX `%s_idx` (`%s` ASC);", schema, tableName, indexName, indexName)
	log.Println(fmt.Sprintf("Create statement: %s", createIndexStmt))
	err := config.DB.Exec(createIndexStmt).Error
	doneChan <- CreateIndexResult{indexName, err}
}

func populateTable(schema string, tableName string) {

	filePath := fmt.Sprintf("resources/ddl/insert-%s.sql", tableName)

	log.Println(fmt.Sprintf("Inserting data into table with name: `gtfs_%s`.`%s` with query from file path: '%s'", schema, tableName, filePath))

	ddl, err := data.Asset(filePath)
	utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' for inserts into table `gtfs_%s`.`%s`", filePath, schema, tableName))

	re := regexp.MustCompile("%s")
	insertStmt := re.ReplaceAllString(string(ddl), schema)

	log.Println(fmt.Sprintf("Insert statement: %s", insertStmt))
	err = config.DB.Exec(insertStmt).Error
	utils.FailOnError(err, fmt.Sprintf("Could not insert into '%s' table", tableName))

}
