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

func (r MySQLGTFSRepository) Transfers() database.GTFSCreatedModelRepository {
	return MySQLTransferRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLTransferRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLTransferRepository) RemoveAllByAgencyKey(agencyKey string) (error) {

	table := fmt.Sprintf("`gtfs_%s`.`transfers`", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}


func (r MySQLTransferRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLTransfersImportTask{mysqlImportTask}
}

func (s MySQLTransferRepository) CreateTableByAgencyKey(agencyKey string) error {

	tmpTable := fmt.Sprintf("`gtfs_%s`.`tranfers`", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", tmpTable))

	ddl, _ := data.Asset("resources/ddl/transfers.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLTransferRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLTransfersImportTask struct {
	MySQLImportTask
}

func (m MySQLTransfersImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLTransfersImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
        fromStopId, _ := strconv.Atoi(record[0])
        toStopId, _ := strconv.Atoi(record[1])
        transferType, _ := strconv.Atoi(record[2])
		minTransferType, _ := strconv.Atoi(record[3])

		st[i] = models.TransferImportRow{
            fromStopId,
            toStopId,
			transferType,
			minTransferType,
		}
	}

	return st
}

func (m MySQLTransfersImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 4)

	table := fmt.Sprintf("`gtfs_%s`.`transfers`", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" from_stop_id," +
		" to_stop_id," +
		" transfer_type," +
		" min_transfer_time" +
		" ) VALUES %s"


	count := 0
	for _, entry := range as {
		t := entry.(models.TransferImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			t.FromStopId,
			t.ToStopId,
			t.TransferType,
			t.MinTransferTime,
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
