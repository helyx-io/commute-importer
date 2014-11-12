package config

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"github.com/jinzhu/gorm"
	"github.com/helyx-io/gtfs-playground/database/mysql"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/auth"
	"github.com/goinggo/workpool"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	DB *gorm.DB
	GTFS database.GTFSRepository
	ConnectInfos *database.DBConnectInfos
	WorkPool *workpool.WorkPool
	OAuthInfos *auth.OAuthInfos
)

////////////////////////////////////////////////////////////////////////////////////////////////
/// Main Function
////////////////////////////////////////////////////////////////////////////////////////////////

func Init() error {

	var err error

	ConnectInfos = &database.DBConnectInfos{"mysql", "gtfs:gtfs@/gtfs?charset=utf8mb4,utf8&parseTime=true", 2, 100}

	// Init Gorm
	if DB, err = mysql.InitDB(ConnectInfos); err != nil {
		return err
	}

	// Init GTFS Repository
	GTFS = mysql.CreateMySQLGTFSRepository(DB, ConnectInfos)

	// Init WorkPool
	WorkPool = workpool.New(32, 10000)

	OAuthInfos = &auth.OAuthInfos{os.Getenv("GOOGLE_AUTH_CLIENT_ID"), os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")}

	return nil
}

func Close() {
	if DB != nil {
		defer DB.Close()
	}
}
