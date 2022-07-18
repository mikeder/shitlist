package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"

	"github.com/mikeder/shitlist/internal/handlers"

	shitlistv1 "github.com/mikeder/shitlist/pkg/go/shitlist/v1"
	"github.com/mikeder/shitlist/pkg/go/shitlist/v1/shitlistv1connect"

	"github.com/bufbuild/connect-go"
	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var clicks = make(map[string]int64)

type ShitlistServer struct {
	clickMux *sync.Mutex
	version  string
}

func (s *ShitlistServer) Greet(
	ctx context.Context,
	req *connect.Request[shitlistv1.GreetRequest]) (*connect.Response[shitlistv1.GreetResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, err
	}
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&shitlistv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Shitlist-Version", s.version)
	return res, nil
}

func (s *ShitlistServer) Click(
	ctx context.Context,
	req *connect.Request[shitlistv1.ClickRequest]) (*connect.Response[shitlistv1.ClickResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, err
	}

	uid := req.Msg.UserId

	s.clickMux.Lock()
	clicks[uid]++
	s.clickMux.Unlock()

	res := connect.NewResponse(&shitlistv1.ClickResponse{
		Clicks: clicks[uid],
	})
	res.Header().Set("Shitlist-Version", s.version)
	return res, nil
}

func (s *ShitlistServer) Leaders(
	ctx context.Context,
	req *connect.Request[shitlistv1.LeadersRequest]) (*connect.Response[shitlistv1.LeadersResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, err
	}

	var clickers, leaders []*shitlistv1.Clicker
	// create a slice of clickers
	for u, c := range clicks {
		clickers = append(clickers, &shitlistv1.Clicker{
			UserId: u,
			Clicks: c,
		})
	}

	// sort clickers highest to lowest
	sort.Slice(clickers, func(i, j int) bool {
		return clickers[i].Clicks > clickers[j].Clicks
	})

	numLeaders := 10
	if len(clickers) < numLeaders {
		numLeaders = len(clickers)
	}
	leaders = clickers[:numLeaders]
	res := connect.NewResponse(&shitlistv1.LeadersResponse{
		TopClickers: leaders,
	})
	res.Header().Set("Shitlist-Version", s.version)
	return res, nil
}

func main() {
	mux := http.NewServeMux()

	// register file handlers
	mux.Handle("/", http.FileServer(http.Dir("templates/")))

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
	shitlistsrv := &ShitlistServer{
		clickMux: new(sync.Mutex),
		version:  "v1",
	}

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
