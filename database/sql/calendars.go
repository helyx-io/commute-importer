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

func (r SQLGTFSRepository) Calendars() database.GTFSCreatedModelRepository {
	return SQLCalendarRepository{
		SQLGTFSModelRepository{r.driver},
	}
}

type SQLCalendarRepository struct {
	SQLGTFSModelRepository
}

func (s SQLCalendarRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.DropTable(schema, "calendars")
}


func (r SQLCalendarRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.driver}
	return SQLCalendarsImportTask{mysqlImportTask}
}

func (s SQLCalendarRepository) CreateTableByAgencyKey(agencyKey string, params map[string]interface{}) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return s.driver.CreateTable(schema, "calendars", params, true)
}

func (s SQLCalendarRepository) AddIndexesByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    err := s.driver.CreateIndex(schema, "calendars", "start_date")
    err = s.driver.CreateIndex(schema, "calendars", "end_date")
    err = s.driver.CreateIndex(schema, "calendars", "monday")
    err = s.driver.CreateIndex(schema, "calendars", "tuesday")
    err = s.driver.CreateIndex(schema, "calendars", "wednesday")
    err = s.driver.CreateIndex(schema, "calendars", "thursday")
    err = s.driver.CreateIndex(schema, "calendars", "friday")
    err = s.driver.CreateIndex(schema, "calendars", "saturday")
    err = s.driver.CreateIndex(schema, "calendars", "sunday")

    return err
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLCalendarsImportTask struct {
	SQLImportTask
}

func (m SQLCalendarsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m SQLCalendarsImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
        serviceId, _ := strconv.Atoi(record[0])
		monday, _ := strconv.Atoi(record[1])
		tuesday, _ := strconv.Atoi(record[2])
		wednesday, _ := strconv.Atoi(record[3])
		thursday, _ := strconv.Atoi(record[4])
		friday, _ := strconv.Atoi(record[5])
		saturday, _ := strconv.Atoi(record[6])
		sunday, _ := strconv.Atoi(record[7])

		st[i] = models.CalendarImportRow{
            serviceId,
			monday,
			tuesday,
			wednesday,
			thursday,
			friday,
			saturday,
			sunday,
			record[8],
			record[9],
		}
	}

	return st
}

func (m SQLCalendarsImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 10)


	table := fmt.Sprintf("gtfs_%s.calendars", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" service_id," +
		" monday," +
		" tuesday," +
		" wednesday," +
		" thursday," +
		" friday," +
		" saturday," +
		" sunday," +
		" start_date," +
		" end_date" +
		" ) VALUES %s"

	count := 0

	for _, entry := range as {
        i := count * 10
		c := entry.(models.CalendarImportRow)

        var args string
        if m.driver.ConnectInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9, i + 10)
        } else {
            args = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
		valueArgs = append(
			valueArgs,
			c.ServiceId,
			c.Monday,
			c.Tuesday,
			c.Wednesday,
			c.Thursday,
			c.Friday,
			c.Saturday,
			c.Sunday,
			c.StartDate,
			c.EndDate,
		)

		count += 1

		if count >= 1024 {
			stmt := fmt.Sprintf(query, strings.Join(valueStrings, ","))

			_, err = dbSql.Exec(stmt, valueArgs...)
			utils.FailOnError(err, fmt.Sprintf("Could not insert into table with name: '%s'", table))

			valueStrings = make([]string, 0, len(as))
			valueArgs = make([]interface{}, 0, len(as) * 10)
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
