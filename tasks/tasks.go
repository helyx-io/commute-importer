package tasks

import (
	"log"
	"fmt"
	"github.com/helyx-io/gtfs-importer/models"
)

type Task interface {
	DoWork(workId int)
}

type ImportTask struct {
	Name string
	JobIndex int
	FileName string
	AgencyKey string
	Headers []string
	Lines []byte
	Done chan error
}

type ModelConverter interface {
	ConvertModels(headers []string, records *models.Records) []interface{}
}

type ModelImporter interface {
	ImportModels(headers []string, models []interface{}) error
}

func NewImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) ImportTask {
	return ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
}


func (it ImportTask) ImportCsv(converter ModelConverter, importer ModelImporter) {

	records, err := models.ParseCsv(it.Lines)

	if err != nil {
		log.Println(fmt.Sprintf("[%s][%d] Could parse CSV File:", it.AgencyKey, it.FileName), err)
		it.Done <- err
	}

	models := converter.ConvertModels(it.Headers, records)
	err = importer.ImportModels(it.Headers, models)

	if err != nil {
		log.Println(fmt.Sprintf("[%s][%d] Could not insert records in database:", it.AgencyKey, it.JobIndex), err)
		it.Done <- err
	}

	log.Println(fmt.Sprintf("[%s][%d] Sending Task Done Event", it.AgencyKey, it.JobIndex))
	it.Done <- nil
}
