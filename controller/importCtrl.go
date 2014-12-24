package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"sort"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/helyx-io/gtfs-playground/config"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/service"
	"github.com/helyx-io/gtfs-playground/database"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	repositoryByFilenameMap map[string]database.GTFSModelRepository
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
	repositoryByFilenameMap = make(map[string]database.GTFSModelRepository)

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

			gaf.ImportGTFSArchiveFile(keyParam, folderFilename, gtfsModelRepository, 2048 * 1000, config.WorkPool)
		}
	}

	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v", sw.ElapsedTime())))
}
