package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"github.com/helyx-io/gtfs-importer/database"
	"github.com/helyx-io/gtfs-importer/models"
	"github.com/helyx-io/gtfs-importer/tasks"
	"github.com/helyx-io/gtfs-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r SQLGTFSRepository) Routes() database.GTFSCreatedModelRepository {
	return SQLRouteRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

type SQLRouteRepository struct {
	SQLGTFSModelRepository
}

func (s SQLRouteRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "routes")
}


func (r SQLRouteRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.driver}
	return SQLRoutesImportTask{mysqlImportTask}
}

func (s SQLRouteRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "routes", params, true)
}

func (s SQLRouteRepository) AddIndexesByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    err := s.driver.CreateIndex(schema, "routes", "agency_id")

    return err
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLRoutesImportTask struct {
	SQLImportTask
}

func (m SQLRoutesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m SQLRoutesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
        routeId, _ := strconv.Atoi(record[0])
        agencyId, _ := strconv.Atoi(record[1])
		routeType, _ := strconv.Atoi(record[5])

		st[i] = models.RouteImportRow{
            routeId,
            agencyId,
			strings.ToUpper(record[2]),
            strings.ToUpper(record[3]),
			record[4],
			routeType,
			record[6],
            strings.ToUpper(record[7]),
            strings.ToUpper(record[8]),
		}
	}

	return st
}

func (m SQLRoutesImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	table := fmt.Sprintf("gtfs_%s.routes", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" route_id," +
		" agency_id," +
		" route_short_name," +
		" route_long_name," +
		" route_desc," +
		" route_type," +
		" route_url," +
		" route_color," +
		" route_text_color" +
		" ) VALUES %s"

	count := 0

	for _, entry := range as {
        i := count * 9
		r := entry.(models.RouteImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9)
        } else {
            args = "(?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
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

		count += 1

		if count >= 1024 {
			stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

			_, err = dbSql.Exec(stmt, valueArgs...)
			utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))

			valueStrings = make([]string, 0, len(as))
			valueArgs = make([]interface{}, 0, len(as) * 9)
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
