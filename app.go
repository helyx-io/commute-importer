package main

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	appHandlers "github.com/helyx-io/gtfs-playground/handlers"
	"github.com/helyx-io/gtfs-playground/config"
	"github.com/helyx-io/gtfs-playground/controller"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/session"
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
	logWriter, err := os.Create("./access.log")
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
		appHandlers.ThrottleHandler,
		appHandlers.TimeoutHandler,
	).Then(router)

	// Init HTTP Server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: handlerChain,
	}

	log.Println("Listening ...")

	server.ListenAndServe()
}



////////////////////////////////////////////////////////////////////////////////////////////////
/// Router Configuration
////////////////////////////////////////////////////////////////////////////////////////////////

func initRouter() *mux.Router {
	r := mux.NewRouter()

	new(controller.IndexController).Init(r.PathPrefix("/").Subrouter())
	new(controller.AuthController).Init(r.PathPrefix("/auth").Subrouter())
	new(controller.ImportController).Init(r.PathPrefix("/import").Subrouter())

	new(controller.ApiController).Init(r.PathPrefix("/api/v1").Subrouter())

	// Add handler for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
