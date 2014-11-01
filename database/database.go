package database

import (
	"github.com/goinggo/workpool"
	"github.com/akinsella/go-playground/models"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// GTFS
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSRepository interface {
	Agencies() GTFSAgencyRepository
	StopTimes() GTFSModelRepository
	Stops() GTFSModelRepository
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// GTFS Model
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSModelRepository interface {
	RemoveAllByAgencyKey(agencyKey string) error
	CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker
}

type GTFSAgencyRepository interface {
	GTFSModelRepository
	FindAll() (*models.Agencies, error)
	FindByKey(agencyKey string) (*models.Agency, error)
}
