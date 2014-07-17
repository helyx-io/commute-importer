package agency

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Agencies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Agency List"))
}

func AgenciesByKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agencyKey := params["agencyKey"]
	log.Printf("Agency Key: %s", agencyKey)
	w.Write([]byte("Agency Key: " + agencyKey))
}
