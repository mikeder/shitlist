package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var githubOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:10000/auth/github/callback",
	ClientID:     os.Getenv("GITHUB_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{}, // scopes derived from GitHub App and user permissions.
	Endpoint:     github.Endpoint,
}

func OauthGithubLogin(w http.ResponseWriter, r *http.Request) {
	// Create oauthState cookie
	oauthState := setOauthStateCookie(w)

	u := githubOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func OauthGithubCallback(w http.ResponseWriter, r *http.Request) {
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

	data, err := getUserDataFromGithub(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// TODO: do something with user data

	t, err := template.New("foo").Parse(`Hello {{.Login}}, your ID is: {{.ID}}.`)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	err = t.ExecuteTemplate(w, t.Name(), data)
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
}

func getUserDataFromGithub(code string) (githubUserData, error) {
	var userData githubUserData

	token, err := githubOauthConfig.Exchange(context.Background(), code)
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
