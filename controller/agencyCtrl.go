package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/config"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type AgencyController struct { }


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	agencyRepository database.GTFSAgencyRepository
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Helper Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func SendJSON(w http.ResponseWriter, data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if  err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}


////////////////////////////////////////////////////////////////////////////////////////////////
/// Agency Controller
////////////////////////////////////////////////////////////////////////////////////////////////

func (agencyController *AgencyController) Init(r *mux.Router) {
	agencyRepository = config.GTFS.Agencies().(database.GTFSAgencyRepository)

	r.HandleFunc("/", agencyController.Agencies)
	r.HandleFunc("/{agencyKey:[0-9]+}", agencyController.AgenciesByKey)
}

func (ac *AgencyController) Agencies(w http.ResponseWriter, r *http.Request) {
	agencies, err := agencyRepository.FindAll()

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else if agencies == nil {
		http.Error(w, "No agency found", 500)
	} else {
		utils.SendJSON(w, agencies)
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
		utils.SendJSON(w, agency)
	}
}
