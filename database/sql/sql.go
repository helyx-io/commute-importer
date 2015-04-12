package sql


////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
    "time"
	"strconv"
	"database/sql"
	"github.com/jinzhu/gorm"
    "github.com/helyx-io/gtfs-importer/data"
    "github.com/helyx-io/gtfs-importer/tasks"
    "github.com/helyx-io/gtfs-importer/utils"
    "github.com/helyx-io/gtfs-importer/config"
    "github.com/helyx-io/gtfs-importer/database"

    _ "github.com/lib/pq"
    _ "github.com/go-sql-driver/mysql"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// SQL
////////////////////////////////////////////////////////////////////////////////////////////////

func CreateSQLGTFSRepository(db *gorm.DB, dbInfos *config.DBConnectInfos) database.GTFSRepository {
	return SQLGTFSRepository{db, dbInfos}
}

type SQLGTFSRepository struct {
	db *gorm.DB
	dbInfos *config.DBConnectInfos
}

type SQLGTFSModelRepository struct {
	db *gorm.DB
	dbInfos *config.DBConnectInfos
}

type SQLConnectionProvider interface {
	OpenSqlConnection() (*sql.DB, error)
}

type SQLImportTask struct {
	tasks.ImportTask
	db *gorm.DB
	dbInfos *config.DBConnectInfos
}

func (r SQLGTFSRepository) CreateSchema(agencyKey string) error {

    filePath := fmt.Sprintf("resources/ddl/%s/create-schema.sql", r.dbInfos.Dialect)

    schema := fmt.Sprintf("gtfs_%s", agencyKey)

    log.Printf("Creating schema: '%s' with query from file path: '%s'", schema, filePath)

    ddlBytes, err := data.Asset(filePath)
    ddl := string(ddlBytes)
    utils.FailOnError(err, fmt.Sprintf("Could get ddl resource at path '%s' to create schema '%s'", filePath, schema))


    log.Println(fmt.Sprintf("Try to create schema: '%s' ...", schema))
	query := fmt.Sprintf(ddl, schema)
	log.Printf("Query: %s", query)

    err = r.db.Exec(query).Error

	if err == nil {
		log.Println(fmt.Sprintf("Created schema: '%s' with success", schema))
	}

	return err
}

func (it *SQLImportTask) OpenSqlConnection() (*sql.DB, error) {
	return sql.Open(it.dbInfos.Dialect, it.dbInfos.URL)
}


func recordValueAsString(record []string, offsets map[string]int, key string) string {

	offset, ok := offsets[key]

	if !ok {
		return ""
	}

	return record[offset]
}

func recordValueAsInt(record []string, offsets map[string]int, key string) int {

    offset, ok := offsets[key]

    if !ok {
        return 0
    }

    intValue, _ := strconv.Atoi(record[offset])

    return intValue
}

func recordValueAsTimeInt(record []string, offsets map[string]int, key string) int {

    offset, ok := offsets[key]

    if !ok {
        return 0
    }

    tc, _ := time.Parse("15:04:05", record[offset])
    intValue := tc.Hour() * 60 + tc.Minute()

    return intValue
}

func recordValueAsFloat(record []string, offsets map[string]int, key string) float64 {

	offset, ok := offsets[key]

	if !ok {
		return 0
	}

	floatValue, _ := strconv.ParseFloat(record[offset], 64)

	return floatValue
}
