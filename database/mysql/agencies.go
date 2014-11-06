package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
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

func (r MySQLGTFSRepository) Agencies() database.GTFSAgencyRepository {
	return MySQLAgencyRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLAgencyRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLAgencyRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("agencies").Where("agency_key = ?", agencyKey).Delete(models.Agency{}).Error
}

func (s MySQLAgencyRepository) FindAll() (*models.Agencies, error) {
	var agencies models.Agencies
	err := s.db.Table("agencies").Find(&agencies).Error

	return &agencies, err
}

func (s MySQLAgencyRepository) FindById(id int) (*models.Agency, error) {
	var agency models.Agency
	err := s.db.Table("agencies").Where("id = ?", id).First(&agency).Error

	return &agency, err
}

func (r MySQLAgencyRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	importTask := tasks.ImportTask{name, agencyKey, lines, workPool}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLAgenciesImportTask{mysqlImportTask}
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLAgenciesImportTask struct {
	MySQLImportTask
}

func (m MySQLAgenciesImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLAgenciesImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		st[i] = models.Agency{
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

func (m MySQLAgenciesImportTask) ImportModels(as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	for _, entry := range as {
		a := entry.(models.Agency)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			a.Id,
			a.Name,
			a.Url,
			a.Timezone,
			a.Lang,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO agencies (" +
			" agency_key," +
			" agency_id," +
			" agency_name," +
			" agency_url," +
			" agency_timezone," +
			" agency_lang" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
