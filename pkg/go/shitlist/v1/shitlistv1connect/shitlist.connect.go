// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: shitlist/v1/shitlist.proto

package shitlistv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/mikeder/shitlist/pkg/go/shitlist/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ShitlistServiceName is the fully-qualified name of the ShitlistService service.
	ShitlistServiceName = "shitlist.v1.ShitlistService"
)

// ShitlistServiceClient is a client for the shitlist.v1.ShitlistService service.
type ShitlistServiceClient interface {
	// Greet performs a greet action.
	Greet(context.Context, *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error)
	// Click records a click action by a user.
	Click(context.Context, *connect_go.Request[v1.ClickRequest]) (*connect_go.Response[v1.ClickResponse], error)
}

// NewShitlistServiceClient constructs a client for the shitlist.v1.ShitlistService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewShitlistServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ShitlistServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &shitlistServiceClient{
		greet: connect_go.NewClient[v1.GreetRequest, v1.GreetResponse](
			httpClient,
			baseURL+"/shitlist.v1.ShitlistService/Greet",
			opts...,
		),
		click: connect_go.NewClient[v1.ClickRequest, v1.ClickResponse](
			httpClient,
			baseURL+"/shitlist.v1.ShitlistService/Click",
			opts...,
		),
	}
}

// shitlistServiceClient implements ShitlistServiceClient.
type shitlistServiceClient struct {
	greet *connect_go.Client[v1.GreetRequest, v1.GreetResponse]
	click *connect_go.Client[v1.ClickRequest, v1.ClickResponse]
}

// Greet calls shitlist.v1.ShitlistService.Greet.
func (c *shitlistServiceClient) Greet(ctx context.Context, req *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error) {
	return c.greet.CallUnary(ctx, req)
}

// Click calls shitlist.v1.ShitlistService.Click.
func (c *shitlistServiceClient) Click(ctx context.Context, req *connect_go.Request[v1.ClickRequest]) (*connect_go.Response[v1.ClickResponse], error) {
	return c.click.CallUnary(ctx, req)
}

// ShitlistServiceHandler is an implementation of the shitlist.v1.ShitlistService service.
type ShitlistServiceHandler interface {
	// Greet performs a greet action.
	Greet(context.Context, *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error)
	// Click records a click action by a user.
	Click(context.Context, *connect_go.Request[v1.ClickRequest]) (*connect_go.Response[v1.ClickResponse], error)
}

// NewShitlistServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewShitlistServiceHandler(svc ShitlistServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/shitlist.v1.ShitlistService/Greet", connect_go.NewUnaryHandler(
		"/shitlist.v1.ShitlistService/Greet",
		svc.Greet,
		opts...,
	))
	mux.Handle("/shitlist.v1.ShitlistService/Click", connect_go.NewUnaryHandler(
		"/shitlist.v1.ShitlistService/Click",
		svc.Click,
		opts...,
	))
	return "/shitlist.v1.ShitlistService/", mux
}

// UnimplementedShitlistServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedShitlistServiceHandler struct{}

func (UnimplementedShitlistServiceHandler) Greet(context.Context, *connect_go.Request[v1.GreetRequest]) (*connect_go.Response[v1.GreetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("shitlist.v1.ShitlistService.Greet is not implemented"))
}

func (UnimplementedShitlistServiceHandler) Click(context.Context, *connect_go.Request[v1.ClickRequest]) (*connect_go.Response[v1.ClickResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("shitlist.v1.ShitlistService.Click is not implemented"))
}
