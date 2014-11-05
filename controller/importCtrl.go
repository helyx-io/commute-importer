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
	r.HandleFunc("/", importController.Import)

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

func (ac *ImportController) Import(w http.ResponseWriter, _ *http.Request) {

	defer utils.RecoverFromError(w)

	sw := stopwatch.Start(0)

	w.Header().Set("Content-Type", "text/html")

	log.Println("Importing agencies ...")

	folderFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502"
	url := "http://localhost/data/gtfs_paris_20140502.zip"
	zipFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502.zip"

	utils.DownloadFile(url, zipFilename)
	utils.UnzipArchive(zipFilename, folderFilename)
	fis := utils.ReadDirectoryFileInfos(folderFilename)
	sort.Sort(utils.FileInfosBySize(fis))

	for _, fi := range fis {
		if fi.Mode().IsRegular() {
			gtfsModelRepository := repositoryByFilenameMap[fi.Name()]

			if (gtfsModelRepository == nil) {
				log.Println(fmt.Sprintf("Filename '%v' is not available in map", fi.Name()))
				continue;
			}

			gaf := service.NewGTFSArchiveFile(fi)

			gaf.ImportGTFSArchiveFile(folderFilename, gtfsModelRepository, config.WorkPool)
		}
	}

	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v ms", sw.ElapsedTime())))
}
