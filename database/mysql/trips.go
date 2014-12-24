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

func (r MySQLGTFSRepository) Trips() database.GTFSCreatedModelRepository {
	return MySQLTripRepository{
		MySQLGTFSModelRepository{r.db,r.dbInfos},
	}
}

type MySQLTripRepository struct {
	MySQLGTFSModelRepository
}

func (s MySQLTripRepository) RemoveAllByAgencyKey(agencyKey string) (error) {

	table := fmt.Sprintf("`gtfs_%s`.`trips`", agencyKey)

	log.Println(fmt.Sprintf("Dropping table: '%s'", table))

	return s.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error
}

func (r MySQLTripRepository) CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, workPool *workpool.WorkPool, done chan error) workpool.PoolWorker {
	importTask := tasks.ImportTask{taskName, jobIndex, fileName, agencyKey, headers, lines, workPool, done}
	mysqlImportTask := MySQLImportTask{importTask, r.db, r.dbInfos}
	return MySQLTripsImportTask{mysqlImportTask}
}

func (s MySQLTripRepository) CreateTableByAgencyKey(agencyKey string) error {

	tmpTable := fmt.Sprintf("`gtfs_%s`.`trips`", agencyKey)

	log.Println(fmt.Sprintf("Creating table: '%s'", tmpTable))

	ddl, _ := data.Asset("resources/ddl/trips.sql")
	stmt := fmt.Sprintf(string(ddl), agencyKey);

	return s.db.Exec(stmt).Error
}

func (s MySQLTripRepository) AddIndexesByAgencyKey(agencyKey string) error {
	return nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQLStopRepository
////////////////////////////////////////////////////////////////////////////////////////////////

type MySQLTripsImportTask struct {
	MySQLImportTask
}

func (m MySQLTripsImportTask) DoWork(_ int) {
	m.ImportCsv(m, m);
}

func(m MySQLTripsImportTask) ConvertModels(headers []string, rs *models.Records) []interface{} {
	var st = make([]interface{}, len(*rs))

	for i, record := range *rs {
		serviceId, _ := strconv.Atoi(record[1])
		directionId, _ := strconv.Atoi(record[4])

		st[i] = models.TripImportRow{
			m.AgencyKey,
			record[0],
			serviceId,
			record[2],
			record[3],
			directionId,
			record[5],
			record[6],
		}
	}

	return st
}

func (m MySQLTripsImportTask) ImportModels(headers []string, as []interface{}) error {

	dbSql, err := m.OpenSqlConnection()

	if err != nil {
		panic(err.Error())
	}

	defer dbSql.Close()

	valueStrings := make([]string, 0, len(as))
	valueArgs := make([]interface{}, 0, len(as) * 7)

	table := fmt.Sprintf("`gtfs_%s`.`trips`", m.AgencyKey)

	log.Println(fmt.Sprintf("[%s][%d] Inserting into table: '%s'", m.AgencyKey, m.JobIndex, table))

	query := "INSERT INTO " + table + " (" +
		" route_id," +
		" service_id," +
		" trip_id," +
		" trip_headsign," +
		" direction_id," +
		" block_id," +
		" shape_id" +
		" ) VALUES %s"


	count := 0
	for _, entry := range as {
		t := entry.(models.TripImportRow)
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(
			valueArgs,
			t.RouteId,
			t.ServiceId,
			t.TripId,
			t.TripHeadsign,
			t.DirectionId,
			t.BlockId,
			t.ShapeId,
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
