package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"github.com/helyx-io/commute-importer/models"
	"github.com/helyx-io/commute-importer/database"
	"github.com/helyx-io/commute-importer/tasks"
	"github.com/helyx-io/commute-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopTimeRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLStopTimeRepository struct {
	SQLGTFSModelRepository
}

func (r SQLGTFSRepository) StopTimes() database.GTFSCreatedModelRepository {
	return SQLStopTimeRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

func (s SQLStopTimeRepository) RemoveAllByAgencyKey(agencyKey string) error {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "stop_times")
}

func (s SQLStopTimeRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "stop_times", params, true)
}

func (s SQLStopTimeRepository) AddIndexesByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    err := s.driver.CreateIndex(schema, "stop_times", "trip_id")
    err = s.driver.CreateIndex(schema, "stop_times", "stop_id")

    return err
}

func (r SQLStopTimeRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.driver}
	return SQLStopTimesImportTask{mysqlImportTask}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopTimesImportTask
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLStopTimesImportTask struct {
	SQLImportTask
}

func (m SQLStopTimesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m)
}

func (m SQLStopTimesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
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
			recordValueAsInt(record, offsets, "trip_id"),
			recordValueAsString(record, offsets, "arrival_time"),
			recordValueAsString(record, offsets, "departure_time"),
			recordValueAsInt(record, offsets, "stop_id"),
			stopSequence,
			recordValueAsString(record, offsets, "stop_headsign"),
			pickup_type,
			drop_off_type,
		}
	}

	return st
}

func (m SQLStopTimesImportTask) ImportModels(headers []string, sts []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(sts))
	valueArgs := make([]interface{}, 0, len(sts) * 8)

	table := fmt.Sprintf("gtfs_%s.stop_times", m.AgencyKey)

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
        i := count * 8
		st := entry.(models.StopTimeImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8)
        } else {
            args = "(?, ?, ?, ?, ?, ?, ?, ?)"
        }

		// Tmp fix for invalid line
		if st.DepartureTime != "" && st.ArrivalTime != "" {
			valueStrings = append(valueStrings, args)
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
				valueArgs = make([]interface{}, 0, len(sts) * 8)
				count = 0
			}

		}
	}

	if count > 0 {
		stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

		_, err = dbSql.Exec(stmt, valueArgs...)
		utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))
	}

	return err
}
