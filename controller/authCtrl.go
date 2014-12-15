package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/gorilla/mux"
//	jwt "github.com/dgrijalva/jwt-go"
	"code.google.com/p/goauth2/oauth"
	"github.com/helyx-io/gtfs-playground/config"
	"github.com/helyx-io/gtfs-playground/session"
	"net/http"
	"log"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Server
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	client *oauth.Config
	ctransport *oauth.Transport
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Auth Controller
////////////////////////////////////////////////////////////////////////////////////////////////

type AuthController struct { }

func (ac *AuthController) Init(r *mux.Router) {

	// Init Router
	r.HandleFunc("/google", ac.AuthGoogle)
	r.HandleFunc("/google/callback", ac.AuthGoogleCallback)

	client = &oauth.Config{
		ClientId:     config.OAuthInfos.ClientId,
		ClientSecret: config.OAuthInfos.ClientSecret,
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scope:		  "https://www.googleapis.com/auth/plus.me https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
	}

	ctransport = &oauth.Transport{Config: client}
}


func (ac *AuthController) AuthGoogle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, client.AuthCodeURL(""), 302)
}

func (ac *AuthController) AuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	code := r.Form.Get("code")
	log.Println("[OAUTH_CONTROLLER] Code retrieved:", code)

	if code == "" {
		w.WriteHeader(500)
		w.Write([]byte("No code found"))
		return
	}

	var jr *oauth.Token
	var err error

	// if parse, download and parse json
	jr, err = ctransport.Exchange(code)
	if err != nil {
		jr = nil
		w.WriteHeader(500)
		w.Write([]byte("Error found: "))
		return
	}

	// show json access token
	if jr == nil {
		w.WriteHeader(500)
		w.Write([]byte("No token retrieved"))
		return
	}

	session.SetToken(w, r, jr)
	log.Println("[OAUTH_CONTROLLER] Token retrieved: { Access Token: ", jr.AccessToken, ", Refresh Token: ", jr.RefreshToken, ", Expiry: ", jr.Expiry, ", Extra: ", jr.Extra)

	http.Redirect(w, r, "/", 302)
}
