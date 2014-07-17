package main

import (
	"agency"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/agencies", agency.Agencies).Methods("GET")
	router.HandleFunc("/agencies/{agencyKey}", agency.AgenciesByKey).Methods("GET")

	http.Handle("/", router)

	log.Println("Listening ...")
	http.ListenAndServe(":3000", nil)
}
