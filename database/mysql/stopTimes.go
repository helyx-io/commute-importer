package mysql

import (
	"fmt"
	"strings"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/database"
	"github.com/akinsella/go-playground/tasks"
	"github.com/jinzhu/gorm"
	"database/sql"
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
		MySQLGTFSModelRepository{r.db},
	}
}

func (s MySQLStopTimeRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("stop_times").Where("agency_key = ?", agencyKey).Delete(models.StopTime{}).Error
}

func (r MySQLStopTimeRepository) CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	return MySQLStopTimesImportTask{
		MySQLImportTask{
			tasks.ImportTask{
				Name: name,
				Lines: lines,
				WP: workPool,
			},
			r.db,
		},
	}
}

type MySQLStopTimesImportTask struct {
	MySQLImportTask
}

func (m MySQLStopTimesImportTask) DoWork(_ int) {
	m.InsertStopTimes(stopTimesInserter(m.db, "RATP"));
}

func stopTimesInserter(db *gorm.DB, agencyKey string) tasks.StopTimesInserter {

	return func(sts *models.StopTimes) (error) {

		dbSql, err := sql.Open("mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8");

		if err != nil {
			panic(err.Error())
		}

		defer dbSql.Close()

		valueStrings := make([]string, 0, len(*sts))
		valueArgs := make([]interface{}, 0, len(*sts) * 9)

		for _, st := range *sts {
			valueStrings = append(valueStrings, "('" + agencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?)")
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

}
