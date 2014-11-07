package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopTimeRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLStopTimeRepository struct {
	MySQLGTFSModelRepository
}

func (r MySQLGTFSRepository) StopTimes() database.GTFSModelRepository {
	return MySQLStopTimeRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

func (s MySQLStopTimeRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("stop_times").Where("agency_key = ?", agencyKey).Delete(models.StopTime{}).Error
}

func (r MySQLStopTimeRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	importTask := tasks.ImportTask{name, agencyKey, lines, workPool}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLStopTimesImportTask{mysqlImportTask}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopTimesImportTask
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLStopTimesImportTask struct {
	MySQLImportTask
}

func (m MySQLStopTimesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m)
}

func (m MySQLStopTimesImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		stopSequence, _ := strconv.Atoi(record[4])
		pickup_type, _ := strconv.Atoi(record[6])
		drop_off_type, _ := strconv.Atoi(record[7])
		st[i] = models.StopTimeImportRow{
			m.AgencyKey,
			record[0],
			record[1],
			record[2],
			record[3],
			stopSequence,
			record[5],
			pickup_type,
			drop_off_type,
		}
	}

	return st
}

func (m MySQLStopTimesImportTask) ImportModels(sts []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(sts))
	valueArgs := make([]interface{}, 0, len(sts) * 9)

	for _, entry := range sts {
		st := entry.(models.StopTimeImportRow)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			st.TripId,
			st.ArrivalTime,
			st.DepartureTime,
			st.StopId,
			st.StopSequence,
			st.StopHeadSign,
			st.PickupType,
			st.DropOffType,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO stop_times (" +
		" agency_key," +
		" trip_id," +
		" arrival_time," +
		" departure_time," +
		" stop_id," +
		" stop_sequence," +
		" stop_head_sign," +
		" pickup_type," +
		" drop_off_type" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err
}
