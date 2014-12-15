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

func (r MySQLGTFSRepository) Transfers() database.GTFSModelRepository {
	return MySQLTransferRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLTransferRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLTransferRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
	return s.db.Table("transfers").Where("agency_key = ?", agencyKey).Delete(models.Transfer{}).Error
}


func (r MySQLTransferRepository) CreateImportTask(name, agencyKey string, lines []byte, workPool *workpool.WorkPool) workpool.PoolWorker {
	importTask := tasks.ImportTask{name, agencyKey, lines, workPool}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLTransfersImportTask{mysqlImportTask}
}

func (s MySQLTransferRepository) FindAll() (*models.Transfers, error) {
	var transfers models.Transfers
	err := s.db.Table("transfers").Limit(1000).Find(&transfers).Error

	return &transfers, err
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

func(m MySQLTransfersImportTask) ConvertModels(rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		transferType, _ := strconv.Atoi(record[2])
		minTransferType, _ := strconv.Atoi(record[3])

		st[i] = models.TransferImportRow{
			m.AgencyKey,
			record[0],
			record[1],
			transferType,
			minTransferType,
		}
	}

	return st
}

func (m MySQLTransfersImportTask) ImportModels(as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 9)

	for _, entry := range as {
		t := entry.(models.TransferImportRow)
		valueStrings = append(valueStrings, "('" + m.AgencyKey + "', ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			t.FromStopId,
			t.ToStopId,
			t.TransferType,
			t.MinTransferTime,
		)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO transfers (" +
			" agency_key," +
			" from_stop_id," +
			" to_stop_id," +
			" transfer_type," +
			" min_transfer_time" +
		" ) VALUES %s", strings.Join(valueStrings, ","))


	_, err = dbSql.Exec(stmt, valueArgs...)

	return err

}
