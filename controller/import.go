package controller

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/akinsella/go-playground/database/mysql"
	"github.com/akinsella/go-playground/utils"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/database"
	"github.com/goinggo/workpool"

)

const (
	url = "http://localhost/data/gtfs_paris_20140502.zip"
	zipFilename = "/Users/akinsella/Desktop/gtfs_paris_20140502.zip"
	folderFilename = "/Users/akinsella/Desktop/gtfs_paris_20140502"
)

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {
	r.HandleFunc("/", importController.Import)
}

func (ac *ImportController) Import(w http.ResponseWriter, _ *http.Request) {

	var err error

	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	sw := stopwatch.Start(0)

	w.Header().Set("Content-Type", "text/html")

	log.Println("Importing agencies ...")

	log.Println(fmt.Sprintf(" - Downloading zip file from url: '%v' to file path: '%v' ...", url, zipFilename))

	writtenBytes, err := utils.DownloadFileFromURL(url, zipFilename)
	utils.FailOnError(err, fmt.Sprintf("Could not download file from url: '%v' to file path: '%v'", url, zipFilename))

	log.Println(fmt.Sprintf(" - Downloaded zip file: '%v' - %v bytes - ElapsedTime: %v", zipFilename, writtenBytes, sw.ElapsedTime()))

	log.Println(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename))

	swZip := stopwatch.Start(0)

	err = utils.Unzip(zipFilename, folderFilename)
	utils.FailOnError(err, fmt.Sprintf("Could unzip filename: '%v' to folder: '%v'", zipFilename, folderFilename))

	log.Println(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v' - ElapsedTime: %v - Duration: %v", zipFilename, folderFilename, sw.ElapsedTime(), swZip.ElapsedTime()))


	d, err := os.Open(folderFilename)
	utils.FailOnError(err, fmt.Sprintf("Could not open directory '%v' for read", folderFilename))
	defer d.Close()

	fi, err := d.Readdir(-1)
	utils.FailOnError(err, fmt.Sprintf("Could not read directory '%v' content", folderFilename))


	workPool := workpool.New(32, 10000)

	db, err := mysql.InitDb(2, 100);
	utils.FailOnError(err, "Could not open database")
	defer db.Close()

	gtfs := mysql.CreateMySQLGTFSRepository(db)
	repositoryByFilenameMap := make(map[string]database.GTFSModelRepository)

	repositoryByFilenameMap["stop_times.txt"] = gtfs.StopTimes()
	repositoryByFilenameMap["stops.txt"] = gtfs.StopTimes()

	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			gtfsModelRepository := repositoryByFilenameMap[fi.Name()]
			if (gtfsModelRepository == nil) {
				log.Println(fmt.Sprintf("Filename '%v' is not available in map", fi.Name()))
				continue;
			}


			log.Println(fmt.Sprintf("Filename '%v' is available in map - Reading File with size: %d bytes ...", fi.Name(), fi.Size()))

			swReadFile := stopwatch.Start(0)

			insertModels(gtfsModelRepository, fi.Name(), workPool)

			log.Println(fmt.Sprintf(" - 	Read file: '%v' - ElapsedTime: %v - Duration: %v", fi.Name(), sw.ElapsedTime(), swReadFile.ElapsedTime()))
		}
	}

	w.Write([]byte(fmt.Sprintf("ElapsedTime: ", sw.ElapsedTime(), "ms")))
}

func insertModels(gtfsModel database.GTFSModelRepository, modelsFilename string, workPool *workpool.WorkPool) {

	offset := 0

	gtfsModel.RemoveAllByAgencyKey("RATP")

	gtfsFile := models.GTFSFile{modelsFilename}

	for lines := range gtfsFile.LinesIterator() {

		offset++
		taskName := fmt.Sprintf("ChunkImport-%d", offset)
		task := gtfsModel.CreateImportTask(taskName, lines, workPool)

		err := workPool.PostWork("import", task)

		utils.FailOnError(err, fmt.Sprintf("Could not post work with offset: %d", offset))
	}

}
