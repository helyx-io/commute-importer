package controller

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
	"path/filepath"
	"github.com/gorilla/mux"
//	"github.com/streadway/amqp"
	"github.com/fatih/stopwatch"
	"github.com/akinsella/go-playground/models"
	"github.com/goinggo/workpool"
	"gopkg.in/mgo.v2"
)

type ImportController struct { }

type ImportTask struct {
	Name string
	Lines []byte
	WP *workpool.WorkPool
}

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

/* // AMQP

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		"chunks",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	failOnError(err, "Failed to declare an exchange")

	chunksQueue, err := channel.QueueDeclare(
		"chunks", // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = channel.QueueBind(chunksQueue.Name, " chunks", "chunks", false, nil)

	failOnError(err, fmt.Sprintf("Failed to bind to queue with name: %v", chunksQueue.Name))
*/

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

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	writtenBytes, err := io.Copy(out, resp.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf(" - Downloaded zip file: '%v' - %v bytes - ElapsedTime: %v", zipFilename, writtenBytes, sw.ElapsedTime())))
	w.Write([]byte("<br/>"))

	w.Write([]byte(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename)))
	w.Write([]byte("<br/>"))

	swZip := stopwatch.Start(0)

	err = Unzip(zipFilename, folderFilename)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v' - ElapsedTime: %v - Duration: %v", zipFilename, folderFilename, sw.ElapsedTime(), swZip.ElapsedTime())))
	w.Write([]byte("<br/>"))

	stopsFilename := fmt.Sprintf("%v/stop_times.txt", folderFilename)

	w.Write([]byte(fmt.Sprintf(" - Reading file: '%v'", stopsFilename)))
	w.Write([]byte("<br/>"))

	swReadFile := stopwatch.Start(0)

	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
			http.Error(w, err.Error(), 500)
			return
		}
	}()


	workPool := workpool.New(32, 10000)

	offset := 0
	for lines := range LinesIterator(stopsFilename) {
		offset++
		task := ImportTask{
			Name: fmt.Sprintf("ChunkImport-%d", offset),
			Lines: lines,
			WP: workPool,
		}

		err := workPool.PostWork("import", &task)

		if err != nil {
			log.Println("Could not post work:", err)
			panic(err)
		}

	}

	w.Write([]byte(fmt.Sprintf(" - 	Read file: '%v' - ElapsedTime: %v - Duration: %v", stopsFilename, sw.ElapsedTime(), swReadFile.ElapsedTime())))
	w.Write([]byte("<br/>"))
}

func (it *ImportTask) DoWork(workRoutine int) {

	lines, err := ParseCsv(it.Lines)

	if err != nil {
		log.Println("Could parse CSV File:", err)
		panic(err)
	}

	_, err = BulkInsertRecords(lines)

	if err != nil {
		log.Println("Could not insert records in database:", err)
		panic(err)
	}

	log.Println(it.Name)
}

func _map(records [][]string) []models.StopTime {
	var stopTimes = make([]models.StopTime, len(records))

	for i, record := range records {
//		fmt.Println("Index:", i, "Record:", record)
		stopTimes[i] = models.StopTime{ "RATP", record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7] }
	}

	return stopTimes
}

func ParseCsv(b []byte) ([][]string, error) {
	r := bytes.NewReader(b)
	reader := csv.NewReader(r)
	records := make([][]string, 0)

	var err error

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err, ok := err.(*csv.ParseError); ok {
			if err.Err != csv.ErrFieldCount {
				fmt.Println(fmt.Sprintf("%#v", err))
				log.Println("2 - Error on line read:", err, "line:", record)
				panic(err)
			}
		} else if err != nil {
			fmt.Println(fmt.Sprintf("%#v", err))
			log.Println("3 - Error on line read:", err, "line:", record)
			break;
		}

//		fmt.Println("Records:", len(records))
		records = append(records, record)
	}

	return records, err
}

func BulkInsertRecords(records [][]string) (*mgo.BulkResult, error)  {

	mSession := getSession()

	defer mSession.Close()

	c := mSession.DB("gtfs").C("stop_times")

	bulk := c.Bulk()

	for _, record := range _map(records) {
		bulk.Insert(record)
	}

	bulkResult, err := bulk.Run()

	return bulkResult, err
}

func LinesIterator(src string) <- chan []byte {
	channel := make(chan []byte)
	go func() {
		ReadFile(src, channel)
		defer close(channel)
	}()
	return channel
}

func ReadFile(src string, channel chan []byte) {

	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	b := []byte{}
	chunk := 0
	i := 0
	for {
		l, err := r.ReadBytes('\n')

		if err == io.EOF {
			break;
		}

		if err != nil {
			panic(err)
		}

		if len(l) == 0 {
			break;
		}

		b = append(b, l...)
		b = append(b, '\n')

		i++

		if len(b) >= 128000 {
			chunk++
//			fmt.Println("Chunk Index: ", chunk, "Number of lines :", i)
			channel <- b
			i = 0
			b = []byte{}
		}

	}
	if len(b) > 0 {
		chunk++
//		fmt.Println("Chunk Index: ", chunk, "Number of lines :", i)
		channel <- b
	}
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()

		if err != nil {
			return err
		}

		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
