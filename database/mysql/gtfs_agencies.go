package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/models"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) GtfsAgencies() database.GTFSModelRepository {
	return MySQLGtfsAgencyRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLGtfsAgencyRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLGtfsAgencyRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Exec("DELETE FROM `gtfs`.`agencies` where agency_key=?", agencyKey).Error
}

func (r MySQLGtfsAgencyRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, workPool *workpool.WorkPool, done chan error) workpool.PoolWorker {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, workPool, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLAgenciesImportTask{mysqlImportTask}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLGtfsAgenciesImportTask struct {
	MySQLImportTask
}

func (m MySQLGtfsAgenciesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLGtfsAgenciesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		st[i] = models.AgencyImportRow{
			m.AgencyKey,
			record[0],
			record[1],
			record[2],
			record[3],
			record[4],
		}
	}

	return st
}

func (m MySQLGtfsAgenciesImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 5)

	for _, entry := range as {
		a := entry.(models.AgencyImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			m.AgencyKey,
			a.AgencyId,
			a.Name,
			a.Url,
			a.Timezone,
			a.Lang,
		)
	}

	table := fmt.Sprintf("`gtfs`.`agencies`", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	stmt := fmt.Sprintf(
		"INSERT INTO " + table + " (" +
		" agency_key, " +
		" agency_id," +
		" agency_name," +
		" agency_url," +
		" agency_timezone," +
		" agency_lang" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
