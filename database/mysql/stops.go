package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/tasks"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Stops() database.GTFSCreatedModelRepository {
	return MySQLStopRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLStopRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLStopRepository) RemoveAllByAgencyKey(agencyKey string) (error) {

	table := fmt.Sprintf("`gtfs_%s`.`stops`", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}

func (r MySQLStopRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, workPool *workpool.WorkPool, done chan error) workpool.PoolWorker {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, workPool, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLStopsImportTask{mysqlImportTask}
}

func (s MySQLStopRepository) CreateTableByAgencyKey(agencyKey string) error {

	tmpTable := fmt.Sprintf("`gtfs_%s`.`stops`", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", tmpTable))

	ddl, _ := data.Asset("resources/ddl/stops.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLStopRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
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

func (m MySQLStopsImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	var offsets = make(map[string]int)

	for i, header := range headers {
		offsets[header] = i
	}

	for i, record := range *rs {
		stopLat := recordValueAsInt(record, offsets, "stop_lat")
		stopLon := recordValueAsInt(record, offsets, "stop_lon")
		locationType := recordValueAsInt(record, offsets, "location_type")
		parentStation := recordValueAsInt(record, offsets, "parent_station")
		st[i] = models.StopImportRow{
			m.AgencyKey,
			recordValueAsString(record, offsets, "stop_id"),
			recordValueAsString(record, offsets, "stop_code"),
			recordValueAsString(record, offsets, "stop_name"),
			recordValueAsString(record, offsets, "stop_desc"),
			stopLat,
			stopLon,
			recordValueAsString(record, offsets, "zone_id"),
			recordValueAsString(record, offsets, "stop_url"),
			locationType,
			parentStation,
		}
	}

	return st
}

func recordValueAsString(record []string, offsets map[string]int, key string) string {

	offset, ok := offsets[key]

	if !ok {
		return ""
	}

	return record[offset]
}

func recordValueAsInt(record []string, offsets map[string]int, key string) int {

	offset, ok := offsets[key]

	if !ok {
		return 0
	}

	intValue, _ := strconv.Atoi(record[offset])

	return intValue
}

func (m MySQLStopsImportTask) ImportModels(headers []string, ss []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(ss))
	valueArgs := make([]interface{}, 0, len(ss) * 10)

	table := fmt.Sprintf("`gtfs_%s`.`stops`", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" stop_id," +
		" stop_code," +
		" stop_name," +
		" stop_desc," +
		" stop_lat," +
		" stop_lon," +
		" zone_id," +
		" stop_url," +
		" location_type," +
		" parent_station" +
		" ) VALUES %s"


	var count int = 0
	for _, entry := range ss {
		s := entry.(models.StopImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
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

		count += 1

		if count >= 4096 {
			stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

			_, err = dbSql.Exec(stmt, valueArgs...)
			utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))

			valueStrings = make([]string, 0, len(ss))
			valueArgs = make([]interface{}, 0, len(ss) * 9)
			count = 0
		}
	}

	if count > 0 {
		stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

		_, err = dbSql.Exec(stmt, valueArgs...)
		utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))
	}

	return err

}
