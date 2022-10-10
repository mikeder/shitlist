package handlers

import (
	"github.com/mikeder/shitlist/internal/config"
	"github.com/mikeder/shitlist/internal/database"
	"github.com/mikeder/shitlist/pkg/go/shitlist/v1/shitlistv1connect"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var _ shitlistv1connect.ShitlistServiceHandler = (*API)(nil)

type API struct {
	clickStore  database.ClickStore
	userStore   database.UserStore
	githubOauth *oauth2.Config
	googleOauth *oauth2.Config
}

func NewAPI(cfg *config.Specification) (*API, error) {
	db, err := database.NewPersistentClickStore(cfg.Database)
	if err != nil {
		return nil, err
	}
	return &API{
		clickStore: db,
		userStore:  db,
		githubOauth: &oauth2.Config{
			RedirectURL:  cfg.GithubOauthRedirectURL,
			ClientID:     cfg.GithubOauthClientID,
			ClientSecret: cfg.GithubOauthClientSecret,
			Scopes:       []string{}, // scopes derived from GitHub App and user permissions.
			Endpoint:     github.Endpoint,
		},
		googleOauth: &oauth2.Config{
			RedirectURL:  cfg.GoogleOauthRedirectURL,
			ClientID:     cfg.GoogleOauthClientID,
			ClientSecret: cfg.GoogleOauthClientSecret,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}, nil
}
