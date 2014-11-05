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

func (r MySQLGTFSRepository) Routes() database.GTFSModelRepository {
	return MySQLRouteRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLRouteRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLRouteRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("routes").Where("agency_key = ?", agencyKey).Delete(models.Route{}).Error
}


func (r MySQLRouteRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	importTask := tasks.ImportTask{name, agencyKey, lines, workPool}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLRoutesImportTask{mysqlImportTask}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLRoutesImportTask struct {
	MySQLImportTask
}

func (m MySQLRoutesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLRoutesImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		routeType, _ := strconv.Atoi(record[5])

		st[i] = models.Route{
			m.AgencyKey,
			record[0],
			record[1],
			record[2],
			record[3],
			record[4],
			routeType,
			record[6],
			record[7],
			record[8],
		}
	}

	return st
}

func (m MySQLRoutesImportTask) ImportModels(as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	for _, entry := range as {
		r := entry.(models.Route)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			r.RouteId,
			r.AgencyId,
			r.RouteShortName,
			r.RouteLongName,
			r.RouteDesc,
			r.RouteType,
			r.RouteUrl,
			r.RouteColor,
			r.RouteTextColor,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO routes (" +
			" agency_key," +
			" route_id," +
			" agency_id," +
			" route_short_name," +
			" route_long_name," +
			" route_desc," +
			" route_type," +
			" route_url," +
			" route_color," +
			" route_text_color" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
