package tasks

import (
	"log"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/goinggo/workpool"
)

type ImportTask struct {
	Name string
	AgencyKey string
	Lines []byte
	WP *workpool.WorkPool
}

type ModelConverter interface {
	ConvertModels(records *models.Records) []interface{}
}

type ModelImporter interface {
	ImportModels(models []interface{}) error
}

func (it ImportTask) ImportCsv(converter ModelConverter, importer ModelImporter) {

	records, err := models.ParseCsv(it.Lines)

	if err != nil {
		log.Println("Could parse CSV File:", err)
		panic(err)
	}

	models := converter.ConvertModels(records)
	err = importer.ImportModels(models)

	if err != nil {
		log.Println("Could not insert records in database:", err)
		panic(err)
	}

	log.Println(it.Name)
}
