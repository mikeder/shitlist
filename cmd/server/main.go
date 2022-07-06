package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mikeder/shitlist/internal/handlers"

	shitlistv1 "github.com/mikeder/shitlist/pkg/go/shitlist/v1"
	"github.com/mikeder/shitlist/pkg/go/shitlist/v1/shitlistv1connect"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ShitlistServer struct{}

func (s *ShitlistServer) Greet(
	ctx context.Context,
	req *connect.Request[shitlistv1.GreetRequest]) (*connect.Response[shitlistv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&shitlistv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Shitlist-Version", "v1")
	return res, nil
}

func main() {
	mux := http.NewServeMux()

	// register file handlers
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

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
	shitlistsrv := &ShitlistServer{}
	path, handler := shitlistv1connect.NewShitlistServiceHandler(shitlistsrv)
	mux.Handle(path, handler)

	// start serving traffic
	if err := http.ListenAndServe(
		":10000",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		panic(err)
	}
}
