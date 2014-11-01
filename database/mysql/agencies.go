package mysql

import (
	"fmt"
	"strings"
	"github.com/akinsella/go-playground/database"
	"github.com/akinsella/go-playground/models"
	"github.com/akinsella/go-playground/tasks"
	"github.com/jinzhu/gorm"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Agencies() database.GTFSAgencyRepository {
	return MySQLAgencyRepository{
		MySQLGTFSModelRepository{r.db},
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

func (s MySQLAgencyRepository) FindByKey(agencyKey string) (*models.Agency, error) {
	var agency models.Agency
	err := s.db.Table("agencies").Where("agency_key = ?", agencyKey).First(&agency).Error

	return &agency, err
}

func (r MySQLAgencyRepository) CreateImportTask(name string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	return MySQLAgenciesImportTask{
		MySQLImportTask{
			tasks.ImportTask{
				Name: name,
				Lines: lines,
				WP: workPool,
			},
			r.db,
		},
	}
}

type MySQLAgenciesImportTask struct {
	MySQLImportTask
}

func (m MySQLAgenciesImportTask) DoWork(_ int) {
	m.InsertAgencies(agenciesInserter(m.db, "RATP"));
}

func agenciesInserter(db *gorm.DB, agencyKey string) tasks.AgenciesInserter {

	return func(as *models.Agencies) (error) {

		dbSql, err := sql.Open("mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8");

		if err != nil {
			panic(err.Error())
		}

		defer dbSql.Close()

		valueStrings := make([]string, 0, len(*as))
		valueArgs := make([]interface{}, 0, len(*as) * 9)

		for _, a := range *as {
			valueStrings = append(valueStrings, "('" + agencyKey + "', ?, ?, ?, ?, ?)")
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

}




























