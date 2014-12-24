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
	CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, workPool *workpool.WorkPool, done chan error) workpool.PoolWorker
	CreateTableByAgencyKey(agencyKey string) error
	AddIndexesByAgencyKey(agencyKey string) error
}

type GTFSAgencyRepository interface {
	GTFSModelRepository
}

type GTFSCalendarRepository interface {
	GTFSModelRepository
}

type GTFSCalendarDateRepository interface {
	GTFSModelRepository
}

type GTFSRouteRepository interface {
	GTFSModelRepository
}

type GTFSTripRepository interface {
	GTFSModelRepository
}

type GTFSTransferRepository interface {
	GTFSModelRepository
}

type GTFSStopRepository interface {
	GTFSModelRepository
}

type GTFSStopTimeRepository interface {
	GTFSModelRepository
}
