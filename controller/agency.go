package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	"github.com/akinsella/go-playground/database/mongo"
)

type AgencyController struct { }

func (agencyController *AgencyController) Init(r *mux.Router) {
	r.HandleFunc("/", agencyController.Agencies)
	r.HandleFunc("/{id:[0-9]+}", agencyController.AgenciesByKey)
}

func sendJson(w http.ResponseWriter, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if  err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (ac *AgencyController) Agencies(w http.ResponseWriter, r *http.Request) {
	results := mongo.FindAll("gtfs", "agency")
	sendJson(w, results)
}

func (ac *AgencyController) AgenciesByKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agencyKey := params["agencyKey"]

	log.Printf("Agency Key: %s", agencyKey)
	w.Write([]byte("Agency Key: " + agencyKey))
}
