package mysql

import (
	"fmt"
	"strings"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/tasks"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)

////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQL
////////////////////////////////////////////////////////////////////////////////////////////////

func InitDb(maxIdelConns, maxOpenConns int) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8")

	if err != nil {
		return nil, err
	}

	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()

	db.DB().SetMaxIdleConns(maxIdelConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	db.SingularTable(true)

	return &db, nil
}

func CreateMySQLGTFSRepository(db *gorm.DB) (GTFSRepository) {
	return MySQLGTFSRepository{db}
}

type GTFSRepository interface {
	StopTimes() StopTimeRepository
	Stops() StopRepository
}

type MySQLGTFSRepository struct {
	db *gorm.DB
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopTimeRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) StopTimes() StopTimeRepository {
	return MySQLStopTimeRepository{r.db}
}

type StopTimeRepository interface {
	RemoveAllByAgencyKey(agencyKey string) (error)
	CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) (MySQLStopTimesImportTask)
}

type MySQLStopTimeRepository struct {
	db *gorm.DB
}

func (s MySQLStopTimeRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("stop_times").Where("agency_key = ?", agencyKey).Delete(models.StopTime{}).Error
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Stops() StopRepository {
	return MySQLStopRepository{r.db}
}

type StopRepository interface {
	RemoveAllByAgencyKey(agencyKey string) (error)
	CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) (MySQLStopsImportTask)
}

type MySQLStopRepository struct {
	db *gorm.DB
}


func (s MySQLStopRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("stops").Where("agency_key = ?", agencyKey).Delete(models.Stop{}).Error
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopTimesImportTask
////////////////////////////////////////////////////////////////////////////////////////////////

type StopTimesImportTask interface {
	CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) (StopTimesImportTask)
}

type MySQLStopTimesImportTask struct {
	tasks.ImportTask
	db *gorm.DB
}

func (r MySQLStopTimeRepository) CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) (MySQLStopTimesImportTask) {
	return MySQLStopTimesImportTask{
		tasks.ImportTask {
			Name: name,
			Lines: lines,
			WP: workPool,
		},
		r.db,
	}
}

func (m *MySQLStopTimesImportTask) DoWork(workRoutine int) {
	m.InsertStopTimes(stopTimesInserter(m.db, "RATP"));
}

func stopTimesInserter(db *gorm.DB, agencyKey string) (tasks.StopTimesInserter) {

	return func(sts *models.StopTimes) (error) {
		valueStrings := make([]string, 0, len(sts.Records))
		valueArgs := make([]interface{}, 0, len(sts.Records) * 9)

		for _, st := range sts.Records {
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

		return db.Exec(stmt, valueArgs...).Error
	}

}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopsImportTask
////////////////////////////////////////////////////////////////////////////////////////////////

type StopsImportTask interface {
	CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) (StopsImportTask)
}

type MySQLStopsImportTask struct {
	tasks.ImportTask
	db *gorm.DB
}

func (r MySQLStopRepository) CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) (MySQLStopsImportTask) {
	return MySQLStopsImportTask{
		tasks.ImportTask {
			Name: name,
			Lines: lines,
			WP: workPool,
		},
		r.db,
	}
}

func (m *MySQLStopsImportTask) DoWork(workRoutine int) {
	m.InsertStops(stopsInserter(m.db, "RATP"));
}

func stopsInserter(db *gorm.DB, agencyKey string) (tasks.StopsInserter) {

	return func(ss *models.Stops) (error) {
		valueStrings := make([]string, 0, len(ss.Records))
		valueArgs := make([]interface{}, 0, len(ss.Records) * 10)

		for _, s := range ss.Records {
			valueStrings = append(valueStrings, "('" + agencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			valueArgs = append(
				valueArgs,
				s.StopId,
				s.StopCode,
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
			"INSERT INTO stop_times (" +
			" agency_key," +
			" stop_id," +
			" stop_code," +
			" stop_name," +
			" stop_desc," +
			" stop_lat," +
			" stop_lon," +
			" zone_id," +
			" stop_url," +
			" location_type," +
			" parent_station," +
			" ) VALUES %s", strings.Join(valueStrings, ","))

		return db.Exec(stmt, valueArgs...).Error
	}

}
