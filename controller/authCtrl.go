package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"github.com/gorilla/mux"
//	jwt "github.com/dgrijalva/jwt-go"
	"code.google.com/p/goauth2/oauth"
	"github.com/helyx-io/gtfs-playground/config"
	"net/http"
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
	r.HandleFunc("/", ac.Login)
	r.HandleFunc("/callback", ac.AuthCallback)

	client = &oauth.Config{
		ClientId:     config.OAuthInfos.ClientId,
		ClientSecret: config.OAuthInfos.ClientSecret,
		RedirectURL:  "http://localhost:3000/auth/callback",
		Scope:		  "https://www.googleapis.com/auth/plus.me https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile",
		AuthURL:      "https://accounts.google.com/o/oauth2/auth",
		TokenURL:     "https://accounts.google.com/o/oauth2/token",
	}

	ctransport = &oauth.Transport{Config: client}
}


func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body>"))
	//w.Write([]byte(fmt.Sprintf("<a href=\"/authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Login</a><br/>", url.QueryEscape("http://localhost:14000/appauth/code"))))
	w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Login</a><br/>", client.AuthCodeURL(""))))
	w.Write([]byte("</body></html>"))
}

func (ac *AuthController) AuthCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	code := r.Form.Get("code")

	w.Write([]byte("<html><body>"))
	w.Write([]byte("APP AUTH - CODE<br/>"))

	if code != "" {

		var jr *oauth.Token
		var err error

		// if parse, download and parse json
		if r.Form.Get("doparse") == "1" {
			jr, err = ctransport.Exchange(code)
			if err != nil {
				jr = nil
				w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", err)))
			}
		}

		// show json access token
		if jr != nil {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", jr.AccessToken)))
			if jr.RefreshToken != "" {
				w.Write([]byte(fmt.Sprintf("REFRESH TOKEN: %s<br/>\n", jr.RefreshToken)))
			}
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		cururl := *r.URL
		curq := cururl.Query()
		curq.Add("doparse", "1")
		cururl.RawQuery = curq.Encode()
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Download Token</a><br/>", cururl.String())))
	} else {
		w.Write([]byte("Nothing to do"))
	}

	w.Write([]byte("</body></html>"))
}
