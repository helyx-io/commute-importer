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
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) CalendarDates() database.GTFSModelRepository {
	return MySQLCalendarDateRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLCalendarDateRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLCalendarDateRepository) RemoveAllByAgencyKey(agencyKey string) (error) {

	table := fmt.Sprintf("%s_calendar_dates", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}

func (r MySQLCalendarDateRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, workPool *workpool.WorkPool, done chan error) workpool.PoolWorker {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, workPool, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLCalendarDatesImportTask{mysqlImportTask}
}

func (s MySQLCalendarDateRepository) CreateTableByAgencyKey(agencyKey string) error {

	tmpTable := fmt.Sprintf("%s_calendar_dates", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", tmpTable))

	ddl, _ := data.Asset("resources/ddl/calendar_dates.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLCalendarDateRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLCalendarDatesImportTask struct {
	MySQLImportTask
}

func (m MySQLCalendarDatesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLCalendarDatesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		serviceId, _ := strconv.Atoi(record[0])
		exceptionType, _ := strconv.Atoi(record[2])

		st[i] = models.CalendarDateImportRow{
			m.AgencyKey,
			serviceId,
			record[1],
			exceptionType,
		}
	}

	return st
}

func (m MySQLCalendarDatesImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 3)

	table := fmt.Sprintf("%s_calendar_dates", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" service_id," +
		" date," +
		" exception_type" +
		" ) VALUES %s"

	count := 0

	for _, entry := range as {
		cd := entry.(models.CalendarDateImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(
			valueArgs,
			cd.ServiceId,
			cd.Date,
			cd.ExceptionType,
		)

		count += 1

		if count >= 4096 {
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
