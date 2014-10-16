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

func (it *ImportTask) InsertStopTimes(insertStopTimes func(sts *models.StopTimes) error) {

	records, err := utils.ParseCsv(it.Lines)

	if err != nil {
		log.Println("Could parse CSV File:", err)
		panic(err)
	}

	stopTimes := records.MapToStopTimes()
	err = insertStopTimes(&stopTimes)

	if err != nil {
		log.Println("Could not insert records in database:", err)
		panic(err)
	}

	log.Println(it.Name)
}
