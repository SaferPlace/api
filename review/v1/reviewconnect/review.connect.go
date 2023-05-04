// Copyright 2022 SaferPlace

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: review/v1/review.proto

package reviewconnect

import (
	v1 "api.safer.place/review/v1"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
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
	// ReviewServiceName is the fully-qualified name of the ReviewService service.
	ReviewServiceName = "review.v1.ReviewService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ReviewServiceViewIncidentProcedure is the fully-qualified name of the ReviewService's
	// ViewIncident RPC.
	ReviewServiceViewIncidentProcedure = "/review.v1.ReviewService/ViewIncident"
	// ReviewServiceReviewIncidentProcedure is the fully-qualified name of the ReviewService's
	// ReviewIncident RPC.
	ReviewServiceReviewIncidentProcedure = "/review.v1.ReviewService/ReviewIncident"
	// ReviewServiceIncidentsWithoutReviewProcedure is the fully-qualified name of the ReviewService's
	// IncidentsWithoutReview RPC.
	ReviewServiceIncidentsWithoutReviewProcedure = "/review.v1.ReviewService/IncidentsWithoutReview"
)

// ReviewServiceClient is a client for the review.v1.ReviewService service.
type ReviewServiceClient interface {
	// ViewIncident displays the incident information for review.
	ViewIncident(context.Context, *connect_go.Request[v1.ViewIncidentRequest]) (*connect_go.Response[v1.ViewIncidentResponse], error)
	// ReviewIncident allows the client to send reviewed incident.
	ReviewIncident(context.Context, *connect_go.Request[v1.ReviewIncidentRequest]) (*connect_go.Response[v1.ReviewIncidentResponse], error)
	// IncidentsWithoutReview shows basic information about incidents which
	// have not been reviewed yet.
	IncidentsWithoutReview(context.Context, *connect_go.Request[v1.IncidentsWithoutReviewRequest]) (*connect_go.Response[v1.IncidentsWithoutReviewResponse], error)
}

// NewReviewServiceClient constructs a client for the review.v1.ReviewService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReviewServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ReviewServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &reviewServiceClient{
		viewIncident: connect_go.NewClient[v1.ViewIncidentRequest, v1.ViewIncidentResponse](
			httpClient,
			baseURL+ReviewServiceViewIncidentProcedure,
			opts...,
		),
		reviewIncident: connect_go.NewClient[v1.ReviewIncidentRequest, v1.ReviewIncidentResponse](
			httpClient,
			baseURL+ReviewServiceReviewIncidentProcedure,
			opts...,
		),
		incidentsWithoutReview: connect_go.NewClient[v1.IncidentsWithoutReviewRequest, v1.IncidentsWithoutReviewResponse](
			httpClient,
			baseURL+ReviewServiceIncidentsWithoutReviewProcedure,
			opts...,
		),
	}
}

// reviewServiceClient implements ReviewServiceClient.
type reviewServiceClient struct {
	viewIncident           *connect_go.Client[v1.ViewIncidentRequest, v1.ViewIncidentResponse]
	reviewIncident         *connect_go.Client[v1.ReviewIncidentRequest, v1.ReviewIncidentResponse]
	incidentsWithoutReview *connect_go.Client[v1.IncidentsWithoutReviewRequest, v1.IncidentsWithoutReviewResponse]
}

// ViewIncident calls review.v1.ReviewService.ViewIncident.
func (c *reviewServiceClient) ViewIncident(ctx context.Context, req *connect_go.Request[v1.ViewIncidentRequest]) (*connect_go.Response[v1.ViewIncidentResponse], error) {
	return c.viewIncident.CallUnary(ctx, req)
}

// ReviewIncident calls review.v1.ReviewService.ReviewIncident.
func (c *reviewServiceClient) ReviewIncident(ctx context.Context, req *connect_go.Request[v1.ReviewIncidentRequest]) (*connect_go.Response[v1.ReviewIncidentResponse], error) {
	return c.reviewIncident.CallUnary(ctx, req)
}

// IncidentsWithoutReview calls review.v1.ReviewService.IncidentsWithoutReview.
func (c *reviewServiceClient) IncidentsWithoutReview(ctx context.Context, req *connect_go.Request[v1.IncidentsWithoutReviewRequest]) (*connect_go.Response[v1.IncidentsWithoutReviewResponse], error) {
	return c.incidentsWithoutReview.CallUnary(ctx, req)
}

// ReviewServiceHandler is an implementation of the review.v1.ReviewService service.
type ReviewServiceHandler interface {
	// ViewIncident displays the incident information for review.
	ViewIncident(context.Context, *connect_go.Request[v1.ViewIncidentRequest]) (*connect_go.Response[v1.ViewIncidentResponse], error)
	// ReviewIncident allows the client to send reviewed incident.
	ReviewIncident(context.Context, *connect_go.Request[v1.ReviewIncidentRequest]) (*connect_go.Response[v1.ReviewIncidentResponse], error)
	// IncidentsWithoutReview shows basic information about incidents which
	// have not been reviewed yet.
	IncidentsWithoutReview(context.Context, *connect_go.Request[v1.IncidentsWithoutReviewRequest]) (*connect_go.Response[v1.IncidentsWithoutReviewResponse], error)
}

// NewReviewServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReviewServiceHandler(svc ReviewServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ReviewServiceViewIncidentProcedure, connect_go.NewUnaryHandler(
		ReviewServiceViewIncidentProcedure,
		svc.ViewIncident,
		opts...,
	))
	mux.Handle(ReviewServiceReviewIncidentProcedure, connect_go.NewUnaryHandler(
		ReviewServiceReviewIncidentProcedure,
		svc.ReviewIncident,
		opts...,
	))
	mux.Handle(ReviewServiceIncidentsWithoutReviewProcedure, connect_go.NewUnaryHandler(
		ReviewServiceIncidentsWithoutReviewProcedure,
		svc.IncidentsWithoutReview,
		opts...,
	))
	return "/review.v1.ReviewService/", mux
}

// UnimplementedReviewServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReviewServiceHandler struct{}

func (UnimplementedReviewServiceHandler) ViewIncident(context.Context, *connect_go.Request[v1.ViewIncidentRequest]) (*connect_go.Response[v1.ViewIncidentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("review.v1.ReviewService.ViewIncident is not implemented"))
}

func (UnimplementedReviewServiceHandler) ReviewIncident(context.Context, *connect_go.Request[v1.ReviewIncidentRequest]) (*connect_go.Response[v1.ReviewIncidentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("review.v1.ReviewService.ReviewIncident is not implemented"))
}

func (UnimplementedReviewServiceHandler) IncidentsWithoutReview(context.Context, *connect_go.Request[v1.IncidentsWithoutReviewRequest]) (*connect_go.Response[v1.IncidentsWithoutReviewResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("review.v1.ReviewService.IncidentsWithoutReview is not implemented"))
}
