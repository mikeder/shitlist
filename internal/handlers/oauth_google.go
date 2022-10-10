package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mikeder/shitlist/internal/database"
	"golang.org/x/oauth2"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (a *API) OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := setOauthStateCookie(w)

	u := a.googleOauth.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (a *API) OauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, err := r.Cookie(stateCookieName)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"), a.googleOauth)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// get or create user in database
	u, err := a.userStore.GetUserByEmail(data.Email)
	if err != nil {
		log.Println("get user: " + err.Error())
	}
	if u == nil {
		u, err = a.userStore.AddUser(data.Name, data.Email)
		if err != nil {
			log.Println("add user: " + err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
		ua, err := a.userStore.AddUserAuthentication(u.ID, database.AuthenticationTypeGoogle)
		if err != nil {
			log.Println("add user authentication: " + err.Error())
		}
		log.Println("added user authentication: " + ua.ID)
	}
	setUserIDCookie(w, u.ID)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func getUserDataFromGoogle(code string, cfg *oauth2.Config) (googleUserData, error) {
	// Use code to get token and get user info from Google.
	var userData googleUserData

	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return userData, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return userData, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&userData); err != nil {
		return userData, err
	}

	return userData, nil
}

type googleUserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
