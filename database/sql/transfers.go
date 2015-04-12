package sql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"github.com/helyx-io/gtfs-importer/database"
	"github.com/helyx-io/gtfs-importer/models"
	"github.com/helyx-io/gtfs-importer/tasks"
	"github.com/helyx-io/gtfs-importer/data"
	"github.com/helyx-io/gtfs-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

func (r SQLGTFSRepository) Transfers() database.GTFSCreatedModelRepository {
	return SQLTransferRepository{
		SQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type SQLTransferRepository struct {
	SQLGTFSModelRepository
}

func (s SQLTransferRepository) RemoveAllByAgencyKey(agencyKey string) (error) {
    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    return database.DropTable(s.db, s.dbInfos, schema, "transfers")
}


func (r SQLTransferRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, done}
	mysqlImportTask := SQLImportTask{importTask, r.db, r.dbInfos}
	return SQLTransfersImportTask{mysqlImportTask}
}

func (s SQLTransferRepository) CreateTableByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)
    table := fmt.Sprintf("%s.transfers", schema)

	log.Println(fmt.Sprintf("Creating table: '%s'", table))

    ddl, _ := data.Asset(fmt.Sprintf("resources/ddl/%s/transfers.sql", s.dbInfos.Dialect))
	stmt := fmt.Sprintf(string(ddl), schema);

    log.Printf("Query: %s", stmt)

	return s.db.Exec(stmt).Error
}

func (s SQLTransferRepository) AddIndexesByAgencyKey(agencyKey string) error {

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    err := database.CreateIndex(s.db, s.dbInfos, schema, "transfers", "from_stop_id")
    err = database.CreateIndex(s.db, s.dbInfos, schema, "transfers", "to_stop_id")
    err = database.CreateIndex(s.db, s.dbInfos, schema, "transfers", "transfer_type")

    return err
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type SQLTransfersImportTask struct {
	SQLImportTask
}

func (m SQLTransfersImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m SQLTransfersImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
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

func (m SQLTransfersImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 4)

	table := fmt.Sprintf("gtfs_%s.transfers", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" from_stop_id," +
		" to_stop_id," +
		" transfer_type," +
		" min_transfer_time" +
		" ) VALUES %s"


	count := 0
    for _, entry := range as {
        i := count * 7
		t := entry.(models.TransferImportRow)

        var args string
        if m.dbInfos.Dialect == "postgres" {
            args = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i + 1, i + 2, i + 3, i + 4)
        } else {
            args = "(?, ?, ?, ?)"
        }

        valueStrings = append(valueStrings, args)
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
