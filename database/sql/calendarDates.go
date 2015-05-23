package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"github.com/helyx-io/commute-importer/database"
	"github.com/helyx-io/commute-importer/models"
	"github.com/helyx-io/commute-importer/tasks"
	"github.com/helyx-io/commute-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r SQLGTFSRepository) CalendarDates() database.GTFSCreatedModelRepository {
	return SQLCalendarDateRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

type SQLCalendarDateRepository struct {
	SQLGTFSModelRepository
}

func (s SQLCalendarDateRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "calendar_dates")
}

func (r SQLCalendarDateRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.driver}
	return SQLCalendarDatesImportTask{mysqlImportTask}
}

func (s SQLCalendarDateRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "calendar_dates", params, true)
}

func (s SQLCalendarDateRepository) AddIndexesByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    err := s.driver.CreateIndex(schema, "calendar_dates", "date")

    return err
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLCalendarDatesImportTask struct {
	SQLImportTask
}

func (m SQLCalendarDatesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m SQLCalendarDatesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
        serviceId, _ := strconv.Atoi(record[0])
        exceptionType, _ := strconv.Atoi(record[2])

		st[i] = models.CalendarDateImportRow{
            serviceId,
			record[1],
			exceptionType,
		}
	}

	return st
}

func (m SQLCalendarDatesImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 3)

	table := fmt.Sprintf("gtfs_%s.calendar_dates", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" service_id," +
		" date," +
		" exception_type" +
		" ) VALUES %s"

	count := 0

    for _, entry := range as {
        i := count * 3
		cd := entry.(models.CalendarDateImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d)", i + 1, i + 2, i + 3)
        } else {
            args = "(?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
		valueArgs = append(
			valueArgs,
			cd.ServiceId,
			cd.Date,
			cd.ExceptionType,
		)

		count += 1

		if count >= 1024 {
			stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

			_, err = dbSql.Exec(stmt, valueArgs...)
			utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))

			valueStrings = make([]string, 0, len(as))
			valueArgs = make([]interface{}, 0, len(as) * 3)
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
