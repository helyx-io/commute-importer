package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/tasks"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/data"
	_ "github.com/go-sql-driver/mysql"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopTimeRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLStopTimeRepository struct {
	MySQLGTFSModelRepository
}

func (r MySQLGTFSRepository) StopTimes() database.GTFSCreatedModelRepository {
	return MySQLStopTimeRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

func (s MySQLStopTimeRepository) RemoveAllByAgencyKey(agencyKey string) error {

	table := fmt.Sprintf("gtfs_%s.stop_times", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}

func (s MySQLStopTimeRepository) CreateTableByAgencyKey(agencyKey string) error {

	table := fmt.Sprintf("`gtfs_%s`.`stop_times`", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", table))

	ddl, _ := data.Asset("resources/ddl/stop_times.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLStopTimeRepository) AddIndexesByAgencyKey(agencyKey string) error {

	table := fmt.Sprintf("`gtfs_%s`.`stop_times`", agencyKey)

	log.Println(fmt.Sprintf("Creating indexes: '%s'", table))

	ddl, _ := data.Asset("resources/ddl/stop_times_indexes.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (r MySQLStopTimeRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
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

func (m MySQLStopTimesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	var offsets = make(map[string]int)

	for i, header := range headers {
		offsets[header] = i
	}

	for i, record := range *rs {
		stopSequence := recordValueAsInt(record, offsets, "stop_sequence")
		pickup_type := recordValueAsInt(record, offsets, "pickup_type")
		drop_off_type := recordValueAsInt(record, offsets, "drop_off_type")
		st[i] = models.StopTimeImportRow{
			m.AgencyKey,
			recordValueAsString(record, offsets, "trip_id"),
			recordValueAsString(record, offsets, "arrival_time"),
			recordValueAsString(record, offsets, "departure_time"),
			recordValueAsString(record, offsets, "stop_id"),
			stopSequence,
			recordValueAsString(record, offsets, "stop_headsign"),
			pickup_type,
			drop_off_type,
		}
	}

	return st
}

func (m MySQLStopTimesImportTask) ImportModels(headers []string, sts []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(sts))
	valueArgs := make([]interface{}, 0, len(sts) * 8)

	table := fmt.Sprintf("`gtfs_%s`.`stop_times`", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" trip_id," +
		" arrival_time," +
		" departure_time," +
		" stop_id," +
		" stop_sequence," +
		" stop_head_sign," +
		" pickup_type," +
		" drop_off_type" +
		" ) VALUES %s"

	var count int = 0
	for _, entry := range sts {
		st := entry.(models.StopTimeImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?)")
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

		count += 1

		if count >= 1024 {
			stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

			_, err = dbSql.Exec(stmt, valueArgs...)
			utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))

			valueStrings = make([]string, 0, len(sts))
			valueArgs = make([]interface{}, 0, len(sts) * 9)
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
