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
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	logWriter, err := os.Create("./access.log")
	check(err)
	defer logWriter.Close()

	router := initRouter()

	http.Handle("/", router)

//	http.Handle("/debug/pprof/", http.HandlerFunc(Index))
//	http.Handle("/debug/pprof/cmdline", http.HandlerFunc(Cmdline))
//	http.Handle("/debug/pprof/profile", http.HandlerFunc(Profile))
//	http.Handle("/debug/pprof/symbol", http.HandlerFunc(Symbol))

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
