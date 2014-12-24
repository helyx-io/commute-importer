package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQL
////////////////////////////////////////////////////////////////////////////////////////////////

func InitDB(dbInfos *database.DBConnectInfos) (*gorm.DB, error) {
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


func CreateMySQLGTFSRepository(db *gorm.DB, dbInfos *database.DBConnectInfos) database.GTFSRepository {
	return MySQLGTFSRepository{db, dbInfos}
}

type MySQLGTFSRepository struct {
	db *gorm.DB
	dbInfos *database.DBConnectInfos
}

type MySQLGTFSModelRepository struct {
	db *gorm.DB
	dbInfos *database.DBConnectInfos
}

type SQLConnectionProvider interface {
	OpenSqlConnection() (*sql.DB, error)
}

type MySQLImportTask struct {
	tasks.ImportTask
	db *gorm.DB
	dbInfos *database.DBConnectInfos
}

func (r MySQLGTFSRepository) CreateSchema(agencyKey string) error {
	log.Println(fmt.Sprintf("Try to create schema for key: '%s' ...", agencyKey))
	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS `gtfs_%s` DEFAULT CHARACTER SET utf8mb4;", agencyKey)
	err := r.db.Exec(query).Error

	if err == nil {
		log.Println(fmt.Sprintf("Created schema for key: '%s' with success", agencyKey))
	}

	return err
}

func NewMySQLImportTask(importTask tasks.ImportTask, db *gorm.DB, dbInfos *database.DBConnectInfos) MySQLImportTask {
	return MySQLImportTask{importTask, db, dbInfos}
}

func (it *MySQLImportTask) OpenSqlConnection() (*sql.DB, error) {
	return sql.Open(it.dbInfos.Dialect, it.dbInfos.URL)
}
