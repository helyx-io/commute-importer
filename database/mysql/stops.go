package mysql

import (
	"fmt"
	"strings"
	"github.com/akinsella/go-playground/database"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/tasks"
	"github.com/jinzhu/gorm"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Stops() database.GTFSModelRepository {
	return MySQLStopRepository{
		MySQLGTFSModelRepository{r.db},
	}
}

type MySQLStopRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLStopRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("stops").Where("agency_key = ?", agencyKey).Delete(models.Stop{}).Error
}

func (r MySQLStopRepository) CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	return MySQLStopsImportTask{
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

type MySQLStopsImportTask struct {
	MySQLImportTask
}

func (m MySQLStopsImportTask) DoWork(_ int) {
	m.InsertStops(stopsInserter(m.db, "RATP"));
}

func stopsInserter(db *gorm.DB, agencyKey string) tasks.StopsInserter {

	return func(ss *models.Stops) (error) {

		dbSql, err := sql.Open("mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8");

		if err != nil {
			panic(err.Error())
		}

		defer dbSql.Close()

		valueStrings := make([]string, 0, len(*ss))
		valueArgs := make([]interface{}, 0, len(*ss) * 9)

		for _, s := range *ss {
			valueStrings = append(valueStrings, "('" + agencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?, ?)")
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

}
