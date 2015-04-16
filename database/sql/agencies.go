package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
    "strconv"
	"strings"
	"github.com/helyx-io/gtfs-importer/database"
	"github.com/helyx-io/gtfs-importer/models"
	"github.com/helyx-io/gtfs-importer/tasks"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r SQLGTFSRepository) Agencies() database.GTFSCreatedModelRepository {
	return SQLAgencyRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

type SQLAgencyRepository struct {
	SQLGTFSModelRepository
}

func (s SQLAgencyRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "agencies")
}

func (r SQLAgencyRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
    importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
    mysqlImportTask := SQLImportTask{importTask, r.driver}
    return SQLAgenciesImportTask{mysqlImportTask}
}

func (s SQLAgencyRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "agencies", params, true)
}

func (s SQLAgencyRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLAgenciesImportTask struct {
	SQLImportTask
}

func (m SQLAgenciesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m SQLAgenciesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {

        id, _ := strconv.Atoi(record[0])

        st[i] = models.AgencyImportRow{
			m.AgencyKey,
            id,
			record[1],
			record[2],
			record[3],
			record[4],
		}
	}

	return st
}

func (m SQLAgenciesImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 5)


	for index, entry := range as {
        i := index * 5
		a := entry.(models.AgencyImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4, i + 5)
        } else {
            args = "(?, ?, ?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
		valueArgs = append(
			valueArgs,
			a.Id,
			a.Name,
			a.Url,
			a.Timezone,
			a.Lang,
		)
	}

	table := fmt.Sprintf("gtfs_%s.agencies", m.AgencyKey)

	log.Printf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table)

	stmt := fmt.Sprintf(
		"INSERT INTO %s (" +
			" agency_id," +
			" agency_name," +
			" agency_url," +
			" agency_timezone," +
			" agency_lang" +
		" ) VALUES %s", table, strings.Join(valueStrings, ","))

	result, err := dbSql.Exec(stmt, valueArgs...)

    log.Printf("*** agency - Error: %v - Result: %v, Values: %v", err, result, valueArgs)

	return err

}
