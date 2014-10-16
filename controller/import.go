package controller

import (
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/akinsella/go-playground/database/mysql"
	"github.com/akinsella/go-playground/tasks"
	"github.com/akinsella/go-playground/utils"
	"github.com/goinggo/workpool"
)

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {
	r.HandleFunc("/", importController.Import)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(err)
	}
}

func (ac *ImportController) Import(w http.ResponseWriter, r *http.Request) {

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

	w.Write([]byte("Importing agencies ..."))
	w.Write([]byte("<br/>"))

	url := "http://localhost/data/gtfs_paris_20140502.zip"
	zipFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502.zip"
	folderFilename := "/Users/akinsella/Desktop/gtfs_paris_20140502"

	w.Write([]byte(fmt.Sprintf(" - Downloading zip file from url: '%v' to file path: '%v' ...", url, zipFilename)))
	w.Write([]byte("<br/>"))

	out, err := os.Create(zipFilename)
	defer out.Close()
	failOnError(err, fmt.Sprintf("Could not create file with name: '%v'", zipFilename))

	resp, err := http.Get(url)
	defer resp.Body.Close()
	failOnError(err, fmt.Sprintf("Could not get content from url: '%v'", url))

	writtenBytes, err := io.Copy(out, resp.Body)
	failOnError(err, "Could not copy response body to out")

	w.Write([]byte(fmt.Sprintf(" - Downloaded zip file: '%v' - %v bytes - ElapsedTime: %v", zipFilename, writtenBytes, sw.ElapsedTime())))
	w.Write([]byte("<br/>"))

	w.Write([]byte(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename)))
	w.Write([]byte("<br/>"))

	swZip := stopwatch.Start(0)

	err = utils.Unzip(zipFilename, folderFilename)
	failOnError(err, fmt.Sprintf("Could unzip filename: '%v' to folder: '%v'", zipFilename, folderFilename))

	w.Write([]byte(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v' - ElapsedTime: %v - Duration: %v", zipFilename, folderFilename, sw.ElapsedTime(), swZip.ElapsedTime())))
	w.Write([]byte("<br/>"))

	stopsFilename := fmt.Sprintf("%v/stop_times.txt", folderFilename)

	w.Write([]byte(fmt.Sprintf(" - Reading file: '%v'", stopsFilename)))
	w.Write([]byte("<br/>"))

	swReadFile := stopwatch.Start(0)

	workPool := workpool.New(32, 10000)

	offset := 0

	for lines := range LinesIterator(stopsFilename) {

		offset++

		task := mysql.MySQLStopTimesImportTask {
			tasks.ImportTask {
				Name: fmt.Sprintf("ChunkImport-%d", offset),
				Lines: lines,
				WP: workPool,
			},
		}

		err := workPool.PostWork("import", &task)

		failOnError(err, fmt.Sprintf("Could not post work with offset: %d", offset))
	}

	w.Write([]byte(fmt.Sprintf(" - 	Read file: '%v' - ElapsedTime: %v - Duration: %v", stopsFilename, sw.ElapsedTime(), swReadFile.ElapsedTime())))
	w.Write([]byte("<br/>"))
}


func LinesIterator(src string) <- chan []byte {
	channel := make(chan []byte)
	go func() {
		utils.ReadCsvFile(src, channel)
		defer close(channel)
	}()
	return channel
}
