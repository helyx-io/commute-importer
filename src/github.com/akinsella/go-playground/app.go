package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/akinsella/go-playground/controller"
	"log"
	"net/http"
	"os"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	logWriter, err := os.Create("./access.log")
	check(err)
	defer logWriter.Close()

	router := initRouter()

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

	new(controller.IndexController).Init(r.PathPrefix("/").Subrouter())
	new(controller.AgencyController).Init(r.PathPrefix("/agencies").Subrouter())

	// Add handler for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
