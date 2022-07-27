package server

import (
	"net/http"
	"time"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"github.com/mikeder/shitlist/internal/config"
	"github.com/mikeder/shitlist/internal/handlers"
	"github.com/mikeder/shitlist/pkg/go/shitlist/v1/shitlistv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Setup(cfg *config.Specification) (*http.Server, error) {
	// construct API with handlers
	api, err := handlers.NewAPI(cfg)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()

	// register file handlers on mux
	mux.Handle("/", http.FileServer(http.Dir("../../templates/")))

	// register GitHub OAuth handlers on mux
	mux.HandleFunc("/auth/github/login", api.OauthGithubLogin)
	mux.HandleFunc("/auth/github/callback", api.OauthGithubCallback)

	// register Google OAuth handlers on mux
	mux.HandleFunc("/auth/google/login", api.OauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", api.OauthGoogleCallback)

	// register reflection handlers on mux
	reflector := grpcreflect.NewStaticReflector(
		"shitlist.v1.ShitlistService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	// Many tools still expect the older version of the server reflection API, so
	// most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	// If you don't need to support HTTP/2 without TLS (h2c), you can drop
	// x/net/http2 and use http.ListenAndServeTLS instead.

	// register ShitListService handlers
	path, handler := shitlistv1connect.NewShitlistServiceHandler(api)
	mux.Handle(path, handler)

	return &http.Server{
		Addr: cfg.ServerListenAddress,
		// use h2c so we can server HTTP/2 w/o TLS
		Handler: h2c.NewHandler(mux, &http2.Server{}),
		// TODO: put timeouts in config if the need to be changed
		ReadHeaderTimeout: time.Second * 5,
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 30,
	}, nil
}

// Start will start the server listen and serve process, it will block.
func Start(srv *http.Server) error {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
