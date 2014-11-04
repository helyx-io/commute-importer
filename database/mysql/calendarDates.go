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

func (r MySQLGTFSRepository) CalendarDates() database.GTFSModelRepository {
	return MySQLCalendarDateRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLCalendarDateRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLCalendarDateRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("calendar_dates").Where("agency_key = ?", agencyKey).Delete(models.CalendarDate{}).Error
}


func (r MySQLCalendarDateRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	return MySQLCalendarDatesImportTask{
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

type MySQLCalendarDatesImportTask struct {
	MySQLImportTask
}

func (m MySQLCalendarDatesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLCalendarDatesImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		serviceId, _ := strconv.Atoi(record[0])
		exceptionType, _ := strconv.Atoi(record[2])

		st[i] = models.CalendarDate{
			m.AgencyKey,
			serviceId,
			record[1],
			exceptionType,
		}
	}

	return st
}

func (m MySQLCalendarDatesImportTask) ImportModels(as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	for _, entry := range as {
		cd := entry.(models.CalendarDate)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			cd.ServiceId,
			cd.Date,
			cd.ExceptionType,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO calendar_dates (" +
			" agency_key," +
			" service_id," +
			" date," +
			" exception_type" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
