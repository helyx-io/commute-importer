package database

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSRepository interface {
	CreateSchema(agencyKey string) error
	GtfsAgencies() GTFSModelRepository
	Agencies()	GTFSCreatedModelRepository
	CalendarDates() GTFSCreatedModelRepository
	Calendars() GTFSCreatedModelRepository
	Routes() GTFSCreatedModelRepository
	Stops() GTFSCreatedModelRepository
	StopTimes() GTFSCreatedModelRepository
	Transfers() GTFSCreatedModelRepository
	Trips() GTFSCreatedModelRepository
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
	CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, workPool *workpool.WorkPool, done chan error) workpool.PoolWorker
}

type GTFSCreatedModelRepository interface {
	GTFSModelRepository
	CreateTableByAgencyKey(agencyKey string) error
	AddIndexesByAgencyKey(agencyKey string) error
}

type GTFSAgencyRepository interface {
	GTFSCreatedModelRepository
}

type GTFSGtfsAgencyRepository interface {
	GTFSModelRepository
}

type GTFSCalendarRepository interface {
	GTFSCreatedModelRepository
}

type GTFSCalendarDateRepository interface {
	GTFSCreatedModelRepository
}

type GTFSRouteRepository interface {
	GTFSCreatedModelRepository
}

type GTFSTripRepository interface {
	GTFSCreatedModelRepository
}

type GTFSTransferRepository interface {
	GTFSCreatedModelRepository
}

type GTFSStopRepository interface {
	GTFSCreatedModelRepository
}

type GTFSStopTimeRepository interface {
	GTFSCreatedModelRepository
}
