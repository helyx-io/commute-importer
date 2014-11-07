package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/config"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type CalendarController struct { }


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	calendarRepository database.GTFSCalendarRepository
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Calendar Controller
////////////////////////////////////////////////////////////////////////////////////////////////

func (calendarController *CalendarController) Init(r *mux.Router) {
	calendarRepository = config.GTFS.Calendars().(database.GTFSCalendarRepository)

	r.HandleFunc("/", calendarController.Calendars)
//	r.HandleFunc("/{id:[0-9]+}", calendarController.CalendarById)
}

func (ac *CalendarController) Calendars(w http.ResponseWriter, r *http.Request) {
	calendars, err := calendarRepository.FindAll()

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else if calendars == nil {
		http.Error(w, "No calendar found", 500)
	} else {
		utils.SendJSON(w, calendars)
	}
}
//
//func (ac *CalendarController) CalendarById(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	idParam := params["id"]
//
//	log.Printf("id: %s", idParam)
//
//	id, _ := strconv.Atoi(idParam)
//
//	calendar, err := calendarRepository.FindById(id)
//
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//	} else if calendar == nil {
//		http.Error(w, fmt.Sprintf("No calendar found for key %v", id), 500)
//	} else {
//		utils.SendJSON(w, calendar)
//	}
//}
