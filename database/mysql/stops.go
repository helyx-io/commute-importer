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

func (r MySQLGTFSRepository) Stops() database.GTFSModelRepository {
	return MySQLStopRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLStopRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLStopRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("stops").Where("agency_key = ?", agencyKey).Delete(models.Stop{}).Error
}

func (r MySQLStopRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	importTask := tasks.ImportTask{name, agencyKey, lines, workPool}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLStopsImportTask{mysqlImportTask}
}

func (s MySQLStopRepository) FindAll() (*models.Stops, error) {
	var stops models.Stops
	err := s.db.Table("stops").Limit(1000).Find(&stops).Error

	return &stops, err
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopsImportTask
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLStopsImportTask struct {
	MySQLImportTask
}

func (m MySQLStopsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m)
}

func (m MySQLStopsImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		stopLat, _ := strconv.Atoi(record[3])
		stopLon, _ := strconv.Atoi(record[4])
		locationType, _ := strconv.Atoi(record[7])
		parentStation, _ := strconv.Atoi(record[8])
		st[i] = models.StopImportRow{
			m.AgencyKey,
			record[0],
			record[1],
			record[2],
			stopLat,
			stopLon,
			record[5],
			record[6],
			locationType,
			parentStation,
		}
	}

	return st
}

func (m MySQLStopsImportTask) ImportModels(ss []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(ss))
	valueArgs := make([]interface{}, 0, len(ss) * 9)

	for _, entry := range ss {
		s := entry.(models.Stop)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			s.StopId,
			s.StopName,
			s.StopDesc,
			s.StopLat,
			s.StopLon,
			s.ZoneId,
			s.StopUrl,
			s.LocationType,
			s.ParentStation,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO stops (" +
		" agency_key," +
		" stop_id," +
		" stop_name," +
		" stop_desc," +
		" stop_lat," +
		" stop_lon," +
		" zone_id," +
		" stop_url," +
		" location_type," +
		" parent_station" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
