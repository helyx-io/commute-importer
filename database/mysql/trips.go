package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Trips() database.GTFSModelRepository {
	return MySQLTripRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLTripRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLTripRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("trips").Where("agency_key = ?", agencyKey).Delete(models.Trip{}).Error
}


func (r MySQLTripRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	return MySQLTripsImportTask{
		MySQLImportTask{
			tasks.ImportTask{
				Name: name,
				AgencyKey: agencyKey,
				Lines: lines,
				WP: workPool,
			},
			r.db,
			r.dbInfos,
		},
	}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLTripsImportTask struct {
	MySQLImportTask
}

func (m MySQLTripsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLTripsImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		serviceId, _ := strconv.Atoi(record[1])
		directionId, _ := strconv.Atoi(record[4])

		st[i] = models.Trip{
			m.AgencyKey,
			record[0],
			serviceId,
			record[2],
			record[3],
			directionId,
			record[5],
			record[6],
		}
	}

	return st
}

func (m MySQLTripsImportTask) ImportModels(as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	for _, entry := range as {
		t := entry.(models.Trip)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			t.RouteId,
			t.ServiceId,
			t.TripId,
			t.TripHeadsign,
			t.DirectionId,
			t.BlockId,
			t.ShapeId,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO trips (" +
			" agency_key," +
			" route_id," +
			" service_id," +
			" trip_id," +
			" trip_headsign," +
			" direction_id," +
			" block_id," +
			" shape_id" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
