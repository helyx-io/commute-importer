package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"github.com/helyx-io/commute-importer/database"
	"github.com/helyx-io/commute-importer/models"
	"github.com/helyx-io/commute-importer/tasks"
	"github.com/helyx-io/commute-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r SQLGTFSRepository) Trips() database.GTFSCreatedModelRepository {
	return SQLTripRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

type SQLTripRepository struct {
	SQLGTFSModelRepository
}

func (s SQLTripRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "trips")
}

func (r SQLTripRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.driver}
	return SQLTripsImportTask{mysqlImportTask}
}

func (s SQLTripRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "trips", params, true)
}

func (s SQLTripRepository) AddIndexesByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    err := s.driver.CreateIndex(schema, "trips", "route_id")
    err = s.driver.CreateIndex(schema, "trips", "service_id")

    return err
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLTripsImportTask struct {
	SQLImportTask
}

func (m SQLTripsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m SQLTripsImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	var offsets = make(map[string]int)

	for i, header := range headers {
		offsets[header] = i
	}

	for i, record := range *rs {
		st[i] = models.TripImportRow{
			recordValueAsInt(record, offsets, "route_id"),
			recordValueAsInt(record, offsets, "service_id"),
			recordValueAsInt(record, offsets, "trip_id"),
			recordValueAsString(record, offsets, "trip_headsign"),
            recordValueAsInt(record, offsets, "direction_id"),
			recordValueAsString(record, offsets, "block_id"),
			recordValueAsString(record, offsets, "shape_id"),
		}
	}

	return st
}

func (m SQLTripsImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 7)

	table := fmt.Sprintf("gtfs_%s.trips", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" route_id," +
		" service_id," +
		" trip_id," +
		" trip_headsign," +
		" direction_id," +
		" block_id," +
		" shape_id" +
		" ) VALUES %s"


	count := 0
    for _, entry := range as {
        i := count * 7
		t := entry.(models.TripImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7)
        } else {
            args = "(?, ?, ?, ?, ?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
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

		count += 1

		if count >= 1024 {
			stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

			_, err = dbSql.Exec(stmt, valueArgs...)
			utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))

			valueStrings = make([]string, 0, len(as))
			valueArgs = make([]interface{}, 0, len(as) * 7)
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
