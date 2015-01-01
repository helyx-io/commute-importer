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
/// Import Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {

	// Init Router
	r.HandleFunc("/{key}", importController.Import)
	r.HandleFunc("/{key}/{file}", importController.Import)

	// Init Repository Map
	initRepositoryMap()
}

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

			err := gaf.ImportGTFSArchiveFileWithTableCreation(keyParam, folderFilename, gtfsModelRepository, 2048 * 1000, config.WorkPool)
			utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive with table creation for key: '%s'", keyParam, fi.Name()))

			if fi.Name() == "agency.txt" {
				log.Println("Importing agencies in GTFS agencies table ...")

				gtfsAgencyModelRepository := config.GTFS.GtfsAgencies()
				gaf := service.NewGTFSArchiveFile(fi)

				err:= gaf.ImportGTFSArchiveFileWithoutTableCreation(keyParam, folderFilename, gtfsAgencyModelRepository, 2048 * 1000, config.WorkPool)
				utils.FailOnError(err, fmt.Sprintf("[%s] Could not import gtfs archive without table creation for key: '%s'", keyParam, fi.Name()))
			}

		}
	}

	createTable(keyParam, "lines")
	createTable(keyParam, "line_stops")
	createTable(keyParam, "stations")
	createTable(keyParam, "station_stops")
	createTable(keyParam, "station_lines")

	populateTable(keyParam, "lines")
	populateTable(keyParam, "line_stops")
	populateTable(keyParam, "stations")
	populateTable(keyParam, "station_stops")
	populateTable(keyParam, "station_lines")

	log.Println("-----------------------------------------------------------------------------------")
	log.Println(fmt.Println("--- All Done. ElapsedTime: %v", sw.ElapsedTime()))
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
