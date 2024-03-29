package handlers

import (
	"context"
	"sort"

	"github.com/mikeder/shitlist/internal/database"
	shitlistv1 "github.com/mikeder/shitlist/pkg/go/shitlist/v1"

	"github.com/bufbuild/connect-go"
)

func (a *API) Click(
	ctx context.Context,
	req *connect.Request[shitlistv1.ClickRequest]) (*connect.Response[shitlistv1.ClickResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, err
	}

	uid := req.Msg.UserId
	clicker, err := a.clickStore.AddClick(uid)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&shitlistv1.ClickResponse{
		Clicks: clicker.ClickCount,
	})
	return res, nil
}

func (a *API) Leaders(
	ctx context.Context,
	req *connect.Request[shitlistv1.LeadersRequest]) (*connect.Response[shitlistv1.LeadersResponse], error) {
	if err := req.Msg.Validate(); err != nil {
		return nil, err
	}

	clickers, err := a.clickStore.GetClickers()
	if err != nil {
		return nil, err
	}

	// sort clickers highest to lowest
	sort.Slice(clickers, func(i, j int) bool {
		return clickers[i].ClickCount > clickers[j].ClickCount
	})

	numLeaders := 10
	if len(clickers) < numLeaders {
		numLeaders = len(clickers)
	}

	res := connect.NewResponse(&shitlistv1.LeadersResponse{
		TopClickers: dbClickersToProto(clickers[:numLeaders]),
	})
	return res, nil
}

func dbClickersToProto(dbc []*database.Clicker) []*shitlistv1.Clicker {
	var pc []*shitlistv1.Clicker
	for _, c := range dbc {
		c := c
		pc = append(pc, &shitlistv1.Clicker{
			UserId: c.UserID,
			Clicks: c.ClickCount,
		})
	}
	return pc
}
