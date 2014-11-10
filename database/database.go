package database

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/goinggo/workpool"
	"github.com/helyx-io/gtfs-playground/models"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSRepository interface {
	Agencies() GTFSAgencyRepository
	CalendarDates() GTFSModelRepository
	Calendars() GTFSModelRepository
	Routes() GTFSModelRepository
	Stops() GTFSModelRepository
	StopTimes() GTFSModelRepository
	Transfers() GTFSModelRepository
	Trips() GTFSModelRepository
}

type DBConnectInfos struct {
	Dialect string
	URL string
	MaxIdelConns int
	MaxOpenConns int
}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Interfaces
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSModelRepository interface {
	RemoveAllByAgencyKey(agencyKey string) error
	CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker
}

type GTFSAgencyRepository interface {
	GTFSModelRepository
	FindAll() (*models.Agencies, error)
	FindById(id int) (*models.Agency, error)
}

type GTFSCalendarRepository interface {
	GTFSModelRepository
	FindAll() (*models.Calendars, error)
}

type GTFSCalendarDateRepository interface {
	GTFSModelRepository
	FindAll() (*models.CalendarDates, error)
}

type GTFSRouteRepository interface {
	GTFSModelRepository
	FindAll() (*models.Routes, error)
}

type GTFSTripRepository interface {
	GTFSModelRepository
	FindAll() (*models.Trips, error)
}

type GTFSTransferRepository interface {
	GTFSModelRepository
	FindAll() (*models.Transfers, error)
}

type GTFSStopRepository interface {
	GTFSModelRepository
	FindAll() (*models.Stops, error)
}

type GTFSStopTimeRepository interface {
	GTFSModelRepository
	FindAll() (*models.StopTimes, error)
}
