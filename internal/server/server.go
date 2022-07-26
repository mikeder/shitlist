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
	mux := http.NewServeMux()

	// register file handlers
	mux.Handle("/", http.FileServer(http.Dir("../../templates/")))

	// OauthGitHub
	mux.HandleFunc("/auth/github/login", handlers.OauthGithubLogin)
	mux.HandleFunc("/auth/github/callback", handlers.OauthGithubCallback)

	// OauthGoogle
	mux.HandleFunc("/auth/google/login", handlers.OauthGoogleLogin)
	mux.HandleFunc("/auth/google/callback", handlers.OauthGoogleCallback)

	// register reflection handlers
	reflector := grpcreflect.NewStaticReflector(
		"shitlist.v1.ShitlistService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	// Many tools still expect the older version of the server reflection API, so
	// most servers should mount both handlers.
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	// If you don't need to support HTTP/2 without TLS (h2c), you can drop
	// x/net/http2 and use http.ListenAndServeTLS instead.

	// register service handlers
	shitlistsrv, err := handlers.NewShitlistService(cfg)
	if err != nil {
		return nil, err
	}

	path, handler := shitlistv1connect.NewShitlistServiceHandler(shitlistsrv)
	mux.Handle(path, handler)

	return &http.Server{
		Addr: cfg.ServerListenAddress,
		//use h2c so we can server HTTP/2 w/o TLS
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
