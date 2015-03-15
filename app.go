package main

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	appHandlers "github.com/helyx-io/gtfs-importer/handlers"
	"github.com/helyx-io/gtfs-importer/config"
	"github.com/helyx-io/gtfs-importer/controller"
	"github.com/helyx-io/gtfs-importer/utils"
	"github.com/helyx-io/gtfs-importer/session"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"log"
	"net/http"
	"os"
	"runtime"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Main Function
////////////////////////////////////////////////////////////////////////////////////////////////

func main() {

	// Init Runtime
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Init Profiling
	//	defer profile.Start(profile.MemProfile).Stop()
	//	defer profile.Start(profile.CPUProfile).Stop()

	// Init Logger
	logWriter, err := os.Create("/var/log/import-gtfs-helyx-io/access.log")
	utils.FailOnError(err, fmt.Sprintf("Could not access log"))
	defer logWriter.Close()

	// Init Config
	err = config.Init();
	utils.FailOnError(err, fmt.Sprintf("Could not init Configuration"))
	defer config.Close()

	session.Init()

	// Init Router
	router := initRouter()
	http.Handle("/", router)

	handlerChain := alice.New(
		appHandlers.LoggingHandler(logWriter),
//		appHandlers.ThrottleHandler,
//		appHandlers.TimeoutHandler,
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



////////////////////////////////////////////////////////////////////////////////////////////////
/// Router Configuration
////////////////////////////////////////////////////////////////////////////////////////////////

func initRouter() *mux.Router {
	r := mux.NewRouter()

	new(controller.IndexController).Init(r.PathPrefix("/").Subrouter())
//	new(controller.AuthController).Init(r.PathPrefix("/auth").Subrouter())
	new(controller.ImportController).Init(r.PathPrefix("/import").Subrouter())
	new(controller.AgencyController).Init(r.PathPrefix("/api/agencies").Subrouter())

	// Add handler for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
