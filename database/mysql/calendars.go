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
	"github.com/helyx-io/gtfs-playground/data"
	"github.com/helyx-io/gtfs-playground/utils"
	_ "github.com/go-sql-driver/mysql"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Calendars() database.GTFSCreatedModelRepository {
	return MySQLCalendarRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLCalendarRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLCalendarRepository) RemoveAllByAgencyKey(agencyKey string) (error) {

	table := fmt.Sprintf("`gtfs_%s`.`calendars`", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}


func (r MySQLCalendarRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLCalendarsImportTask{mysqlImportTask}
}

func (s MySQLCalendarRepository) CreateTableByAgencyKey(agencyKey string) error {

	tmpTable := fmt.Sprintf("`gtfs_%s`.`calendars`", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", tmpTable))

	ddl, _ := data.Asset("resources/ddl/calendars.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLCalendarRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLCalendarsImportTask struct {
	MySQLImportTask
}

func (m MySQLCalendarsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLCalendarsImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
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

func (m MySQLCalendarsImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 10)


	table := fmt.Sprintf("`gtfs_%s`.`calendars`", m.AgencyKey)

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
		c := entry.(models.CalendarImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
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
