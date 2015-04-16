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
	"net/http"

    "gopkg.in/redis.v2"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/helyx-io/gtfs-importer/utils"
	"github.com/helyx-io/gtfs-importer/service"
	"github.com/helyx-io/gtfs-importer/database"
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
    driver *database.Driver
    gtfs database.GTFSRepository
    redis *redis.Client
    dataResources map[string]string
    tmpDir string
}

func (ic *ImportController) Init(r *mux.Router, dataResources map[string]string, tmpDir string, driver *database.Driver, redis *redis.Client, gtfs database.GTFSRepository) {

    ic.gtfs = gtfs
    ic.driver = driver
    ic.redis = redis
    ic.tmpDir = tmpDir
    ic.dataResources = dataResources


	// Init Router
    r.HandleFunc("/{key}", ic.Import)
    r.HandleFunc("/{key}/caches/trips", ic.BuildTripCache)

	// Init Repository Map
	initRepositoryMap(gtfs)
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

    columnLengthsByFiles, err := service.NewCsvFileRewriter(ic.tmpDir).RewriteCsvFiles(keyParam, "out")

	fis := utils.ReadDirectoryFileInfos(outFolderFilename)
	sort.Sort(utils.FileInfosBySize(fis))

	err = ic.gtfs.CreateSchema(keyParam)
	utils.FailOnError(err, fmt.Sprintf("Could not create schema for key: '%s'", keyParam))

	for _, fi := range fis {
		if fi.Mode().IsRegular() {
			gtfsModelRepository := repositoryByFilenameMap[fi.Name()]

			if gtfsModelRepository == nil {
				log.Printf("Filename '%v' is not available in map", fi.Name())
				continue;
			}

			gaf := service.NewGTFSArchiveFile(fi)

            log.Printf("columnLengthsByFiles:  %v - gaf.Name(): %v", columnLengthsByFiles, gaf.Name())

            columnLengthsByFile := columnLengthsByFiles[gaf.Name()]

            columnLengthsByFileTmp := make(map[string]interface{})
            for key, value := range columnLengthsByFile {
                columnLengthsByFileTmp[key] = value
            }

            log.Printf("ImportGTFSArchiveFileWithTableCreation:  %v", columnLengthsByFileTmp)

            err := gaf.ImportGTFSArchiveFileWithTableCreation(keyParam, outFolderFilename, gtfsModelRepository, columnLengthsByFileTmp, 512 * 1000)
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

    service.NewComplementaryTablePopuler(ic.driver).Populate(schema)
    service.NewStopTimesFullImporter(ic.driver).ImportStopTimesFull(schema)
    service.NewAgenciesMetadataUpdater(ic.driver).UpdateAgenciesMetaData(agencyKey, schema)
    service.NewTripCacheBuilder(ic.driver, ic.redis).BuildTripCache(agencyKey, schema)

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

    service.NewComplementaryTablePopuler(ic.driver).Populate(schema)
    service.NewStopTimesFullImporter(ic.driver).ImportStopTimesFull(schema)
    service.NewAgenciesMetadataUpdater(ic.driver).UpdateAgenciesMetaData(agencyKey, keyParam)

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

    service.NewTripCacheBuilder(ic.driver, ic.redis).BuildTripCache(agencyKey, schema)

    log.Printf("-----------------------------------------------------------------------------------")
    log.Printf("--- All Done. ElapsedTime: %v", sw.ElapsedTime())
    log.Printf("-----------------------------------------------------------------------------------")

    w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}
