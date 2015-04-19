package config

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "os"
    "fmt"
    "log"
    "strconv"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type Config struct {
    Http          *HttpConfig
    DataResources map[string]string
    ConnectInfos  *DBConnectInfos
    RedisInfos    *RedisConfig
    OAuthInfos    *OAuthInfos
    LoggerInfos   *LoggerConfig
    Session       *SessionConfig
    TmpDir        string
    BaseURL       string
}

type DBConnectInfos struct {
    Dialect string
    URL string
    MaxIdelConns int
    MaxOpenConns int
}

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

type LoggerConfig struct {
    Path string
}

type OAuthInfos struct {
    ClientId string
    ClientSecret string
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func Init() *Config {

    connectInfos := createConnectInfos()
    redisInfos := createRedisConfig()
    oauthInfos := createOAuthInfos()
    http := createHttpConfig()
    session := createSessionConfig()
    tmpDir := getTmpDir()
    baseURL := getBaseURL()
    loggerInfos := createLoggerConfig()

    dataResources := initDataSources()

    return &Config{http, dataResources, connectInfos, redisInfos, oauthInfos, loggerInfos, session, tmpDir, baseURL}
}

func getBaseURL() string {
    baseURL := os.Getenv("GTFS_BASE_URL")

    log.Printf("[CONFIG] Application - Base URL : '%s'", baseURL)

    return baseURL
}

func getTmpDir() string {
    tmpDir := os.Getenv("GTFS_TMP_DIR")

    log.Printf("[CONFIG] Application - Temp Directory : '%s'", tmpDir)

    return tmpDir
}

func createSessionConfig() *SessionConfig {

    sessionConfig := &SessionConfig{os.Getenv("SESSION_SECRET")}

    log.Printf("[CONFIG] Session - Secret :'%s'", sessionConfig.Secret)

    return sessionConfig
}

func createHttpConfig() *HttpConfig {

    httpPort, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
    if httpPort == 0 {
        httpPort = 3000
    }

    httpConfig := &HttpConfig{httpPort}

    log.Printf("[CONFIG] Application - HTTP Port : %d", httpConfig.Port)

    return httpConfig
}

func createLoggerConfig() *LoggerConfig {

    loggerFilePath := os.Getenv("LOGGER_FILE_PATH")
    if loggerFilePath == "" {
        loggerFilePath = "/var/log/gtfs-importer/access.log"
    }

    loggerConfig := &LoggerConfig{loggerFilePath}

    log.Printf("[CONFIG] Application - Logger File Path : %d", loggerConfig.Path)

    return loggerConfig
}

func createOAuthInfos() *OAuthInfos {

    oauthInfos := &OAuthInfos{os.Getenv("GOOGLE_AUTH_CLIENT_ID"), os.Getenv("GOOGLE_AUTH_CLIENT_SECRET")}

    log.Printf("[CONFIG] OAuth infos - ClientId : '%s'", oauthInfos.ClientId)
    log.Printf("[CONFIG] OAuth infos - ClientSecret : '%s'", oauthInfos.ClientSecret)

    return oauthInfos
}

func createRedisConfig() *RedisConfig {

    redisHost := os.Getenv("REDIS_HOST")

    if redisHost == "" {
        redisHost = "localhost"
    }

    redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
    if redisPort == 0 {
        redisPort = 8888
    }

    redisInfos := &RedisConfig{redisHost, redisPort}

    log.Printf("[CONFIG] Redis infos - Host : '%s'", redisInfos.Host)
    log.Printf("[CONFIG] Redis infos - Port : '%d'", redisInfos.Port)

    return redisInfos
}

func createConnectInfos() *DBConnectInfos {

    dbDialect := os.Getenv("GTFS_DB_DIALECT")
    if dbDialect == "" {
        dbDialect = "mysql"
    }

    dbHostname := os.Getenv("GTFS_DB_HOSTNAME")
    if dbHostname == "" {
        dbHostname = "localhost"
    }

    dbPort := os.Getenv("GTFS_DB_PORT")
    if dbPort == "" {
        dbPort = "3306"
    }

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

    dbURL := os.Getenv("GTFS_DB_URL")
    if dbURL == "" {
        log.Printf("[CONFIG] DB infos - Dialect : '%s'", dbDialect)
        log.Printf("[CONFIG] DB infos - Hostname : '%s'", dbHostname)
        log.Printf("[CONFIG] DB infos - Port : '%s'", dbPort)
        log.Printf("[CONFIG] DB infos - Database : '%s'", dbDatabase)
        log.Printf("[CONFIG] DB infos - Username : '%s'", dbUsername)
        log.Printf("[CONFIG] DB infos - Password : '%s'", "********")

        if dbDialect == "mysql" {

            dbURL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4,utf8&parseTime=true", dbUsername, dbPassword, dbHostname, dbPort, dbDatabase)
        } else if dbDialect == "postgres" {
            // postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full
            dbURL = fmt.Sprintf("%s://%v:%v@%v:%v/%v?sslmode=disable", dbDialect, dbUsername, dbDatabase, dbHostname, dbPort, dbDatabase)
        }
    }

    log.Printf("[CONFIG] DB infos - URL : '%s'", dbURL)

    dbMinCnx, _ := strconv.Atoi(os.Getenv("GTFS_DB_MIN_CNX"))
    if dbMinCnx == 0 {
        dbMinCnx = 128
    }

    dbMaxCnx, _ := strconv.Atoi(os.Getenv("GTFS_DB_MAX_CNX"))
    if dbMaxCnx == 0 {
        dbMaxCnx = 1000
    }

    log.Printf("[CONFIG] DB infos - Min Connections : %d", dbMinCnx)
    log.Printf("[CONFIG] DB infos - Max Connections : %d", dbMaxCnx)

    return &DBConnectInfos{dbDialect, dbURL, dbMinCnx, dbMaxCnx}
}

func initDataSources() map[string]string {
    
    dataResources := make(map[string]string)

    /**
     * Région: Ile-de-France
     * Agence: RATP
     * Site: http://data.ratp.fr/fr/les-donnees/fiche-de-jeu-de-donnees/dataset/offre-transport-de-la-ratp-format-gtfs.html?tx_icsoddatastore_pi1%5BreturnID%5D=38
     * Import: OK
     **/
//    dataResources["RATP"] = "http://dataratp.download.opendatasoft.com/RATP_GTFS_FULL.zip"
    dataResources["RATP"] = "http://localhost/data/RATP_GTFS_FULL_2015-2.zip"

    /**
     * Région: Ile-de-France
     * Agence: INTERCITES
     * Site: https://ressources.data.sncf.com/explore/dataset/sncf-intercites-gtfs/?tab=metas
     * Import: OK
     **/
    dataResources["INTERCITES"] = "http://ressources.data.sncf.com/api/datasets/1.0/sncf-intercites-gtfs/attachments/export_intercites_gtfs_last_zip/"

    /**
     * Région: Ile-de-France
     * Agence: TER
     * Site: https://ressources.data.sncf.com/explore/dataset/sncf-ter-gtfs/?tab=metas
     * Import: OK
     **/
    dataResources["TER"] = "http://ressources.data.sncf.com/api/datasets/1.0/sncf-ter-gtfs/attachments/export_ter_gtfs_last_zip/"

    /**
     * Région: Ile-de-France
     * Agence: Transilien
     * Site: https://ressources.data.sncf.com/explore/dataset/sncf-transilien-gtfs/?tab=metas
     * Import: OK
     **/
    dataResources["TRANSILIEN"] = "http://ressources.data.sncf.com/api/datasets/1.0/sncf-transilien-gtfs/attachments/export_tn_gtfs_last_zip/"

    /**
     * Région: REnnes
     * Agence: STAR (Keolis)
     * Site: http://data.keolis-rennes.com/fr/les-donnees/donnees-telechargeables.html
     * Import: OK
     **/
    dataResources["STAR"] = "http://data.keolis-rennes.com/fileadmin/OpenDataFiles/GTFS/GTFS-20150120.zip"

    /**
     * Ville: Nantes
     * Agence: SEMITAN
     * Site: http://data.nantes.fr/donnees/detail/arrets-horaires-et-circuits-tan/
     * Import: KO - Zip inside zip
     **/
    dataResources["SEMITAN"] = "http://data.nantes.fr/fileadmin/data/datastore/nm/mobilite/24440040400129_NM_TAN_00005/ARRETS_HORAIRES_CIRCUITS_TAN_gtfs.zip"

    /**
     * Département: Loires-Atlantique
     * Agence: LILA
     * Site: http://data.loire-atlantique.fr/donnees/detail/horaires-et-points-darrets-du-reseau-de-transport-lila-lignes-regulieres/
     * Import: KO - Empty zip
     **/
    dataResources["LILA"] = "http://data.loire-atlantique.fr/fileadmin/data/datastore/cg44/mobilite/22440002800011_CG44_MOB_09003/MOB09003_gtfs.zip"

    /**
     * Ville: Toulouse
     * Agence: Tisseo
     * Site: http://data.toulouse-metropole.fr/les-donnees/-/opendata/card/16271-reseau-tisseo-metro-bus-tram-gtfs
     * Import: KO - Duplicate key
     **/
    dataResources["TISSEO"] = "http://data.toulouse-metropole.fr/les-donnees/-/opendata/card/16271-reseau-tisseo-metro-bus-tram-gtfs/resource/document?p_p_state=exclusive&_5_WAR_opendataportlet_jspPage=%2Fsearch%2Fview_card.jsp"

    /**
     * Ville: Lille
     * Agence: Transpole
     * Site: http://www.transpole.fr/
     * Url: https://twitter.com/schignard/status/509780362785218560
     * Import: KO - line 3, column 29: bare " in non-quoted-field
     **/
    dataResources["LILLE"] = "https://www.data.gouv.fr/_uploads/resources/GTFS-Lille-20140801_1.zip"

    /**
     * Région: Gironde
     * Agence: Transgironde
     * Site: http://transgironde.gironde.fr/
     * Import: KO - STOP_TIMES - Duplicate entry '310-001-A|20120901|2-STOPPOINT:12630' for key 'PRIMARY')
     **/
    dataResources["TRANSGIRONDE"] = "http://catalogue.datalocale.fr/storage/f/2013-03-19T174734/ExportGTFS20130319.zip"

    /**
     * Région: Gironde
     * Agence: Cub
     * Site: http://www.infotbc.com/
     * Import: OK
     **/
    dataResources["CUB"] = "http://data.lacub.fr/files.php?gid=67&format=14"

    /**
     * Ville: Metz
     * Agence: LeMet / Mettis
     * Site: https://github.com/ridem/fr.lemet/tree/master/TransportsMetzHelper/src/fr/ybo/transportsrenneshelper/gtfs
     * Import: KO - File is missing
     **/
    dataResources["LEMET"] = "https://geo-ws.metzmetropole.fr/services/opendata/gtfs_hiver_nov14_031114-050715.zip"

    /**
     * Ville: Angers
     * Agence: Irigo
     * Site: http://data.angers.fr/donnees/mobilite-horaire-de-passage-theorique/
     * Import: KO - Data available in subfolder (20141222-20150208)
     **/
    dataResources["IRIGO"] = "http://data.angers.fr/?eID=ics_od_datastoredownload&file=824"

    /**
     * Ville: Nancy
     * Agence: STAN
     * Site: http://opendata.grand-nancy.org/
     * Import: KO - Agency file custom
     **/
    dataResources["STAN"] = "http://opendata.grand-nancy.org/?eID=ics_od_datastoredownload&file=333"


    return dataResources
}

