package database

import (
	"github.com/goinggo/workpool"
	"github.com/helyx-io/gtfs-playground/models"
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
	CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker
}

type GTFSAgencyRepository interface {
	GTFSModelRepository
	FindAll() (*models.Agencies, error)
	FindByKey(agencyKey string) (*models.Agency, error)
}
