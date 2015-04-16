package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"github.com/helyx-io/gtfs-importer/database"
	"github.com/helyx-io/gtfs-importer/models"
	"github.com/helyx-io/gtfs-importer/tasks"
	"github.com/helyx-io/gtfs-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r SQLGTFSRepository) Stops() database.GTFSCreatedModelRepository {
	return SQLStopRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

type SQLStopRepository struct {
	SQLGTFSModelRepository
}

func (s SQLStopRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "stops")
}

func (r SQLStopRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.driver}
	return SQLStopsImportTask{mysqlImportTask}
}

func (s SQLStopRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "stops", params, true)
}

func (s SQLStopRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopsImportTask
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLStopsImportTask struct {
	SQLImportTask
}

func (m SQLStopsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m)
}

func (m SQLStopsImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	var offsets = make(map[string]int)

	for i, header := range headers {
		offsets[header] = i
	}

	for i, record := range *rs {
		stopLat := recordValueAsFloat(record, offsets, "stop_lat")
		stopLon := recordValueAsFloat(record, offsets, "stop_lon")
		locationType := recordValueAsInt(record, offsets, "location_type")
		parentStation := recordValueAsInt(record, offsets, "parent_station")
		st[i] = models.StopImportRow{
			recordValueAsInt(record, offsets, "stop_id"),
			recordValueAsString(record, offsets, "stop_code"),
            strings.ToUpper(recordValueAsString(record, offsets, "stop_name")),
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

func (m SQLStopsImportTask) ImportModels(headers []string, ss []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(ss))
	valueArgs := make([]interface{}, 0, len(ss) * 10)

	table := fmt.Sprintf("gtfs_%s.stops", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" stop_id," +
		" stop_code," +
		" stop_name," +
		" stop_desc," +
		" stop_lat," +
		" stop_lon," +
		" stop_geo," +
		" zone_id," +
		" stop_url," +
		" location_type," +
		" parent_station" +
		" ) VALUES %s"


	var count int = 0
    for _, entry := range ss {
        i := count * 11
		s := entry.(models.StopImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, ST_GeomFromText($%d), $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9, i + 10, i + 11)
        } else {
            args = "(?, ?, ?, ?, ?, ?, GeomFromText(?), ?, ?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
		valueArgs = append(
			valueArgs,
			s.StopId,
			s.StopCode,
			s.StopName,
			s.StopDesc,
			s.StopLat,
			s.StopLon,
			fmt.Sprintf("POINT(%f %f)", s.StopLat, s.StopLon),
			s.ZoneId,
			s.StopUrl,
			s.LocationType,
			s.ParentStation,
		)

		count += 1

		if count >= 1024 {
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
