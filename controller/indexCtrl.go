package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
//	appHandlers "github.com/helyx-io/commute-importer/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Index Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type IndexController struct {
	*mux.Router
}

func (c *IndexController) Init(r *mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/", c.indexHandler)

	handlerChain := alice.New(
//		appHandlers.LoggedInHandler,
	).Then(router)

	r.Handle("/", http.Handler(handlerChain))
}

func (c *IndexController) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, "Hello World.")
}
