package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	appHandlers "github.com/helyx-io/gtfs-playground/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

////////////////////////////////////////////////////////////////////////////////////////////////
/// Auth Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type ApiController struct { }

func (ac *ApiController) Init(r *mux.Router) {

	// Init Router
	router := mux.NewRouter()

	new(ImportController).Init(r.PathPrefix("/import").Subrouter())
	new(AgencyController).Init(r.PathPrefix("/agencies").Subrouter())
	new(CalendarController).Init(r.PathPrefix("/calendars").Subrouter())
	new(CalendarDateController).Init(r.PathPrefix("/calendar-dates").Subrouter())
	new(RouteController).Init(r.PathPrefix("/routes").Subrouter())
	new(TripController).Init(r.PathPrefix("/trips").Subrouter())
	new(TransferController).Init(r.PathPrefix("/transfers").Subrouter())
	new(StopController).Init(r.PathPrefix("/stops").Subrouter())
	new(StopTimeController).Init(r.PathPrefix("/stop-times").Subrouter())


	handlerChain := alice.New(
		appHandlers.LoggedInHandler,
	).Then(router)

	r.Handle("/", handlerChain)
}
