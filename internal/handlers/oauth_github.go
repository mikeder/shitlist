package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

func (a *API) OauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := setOauthStateCookie(w)

	u := a.githubOauth.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (a *API) OauthGithubCallback(w http.ResponseWriter, r *http.Request) {
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

	data, err := getUserDataFromGithub(r.FormValue("code"), a.githubOauth)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// get or create user in database
	u, err := a.us.GetUserByEmail(data.Email)
	if err != nil {
		log.Println("get user: " + err.Error())
	}
	if u == nil {
		u, err = a.us.AddUser(data.Login, data.Email)
		if err != nil {
			log.Println("add user: " + err.Error())
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
	setUserIDCookie(w, u.ID)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func getUserDataFromGithub(code string, cfg *oauth2.Config) (githubUserData, error) {
	var userData githubUserData

	token, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		return userData, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", http.NoBody)
	if err != nil {
		return userData, fmt.Errorf("failed creating user request: %s", err.Error())
	}
	req.Header.Set("Authorization", "token "+token.AccessToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return userData, fmt.Errorf("failed performing user request: %s", err.Error())
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&userData); err != nil {
		return userData, err
	}

	return userData, nil
}

type githubUserData struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
