package mysql

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
	"github.com/helyx-io/gtfs-importer/data"
	_ "github.com/go-sql-driver/mysql"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r MySQLGTFSRepository) Agencies() database.GTFSCreatedModelRepository {
	return MySQLAgencyRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLAgencyRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLAgencyRepository) RemoveAllByAgencyKey(agencyKey string) (error) {

	table := fmt.Sprintf("`gtfs_%s`.`agencies`", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}

func (r MySQLAgencyRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLAgenciesImportTask{mysqlImportTask}
}

func (s MySQLAgencyRepository) CreateTableByAgencyKey(agencyKey string) error {

	tmpTable := fmt.Sprintf("`gtfs_%s`.`agencies`", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", tmpTable))

	ddl, _ := data.Asset("resources/ddl/agencies.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLAgencyRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
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

func(m MySQLAgenciesImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
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

func (m MySQLAgenciesImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 5)

	for _, entry := range as {
		a := entry.(models.AgencyImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			a.Id,
			a.Name,
			a.Url,
			a.Timezone,
			a.Lang,
		)
	}

	table := fmt.Sprintf("`gtfs_%s`.`agencies`", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	stmt := fmt.Sprintf(
		"INSERT INTO " + table + " (" +
			" agency_id," +
			" agency_name," +
			" agency_url," +
			" agency_timezone," +
			" agency_lang" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
