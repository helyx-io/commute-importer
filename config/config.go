package config

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"os"
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/helyx-io/gtfs-playground/database/mysql"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/auth"
	"github.com/goinggo/workpool"
	"strconv"
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
	Session *SessionConfig
	TmpDir string
	BaseURL string
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type SessionConfig struct{
	Secret string
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func Init() error {

	var err error

	dbDialect := "mysql"

	dbUsername := os.Getenv("GTFS_DB_USERNAME")
	if dbUsername == "" {
		dbUsername = "gtfs"
	}

	dbPassword := os.Getenv("GTFS_DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "gtfs"
	}

	dbDatabase := os.Getenv("GTFS_DB_DATABASE")
	if dbDatabase == "" {
		dbDatabase = "gtfs"
	}

	dbURL := fmt.Sprintf("%v:%v@/%v?charset=utf8mb4,utf8&parseTime=true", dbUsername, dbPassword, dbDatabase)

	dbMinCnx, _ := strconv.Atoi(os.Getenv("GTFS_DB_MIN_CNX"))
	if (dbMinCnx == 0) {
		dbMinCnx = 2
	}

	dbMaxCnx, _ := strconv.Atoi(os.Getenv("GTFS_DB_MAX_CNX"))
	if (dbMaxCnx == 0) {
		dbMaxCnx = 100
	}

	ConnectInfos = &database.DBConnectInfos{dbDialect, dbURL, dbMinCnx, dbMaxCnx}

	// Init Gorm
	if DB, err = mysql.InitDB(ConnectInfos); err != nil {
		return err
	}

	// Init GTFS Repository
	GTFS = mysql.CreateMySQLGTFSRepository(DB, ConnectInfos)

	// Init WorkPool
	WorkPool = workpool.New(32, 10000)

	OAuthInfos = &auth.OAuthInfos{os.Getenv("GOOGLE_AUTH_CLIENT_ID"), os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")}

	log.Println("[CONFIG] OAuth infos - ClientId :", "'" + OAuthInfos.ClientId + "'")
	log.Println("[CONFIG] OAuth infos - ClientSecret :", "'" + OAuthInfos.ClientSecret + "'")

	Session = &SessionConfig{os.Getenv("SESSION_SECRET")}

	TmpDir = os.Getenv("GTFS_TMP_DIR")

	BaseURL = os.Getenv("GTFS_BASE_URL")

	return nil
}

func Close() {
	if DB != nil {
		defer DB.Close()
	}
}
