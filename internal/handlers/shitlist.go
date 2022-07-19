package handlers

import (
	"context"
	"fmt"
	"log"
	"sort"
	"sync"

	shitlistv1 "github.com/mikeder/shitlist/pkg/go/shitlist/v1"
	"github.com/mikeder/shitlist/pkg/go/shitlist/v1/shitlistv1connect"

	"github.com/bufbuild/connect-go"
)

const version = "v1"

var clicks = make(map[string]int64)

func NewShitlistService() shitlistv1connect.ShitlistServiceHandler {
	return &srv{
		clickMux: new(sync.Mutex),
		version:  version,
	}
}

type srv struct {
	clickMux *sync.Mutex
	version  string
}

func (s *srv) Greet(
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

func (s *srv) Click(
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

func (s *srv) Leaders(
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
