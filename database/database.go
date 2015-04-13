package database

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/helyx-io/gtfs-importer/tasks"
    "github.com/helyx-io/gtfs-importer/config"
    "github.com/jinzhu/gorm"
    "fmt"
    "log"
    "github.com/helyx-io/gtfs-importer/data"
    "github.com/helyx-io/gtfs-importer/utils"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSRepository interface {
	CreateSchema(agencyKey string) error
	GtfsAgencies() GTFSModelRepository
	Agencies()	GTFSCreatedModelRepository
	CalendarDates() GTFSCreatedModelRepository
	Calendars() GTFSCreatedModelRepository
	Routes() GTFSCreatedModelRepository
	Stops() GTFSCreatedModelRepository
	StopTimes() GTFSCreatedModelRepository
	Transfers() GTFSCreatedModelRepository
	Trips() GTFSCreatedModelRepository
}

func InitDB(dbInfos *config.DBConnectInfos) (*gorm.DB, error) {
    db, err := gorm.Open(dbInfos.Dialect, dbInfos.URL)

    if err != nil {
        return nil, err
    }

    db.DB()

    // Then you could invoke `*sql.DB`'s functions with it
    db.DB().Ping()

    db.DB().SetMaxIdleConns(dbInfos.MaxIdelConns)
    db.DB().SetMaxOpenConns(dbInfos.MaxOpenConns)

    db.SingularTable(true)

    return &db, nil
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func Exec(db *gorm.DB, connectInfos *config.DBConnectInfos, filename string, params ...string) error {
    filePath := fmt.Sprintf("resources/ddl/%s/%s.sql", connectInfos.Dialect, filename)
    log.Printf("Executing query from file path: '%s'", filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for exec", filePath))
    execStmt := fmt.Sprintf(string(dml))
    log.Printf("Exec Stmt: '%s' - Params: %v", execStmt, params)

    return db.Exec(execStmt, params).Error
}

func DropTable(db *gorm.DB, connectInfos *config.DBConnectInfos, schema, tableName string) error {
    filePath := fmt.Sprintf("resources/ddl/%s/drop-table.sql", connectInfos.Dialect)
    log.Printf("Dropping table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for drop of table '%s.%s'", filePath, schema, tableName))
    dropStmt := fmt.Sprintf(string(dml), schema, tableName)
    log.Printf("Drop statement: %s", dropStmt)

    return db.Exec(dropStmt).Error
}

func CreateTable(db *gorm.DB, connectInfos *config.DBConnectInfos, schema, tableName string, dropIfExists bool) error {

    if dropIfExists {
        DropTable(db, connectInfos, schema, tableName)
    }

    filePath := fmt.Sprintf("resources/ddl/%s/create-table-%s.sql", connectInfos.Dialect, tableName)
    log.Printf("Creating table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for create of table '%s.%s'", filePath, schema, tableName))
    createStmt := fmt.Sprintf(string(dml), schema)
    log.Printf("Create statement: %s", createStmt)
    return db.Exec(createStmt).Error
}

func CreateIndex(db *gorm.DB, connectInfos *config.DBConnectInfos, schema, tableName, indexName string) error {
    return CreateIndexWithScript(db, connectInfos, "create-index", schema, tableName, indexName)

}

func CreateSpatialIndex(db *gorm.DB, connectInfos *config.DBConnectInfos, schema, tableName, indexName string) error {
    return CreateIndexWithScript(db, connectInfos, "create-spatial-index", schema, tableName, indexName)
}


func CreateIndexWithScript(db *gorm.DB, connectInfos *config.DBConnectInfos, script, schema, tableName, indexName string) error {
    filePath := fmt.Sprintf("resources/ddl/%s/%s.sql", connectInfos.Dialect, script)
    log.Printf("Creating index with for table with name: '%s.%s' with query from file path: '%s'", schema, tableName, filePath)

    dml, err := data.Asset(filePath)
    utils.FailOnError(err, fmt.Sprintf("Could get dml resource at path '%s' for create index of table '%s.%s'", filePath, schema, tableName))

    var createIndexStmt string
    if connectInfos.Dialect == "postgres" {
        createIndexStmt = fmt.Sprintf(string(dml), tableName, indexName, schema, tableName, indexName)
    } else {
        createIndexStmt = fmt.Sprintf(string(dml), schema, tableName, tableName, indexName, indexName)
    }
    log.Printf("Create statement for index: %s", createIndexStmt)
    return  db.Exec(createIndexStmt).Error
}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Interfaces
////////////////////////////////////////////////////////////////////////////////////////////////

type GTFSModelRepository interface {
	RemoveAllByAgencyKey(agencyKey string) error
	CreateImportTask(taskName string, jobIndex int, fileName, agencyKey string, headers []string, lines []byte, done chan error) tasks.Task
}

type GTFSCreatedModelRepository interface {
	GTFSModelRepository
	CreateTableByAgencyKey(agencyKey string) error
	AddIndexesByAgencyKey(agencyKey string) error
}

type GTFSAgencyRepository interface {
	GTFSCreatedModelRepository
}

type GTFSGtfsAgencyRepository interface {
	GTFSModelRepository
}

type GTFSCalendarRepository interface {
	GTFSCreatedModelRepository
}

type GTFSCalendarDateRepository interface {
	GTFSCreatedModelRepository
}

type GTFSRouteRepository interface {
	GTFSCreatedModelRepository
}

type GTFSTripRepository interface {
	GTFSCreatedModelRepository
}

type GTFSTransferRepository interface {
	GTFSCreatedModelRepository
}

type GTFSStopRepository interface {
	GTFSCreatedModelRepository
}

type GTFSStopTimeRepository interface {
	GTFSCreatedModelRepository
}
