package tasks

import (
	"log"
	"github.com/goinggo/workpool"
	"github.com/akinsella/go-playground/models"
)


type ImportTask struct {
	Name string
	Lines []byte
	WP *workpool.WorkPool
}


type StopTimesInserter func(sts *models.StopTimes) error

func (it *ImportTask) InsertStopTimes(stopTimesInserter StopTimesInserter) {

	records, err := models.ParseCsv(it.Lines)

	if err != nil {
		log.Println("Could parse CSV File:", err)
		panic(err)
	}

	stopTimes := records.MapToStopTimes()
	err = stopTimesInserter(&stopTimes)

	if err != nil {
		log.Println("Could not insert records in database:", err)
		panic(err)
	}

	log.Println(it.Name)
}


type StopsInserter func(sts *models.Stops) error

func (it *ImportTask) InsertStops(stopsInserter StopsInserter) {

	records, err := models.ParseCsv(it.Lines)

	if err != nil {
		log.Println("Could parse CSV File:", err)
		panic(err)
	}

	stops := records.MapToStops()
	err = stopsInserter(&stops)

	if err != nil {
		log.Println("Could not insert records in database:", err)
		panic(err)
	}

	log.Println(it.Name)
}
