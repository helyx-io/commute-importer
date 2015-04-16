package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "os"
	"fmt"
	"log"
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
}

func (ic *ImportController) Import(w http.ResponseWriter, r *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	params := mux.Vars(r)
    keyParam := params["key"]

    agencyKey := keyParam
    schema := fmt.Sprintf("gtfs_%s", agencyKey)

	log.Printf("Importing agencies for Key: '%s' ...", agencyKey)

	w.Header().Set("Content-Type", "text/html")

	folderFilename := ic.tmpDir + "/" + agencyKey
	url := ic.dataResources[agencyKey]

	zipFilename := ic.tmpDir + "/" + agencyKey + "-" + time.Now().Format("20060102-150405") + ".zip"

	utils.DownloadFile(url, zipFilename)
	utils.UnzipArchive(zipFilename, folderFilename)

    outFolderFilename := path.Join(folderFilename, "out")

    if os.MkdirAll(outFolderFilename, 0755) != nil {
        panic("Unable to create directory for tagfile!")
    }

    columnLengthsByFiles, err := service.NewCsvFileRewriter(ic.tmpDir).RewriteCsvFiles(agencyKey, "out")
    utils.FailOnError(err, "Could not rewrite csv files with success")

    service.NewCsvFileImporter(ic.driver, ic.gtfs).ImportCsvFiles(agencyKey, outFolderFilename, columnLengthsByFiles)
    service.NewComplementaryTablePopuler(ic.driver).Populate(schema, columnLengthsByFiles)
    service.NewStopTimesFullImporter(ic.driver).ImportStopTimesFull(schema, columnLengthsByFiles)
    service.NewAgenciesMetadataUpdater(ic.driver).UpdateAgenciesMetaData(agencyKey, schema)
    service.NewTripCacheBuilder(ic.driver, ic.redis).BuildTripCache(agencyKey, schema)

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
