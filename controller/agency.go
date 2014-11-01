package controller

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/helyx-io/gtfs-playground/database/mysql"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/database"
)

type AgencyController struct { }

var (
	gtfs database.GTFSRepository
	agencyRepository database.GTFSAgencyRepository
)

func (agencyController *AgencyController) Init(r *mux.Router) {
	db, err := mysql.InitDb(2, 100);
	gtfs = mysql.CreateMySQLGTFSRepository(db)
	agencyRepository = gtfs.Agencies().(database.GTFSAgencyRepository)

	utils.FailOnError(err, fmt.Sprintf("Could not open database"))

	r.HandleFunc("/", agencyController.Agencies)
	r.HandleFunc("/{agencyKey:[0-9]+}", agencyController.AgenciesByKey)
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
	agencies, err := agencyRepository.FindAll()

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else if agencies == nil {
		http.Error(w, "No agency found", 500)
	} else {
		sendJson(w, agencies)
	}
}

func (ac *AgencyController) AgenciesByKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agencyKey := params["agencyKey"]

	log.Printf("Agency Key: %s", agencyKey)

	agency, err := agencyRepository.FindByKey(agencyKey)

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else if agency != nil {
		http.Error(w, fmt.Sprintf("No agency found for key %v", agencyKey), 500)
	} else {
		sendJson(w, agency)
	}
}
