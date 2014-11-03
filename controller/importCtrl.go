package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"fmt"
	"log"
	"path"
	"sort"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/helyx-io/gtfs-playground/config"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	repositoryByFilenameMap map[string]database.GTFSModelRepository
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type FileInfos []os.FileInfo

func (fis FileInfos) Len() int {
	return len(fis)
}
func (fis FileInfos) Less(i, j int) bool {
	return fis[i].Size() < fis[j].Size()
}
func (fis FileInfos) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func recoverFromError(w http.ResponseWriter) {
	if r := recover(); r != nil {
		err, _ := r.(error)
		http.Error(w, err.Error(), 500)
		return
	}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Import Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {

	// Init Router
	r.HandleFunc("/", importController.Import)

	// Init Repository Map
	repositoryByFilenameMap = make(map[string]database.GTFSModelRepository)

	repositoryByFilenameMap["stop_times.txt"] = config.GTFS.StopTimes()
	repositoryByFilenameMap["stops.txt"] = config.GTFS.Stops()
	repositoryByFilenameMap["agency.txt"] = config.GTFS.Agencies()
}

func (ac *ImportController) Import(w http.ResponseWriter, _ *http.Request) {

	defer recoverFromError(w)

	sw := stopwatch.Start(0)

	w.Header().Set("Content-Type", "text/html")

	log.Println("Importing agencies ...")

	folderFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502"
	url := "http://localhost/data/gtfs_paris_20140502.zip"
	zipFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502.zip"

	downloadGTFSArchive(url, zipFilename)
	unzipGTFSArchive(zipFilename, folderFilename)
	fis := readGTFSArchiveFileInfos(folderFilename)

	for _, fi := range fis {
		processGTFSArchiveFile(folderFilename, fi, config.WorkPool)
	}

	w.Write([]byte(fmt.Sprintf("ElapsedTime: %v ms", sw.ElapsedTime())))
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Private Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func downloadGTFSArchive(url, zipFilename string) {

	log.Println(fmt.Sprintf(" - Downloading zip file from url: '%v' to file path: '%v' ...", url, zipFilename))

	sw := stopwatch.Start(0)

	writtenBytes, err := utils.DownloadFileFromURL(url, zipFilename)
	utils.FailOnError(err, fmt.Sprintf("Could not download file from url: '%v' to file path: '%v'", url, zipFilename))

	log.Println(fmt.Sprintf(" - Downloaded zip file: '%v' - %v bytes - Duration: %v", zipFilename, writtenBytes, sw.ElapsedTime()))

}

func unzipGTFSArchive(zipFilename, folderFilename string) {

	log.Println(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename))

	sw := stopwatch.Start(0)

	err := utils.Unzip(zipFilename, folderFilename)
	utils.FailOnError(err, fmt.Sprintf("Could unzip filename: '%v' to folder: '%v'", zipFilename, folderFilename))

	log.Println(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v' - Duration: %v", zipFilename, folderFilename, sw.ElapsedTime()))
}

func readGTFSArchiveFileInfos(folderFilename string) FileInfos {

	d, err := os.Open(folderFilename)
	utils.FailOnError(err, fmt.Sprintf("Could not open directory '%v' for read", folderFilename))
	defer d.Close()

	fisArr, err := d.Readdir(-1)
	fis := FileInfos(fisArr)
	utils.FailOnError(err, fmt.Sprintf("Could not read directory '%v' content", folderFilename))

	sort.Sort(fis)

	return fis
}

func processGTFSArchiveFile(folderFilename string, fi os.FileInfo, workPool *workpool.WorkPool) {
	if fi.Mode().IsRegular() {
		gtfsModelRepository := repositoryByFilenameMap[fi.Name()]
		if (gtfsModelRepository == nil) {
			log.Println(fmt.Sprintf("Filename '%v' is not available in map", fi.Name()))
			return;
		}

		log.Println(fmt.Sprintf("Filename '%v' is available in map - Reading File with size: %d bytes ...", fi.Name(), fi.Size()))

		sw := stopwatch.Start(0)

		insertGTFSModels(folderFilename, gtfsModelRepository, fi.Name(), workPool)

		log.Println(fmt.Sprintf(" - 	Read file: '%v' - Duration: %v", fi.Name(), sw.ElapsedTime()))
	}
}

func insertGTFSModels(folderFilename string, gtfsModel database.GTFSModelRepository, modelsFilename string, workPool *workpool.WorkPool) {

	offset := 0

	log.Println(fmt.Sprintf(" - Removing entries from repository related to file with name: '%v' ...", modelsFilename))
	gtfsModel.RemoveAllByAgencyKey("RATP")
	log.Println(fmt.Sprintf(" - Removed entries from repository related to file with name: '%v'", modelsFilename))

	gtfsFile := models.GTFSFile{path.Join(folderFilename, modelsFilename)}

	for lines := range gtfsFile.LinesIterator() {

		offset++

		log.Println(fmt.Sprintf(" - Inserting chunk of data with offset: '%d' related to file with name: '%v'", offset, modelsFilename))

		taskName := fmt.Sprintf("ChunkImport-%d", offset)
		task := gtfsModel.CreateImportTask(taskName, "RATP", lines, workPool)

		err := workPool.PostWork("import", task)

		utils.FailOnError(err, fmt.Sprintf("Could not post work with offset: %d", offset))
	}

}
