package session

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"log"
	"net/http"
	"github.com/gorilla/sessions"
	"code.google.com/p/goauth2/oauth"
	"github.com/helyx-io/commute-importer/config"
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

func Init(sessionConfig *config.SessionConfig) {
	log.Println("[SESSION] Initializing cookie store with secret: ", "'" + sessionConfig.Secret + "'");
	store = sessions.NewCookieStore([]byte(sessionConfig.Secret))
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

	token := session.Values["token"].(oauth.Token)
	return &token, nil
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

	session.Values["token"] = *token
	log.Println("[SETTING_TOKEN] Session values: ", session.Values)

	session.Save(r, w)

	return nil
}


