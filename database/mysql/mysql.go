package mysql

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/jinzhu/gorm"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/tasks"
	_ "github.com/go-sql-driver/mysql"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// MySQL
////////////////////////////////////////////////////////////////////////////////////////////////

func InitDB(dbInfos *database.ConnectInfos) (*gorm.DB, error) {
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


func CreateMySQLGTFSRepository(db *gorm.DB) database.GTFSRepository {
	return MySQLGTFSRepository{db}
}

type MySQLGTFSRepository struct {
	db *gorm.DB
}

type MySQLGTFSModelRepository struct {
	db *gorm.DB
}

type MySQLImportTask struct {
	tasks.ImportTask
	db *gorm.DB
}
