package session

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"net/http"
	"github.com/gorilla/sessions"
	"code.google.com/p/goauth2/oauth"
	"github.com/helyx-io/gtfs-playground/config"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	store *sessions.CookieStore
)




////////////////////////////////////////////////////////////////////////////////////////////////
/// Functions
////////////////////////////////////////////////////////////////////////////////////////////////

func Init() {
	store = sessions.NewCookieStore([]byte(config.Session.Secret))
}

func Get(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "session")
}

func GetToken(r *http.Request) (*oauth.Token, error) {
	session, err := Get(r)

	if err != nil {
		return nil, err
	}

	log.Println("[GET_TOKEN] Session values: ", session.Values)

	return session.Values["token"].(*oauth.Token), nil
}

func HasToken(r *http.Request) bool {
	session, err := Get(r)

	if err != nil {
		return false
	}

	log.Println("[CHECK_TOKEN] Session values: ", session.Values)

	return session.Values["token"] != nil
}

func SetToken(w http.ResponseWriter, r *http.Request, token *oauth.Token) error {
	session, err := Get(r)

	if err != nil {
		return err
	}

	log.Println("Writing token", token)

	session.Values["token"] = token

	session.Save(r, w)

	return nil
}


