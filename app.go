package main

import (
	"fmt"
	"github.com/akinsella/go-playground/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"runtime"
	"net/http/pprof"
	"github.com/davecheney/profile"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//	defer profile.Start(profile.MemProfile).Stop()
		defer profile.Start(profile.CPUProfile).Stop()

	runtime.GOMAXPROCS(runtime.NumCPU())

	logWriter, err := os.Create("./access.log")
	check(err)
	defer logWriter.Close()

	router := initRouter()


	router.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))


	http.Handle("/", router)

	loggingHandler := handlers.LoggingHandler(logWriter, router)

	log.Println("Listening ...")

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: loggingHandler,
	}

	server.ListenAndServe()
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))

	new(controller.IndexController).Init(r.PathPrefix("/").Subrouter())
	new(controller.ImportController).Init(r.PathPrefix("/import").Subrouter())
	new(controller.AgencyController).Init(r.PathPrefix("/agencies").Subrouter())

	// Add handler for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
