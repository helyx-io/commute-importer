package database

import (
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// GTFS
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSRepository interface {
	StopTimes() GTFSModelRepository
	Stops() GTFSModelRepository
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// GTFS Model
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSModelRepository interface {
	RemoveAllByAgencyKey(agencyKey string) (error)
	CreateImportTask(name string, lines *[]byte, workPool *workpool.WorkPool) workpool.PoolWorker
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Stops
////////////////////////////////////////////////////////////////////////////////////////////////

type StopRepository interface {
	GTFSModelRepository
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// StopTimes
////////////////////////////////////////////////////////////////////////////////////////////////

type StopTimeRepository interface {
	GTFSModelRepository
}

