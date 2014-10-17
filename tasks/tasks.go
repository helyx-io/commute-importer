package tasks

import (
	"log"
	"github.com/goinggo/workpool"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/utils"
)

type ImportTask struct {
	Name string
	Lines []byte
	WP *workpool.WorkPool
}

type StopTimesInserter func(sts *models.StopTimes) error

func (it *ImportTask) InsertStopTimes(stopTimesInserter StopTimesInserter) {

	records, err := utils.ParseCsv(it.Lines)

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
