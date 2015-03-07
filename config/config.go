package config

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"github.com/helyx-io/gtfs-playground/auth"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/database/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	DB            *gorm.DB
	Http          *HttpConfig
	GTFS          database.GTFSRepository
	DataResources map[string]string
	ConnectInfos  *database.DBConnectInfos
    RedisInfos    *RedisConfig
	OAuthInfos    *auth.OAuthInfos
	Session       *SessionConfig
	TmpDir        string
	BaseURL       string
)

////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type SessionConfig struct {
	Secret string
}

type HttpConfig struct {
    Port int
}

type RedisConfig struct {
    Host string
    Port int
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

	log.Printf("[CONFIG] DB infos - Database : '%s'", dbDatabase)
	log.Printf("[CONFIG] DB infos - Username : '%s'", dbUsername)
	log.Printf("[CONFIG] DB infos - Password : '%s'", "********")

	dbURL := fmt.Sprintf("%v:%v@/%v?charset=utf8mb4,utf8&parseTime=true", dbUsername, dbPassword, dbDatabase)

	dbMinCnx, _ := strconv.Atoi(os.Getenv("GTFS_DB_MIN_CNX"))
	if dbMinCnx == 0 {
		dbMinCnx = 2
	}

	dbMaxCnx, _ := strconv.Atoi(os.Getenv("GTFS_DB_MAX_CNX"))
	if dbMaxCnx == 0 {
		dbMaxCnx = 100
	}

	log.Printf("[CONFIG] DB infos - Min Connections : %d", dbMinCnx)
	log.Printf("[CONFIG] DB infos - Max Connections : %d", dbMaxCnx)

	ConnectInfos = &database.DBConnectInfos{dbDialect, dbURL, dbMinCnx, dbMaxCnx}

	// Init Gorm
	if DB, err = mysql.InitDB(ConnectInfos); err != nil {
		return err
	}

	// Init GTFS Repository
	GTFS = mysql.CreateMySQLGTFSRepository(DB, ConnectInfos)

    redisHost := os.Getenv("REDIS_HOST")

    if redisHost == "" {
        redisHost = "localhost"
    }

    redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
    if redisPort == 0 {
        redisPort = 6380
    }

    RedisInfos = &RedisConfig{redisHost, redisPort}

    log.Printf("[CONFIG] Redis infos - Host : '%s'", RedisInfos.Host)
    log.Printf("[CONFIG] Redis infos - Port : '%d'", RedisInfos.Port)


    OAuthInfos = &auth.OAuthInfos{os.Getenv("GOOGLE_AUTH_CLIENT_ID"), os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")}

	log.Printf("[CONFIG] OAuth infos - ClientId : '%s'", OAuthInfos.ClientId)
	log.Printf("[CONFIG] OAuth infos - ClientSecret : '%s'", OAuthInfos.ClientSecret)

	Session = &SessionConfig{os.Getenv("SESSION_SECRET")}

	log.Printf("[CONFIG] Session - Secret :'%s'", Session.Secret)

	TmpDir = os.Getenv("GTFS_TMP_DIR")

	log.Printf("[CONFIG] Application - Temp Directory : '%s'", TmpDir)

	BaseURL = os.Getenv("GTFS_BASE_URL")

	log.Printf("[CONFIG] Application - Base URL : '%s'", BaseURL)

	DataResources = make(map[string]string)

	/**
	 * Région: Ile-de-France
	 * Agence: RATP
	 * Site: http://data.ratp.fr/fr/les-donnees/fiche-de-jeu-de-donnees/dataset/offre-transport-de-la-ratp-format-gtfs.html?tx_icsoddatastore_pi1%5BreturnID%5D=38
	 * Import: OK
	 **/
	DataResources["RATP"] = "http://dataratp.download.opendatasoft.com/RATP_GTFS_FULL.zip"

	/**
	 * Région: Ile-de-France
	 * Agence: INTERCITES
	 * Site: https://ressources.data.sncf.com/explore/dataset/sncf-intercites-gtfs/?tab=metas
	 * Import: OK
	 **/
	DataResources["INTERCITES"] = "http://ressources.data.sncf.com/api/datasets/1.0/sncf-intercites-gtfs/attachments/export_intercites_gtfs_last_zip/"

	/**
	 * Région: Ile-de-France
	 * Agence: TER
	 * Site: https://ressources.data.sncf.com/explore/dataset/sncf-ter-gtfs/?tab=metas
	 * Import: OK
	 **/
	DataResources["TER"] = "http://ressources.data.sncf.com/api/datasets/1.0/sncf-ter-gtfs/attachments/export_ter_gtfs_last_zip/"

	/**
	 * Région: Ile-de-France
	 * Agence: Transilien
	 * Site: https://ressources.data.sncf.com/explore/dataset/sncf-transilien-gtfs/?tab=metas
	 * Import: OK
	 **/
	DataResources["TRANSILIEN"] = "http://ressources.data.sncf.com/api/datasets/1.0/sncf-transilien-gtfs/attachments/export_tn_gtfs_last_zip/"

	/**
	 * Région: REnnes
	 * Agence: STAR (Keolis)
	 * Site: http://data.keolis-rennes.com/fr/les-donnees/donnees-telechargeables.html
	 * Import: OK
	 **/
	DataResources["STAR"] = "http://data.keolis-rennes.com/fileadmin/OpenDataFiles/GTFS/GTFS-20150120.zip"

	/**
	 * Ville: Nantes
	 * Agence: SEMITAN
	 * Site: http://data.nantes.fr/donnees/detail/arrets-horaires-et-circuits-tan/
	 * Import: KO - Zip inside zip
	 **/
	DataResources["SEMITAN"] = "http://data.nantes.fr/fileadmin/data/datastore/nm/mobilite/24440040400129_NM_TAN_00005/ARRETS_HORAIRES_CIRCUITS_TAN_gtfs.zip"

	/**
	 * Département: Loires-Atlantique
	 * Agence: LILA
	 * Site: http://data.loire-atlantique.fr/donnees/detail/horaires-et-points-darrets-du-reseau-de-transport-lila-lignes-regulieres/
	 * Import: KO - Empty zip
	 **/
	DataResources["LILA"] = "http://data.loire-atlantique.fr/fileadmin/data/datastore/cg44/mobilite/22440002800011_CG44_MOB_09003/MOB09003_gtfs.zip"

	/**
	 * Ville: Toulouse
	 * Agence: Tisseo
	 * Site: http://data.toulouse-metropole.fr/les-donnees/-/opendata/card/16271-reseau-tisseo-metro-bus-tram-gtfs
	 * Import: KO - Duplicate key
	 **/
	DataResources["TISSEO"] = "http://data.toulouse-metropole.fr/les-donnees/-/opendata/card/16271-reseau-tisseo-metro-bus-tram-gtfs/resource/document?p_p_state=exclusive&_5_WAR_opendataportlet_jspPage=%2Fsearch%2Fview_card.jsp"

	/**
	 * Ville: Lille
	 * Agence: Transpole
	 * Site: http://www.transpole.fr/
	 * Url: https://twitter.com/schignard/status/509780362785218560
	 * Import: KO - line 3, column 29: bare " in non-quoted-field
	 **/
	DataResources["LILLE"] = "https://www.data.gouv.fr/_uploads/resources/GTFS-Lille-20140801_1.zip"

	/**
	 * Région: Gironde
	 * Agence: Transgironde
	 * Site: http://transgironde.gironde.fr/
	 * Import: KO - STOP_TIMES - Duplicate entry '310-001-A|20120901|2-STOPPOINT:12630' for key 'PRIMARY')
	 **/
	DataResources["TRANSGIRONDE"] = "http://catalogue.datalocale.fr/storage/f/2013-03-19T174734/ExportGTFS20130319.zip"

	/**
	 * Région: Gironde
	 * Agence: Cub
	 * Site: http://www.infotbc.com/
	 * Import: OK
	 **/
	DataResources["CUB"] = "http://data.lacub.fr/files.php?gid=67&format=14"

	/**
	 * Ville: Metz
	 * Agence: LeMet / Mettis
	 * Site: https://github.com/ridem/fr.lemet/tree/master/TransportsMetzHelper/src/fr/ybo/transportsrenneshelper/gtfs
	 * Import: KO - File is missing
	 **/
	DataResources["LEMET"] = "https://geo-ws.metzmetropole.fr/services/opendata/gtfs_hiver_nov14_031114-050715.zip"

	/**
	 * Ville: Angers
	 * Agence: Irigo
	 * Site: http://data.angers.fr/donnees/mobilite-horaire-de-passage-theorique/
	 * Import: KO - Data available in subfolder (20141222-20150208)
	 **/
	DataResources["IRIGO"] = "http://data.angers.fr/?eID=ics_od_datastoredownload&file=824"

	/**
	 * Ville: Nancy
	 * Agence: STAN
	 * Site: http://opendata.grand-nancy.org/
	 * Import: KO - Agency file custom
	 **/
	DataResources["STAN"] = "http://opendata.grand-nancy.org/?eID=ics_od_datastoredownload&file=333"

	httpPort, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if httpPort == 0 {
		httpPort = 3000
	}

	Http = &HttpConfig{httpPort}

	log.Printf("[CONFIG] Application - HTTP Port : %d", Http.Port)

	return nil
}

func Close() {
	if DB != nil {
		defer DB.Close()
	}
}
