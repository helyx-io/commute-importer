package main

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
    "os"
	"fmt"
    "log"
    "runtime"
    "net/http"

    "github.com/jinzhu/gorm"
    "github.com/gorilla/mux"
    "github.com/justinas/alice"

    "github.com/helyx-io/gtfs-importer/utils"
    "github.com/helyx-io/gtfs-importer/config"
	"github.com/helyx-io/gtfs-importer/session"
    "github.com/helyx-io/gtfs-importer/database"
    "github.com/helyx-io/gtfs-importer/database/sql"
    "github.com/helyx-io/gtfs-importer/handlers"
    "github.com/helyx-io/gtfs-importer/controller"
    "gopkg.in/redis.v2"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
    DB              *gorm.DB
    GTFS            database.GTFSRepository
    RedisClient     *redis.Client
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Main Function
////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
    defer Close()

	// Init Runtime
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Init Profiling
	//	defer profile.Start(profile.MemProfile).Stop()
	//	defer profile.Start(profile.CPUProfile).Stop()

	// Init Config
	config := config.Init();

    // Init Logger
    logWriter, err := os.Create(config.LoggerInfos.Path)
    utils.FailOnError(err, fmt.Sprintf("Could not access log"))
    defer logWriter.Close()

    DB, err = database.InitDB(config.ConnectInfos)
    utils.FailOnError(err, fmt.Sprintf("Could not init Database"))

    // Init GTFS Repository
    GTFS = sql.CreateSQLGTFSRepository(DB, config.ConnectInfos)


    RedisClient = redis.NewTCPClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d", config.RedisInfos.Host, config.RedisInfos.Port),
        Password: "", // no password set
        DB:       0,  // use default DB
        PoolSize: 16,
    })


    session.Init(config.Session)

	// Init Router
	router := initRouter(config)
	http.Handle("/", router)

	handlerChain := alice.New(
        handlers.LoggingHandler(logWriter),
//		handlers.ThrottleHandler,
//		handlers.TimeoutHandler,
	).Then(router)


	// Init HTTP Server
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Http.Port),
		Handler: handlerChain,
	}

	log.Println(fmt.Sprintf("Listening on port '%d' ...", config.Http.Port))


	err = server.ListenAndServe()
    utils.FailOnError(err, fmt.Sprintf("Could not listen and server"))
}

func Close() {
    if DB != nil {
        defer DB.Close()
    }
}

////////////////////////////////////////////////////////////////////////////////////////////////
/// Router Configuration
////////////////////////////////////////////////////////////////////////////////////////////////

func initRouter(config *config.Config) *mux.Router {
	r := mux.NewRouter()

	new(controller.IndexController).Init(r.PathPrefix("/").Subrouter())
	new(controller.ImportController).Init(r.PathPrefix("/import").Subrouter(), config.DataResources, config.TmpDir, RedisClient, DB, config.ConnectInfos, GTFS)

	// Add handler for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
