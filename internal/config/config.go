package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mikeder/shitlist/internal/database"
)

const envPrefix = ""

type Specification struct {
	// Database Config
	Database database.PersistentDataStoreConfig `required:"true" split_words:"true"`

	// OAuth Configs
	GithubOauthClientID     string `required:"true" split_words:"true"`
	GithubOauthClientSecret string `required:"true" split_words:"true"`
	GithubOauthRedirectURL  string `default:"http://localhost:10000/auth/github/callback" split_words:"true"`
	GoogleOauthClientID     string `required:"true" split_words:"true"`
	GoogleOauthClientSecret string `required:"true" split_words:"true"`
	GoogleOauthRedirectURL  string `default:"http://localhost:10000/auth/google/callback" split_words:"true"`

	// Server Config
	TemplatesDirectory  string `default:"../../templates" split_words:"true"`
	ServerListenAddress string `default:":10000" split_words:"true"`
}

func (s *Specification) LoadFromEnvironment() error {
	return envconfig.Process(envPrefix, s)
}

func (s *Specification) Usage() error {
	return envconfig.Usage(envPrefix, s)
}
