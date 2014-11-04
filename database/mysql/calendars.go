package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Calendars() database.GTFSModelRepository {
	return MySQLCalendarRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLCalendarRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLCalendarRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("calendars").Where("agency_key = ?", agencyKey).Delete(models.Calendar{}).Error
}


func (r MySQLCalendarRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	return MySQLCalendarsImportTask{
		MySQLImportTask{
			tasks.ImportTask{
				Name: name,
				AgencyKey: agencyKey,
				Lines: lines,
				WP: workPool,
			},
			r.db,
			r.dbInfos,
		},
	}
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

func(m MySQLCalendarsImportTask) ConvertModels(rs *models.Records) []interface{} {
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

		st[i] = models.Calendar{
			m.AgencyKey,
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

func (m MySQLCalendarsImportTask) ImportModels(as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	for _, entry := range as {
		c := entry.(models.Calendar)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
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
	}

	stmt := fmt.Sprintf(
		"INSERT INTO calendars (" +
			" agency_key," +
			" monday," +
			" tuesday," +
			" wednesday," +
			" thursday," +
			" friday," +
			" saturday," +
			" sunday," +
			" start_date," +
			" end_date" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
