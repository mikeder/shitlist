package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	shitlistv1 "shitlist/gen/shitlist/v1"
	"shitlist/gen/shitlist/v1/shitlistv1connect"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ShitlistServer struct{}

func (s *ShitlistServer) Shitlist(
	ctx context.Context,
	req *connect.Request[shitlistv1.ShitlistRequest],
) (*connect.Response[shitlistv1.ShitlistResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&shitlistv1.ShitlistResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Shitlist-Version", "v1")
	return res, nil
}

func main() {
	shitlistsrv := &ShitlistServer{}
	mux := http.NewServeMux()
	path, handler := shitlistv1connect.NewShitlistServiceHandler(shitlistsrv)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
