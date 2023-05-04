// Copyright 2022 SaferPlace

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: report/v1/report.proto

package reportconnect

import (
	v1 "api.safer.place/report/v1"
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
	// ReportServiceName is the fully-qualified name of the ReportService service.
	ReportServiceName = "report.v1.ReportService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ReportServiceSendReportProcedure is the fully-qualified name of the ReportService's SendReport
	// RPC.
	ReportServiceSendReportProcedure = "/report.v1.ReportService/SendReport"
)

// ReportServiceClient is a client for the report.v1.ReportService service.
type ReportServiceClient interface {
	// SendReport
	SendReport(context.Context, *connect_go.Request[v1.SendReportRequest]) (*connect_go.Response[v1.SendReportResponse], error)
}

// NewReportServiceClient constructs a client for the report.v1.ReportService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReportServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ReportServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &reportServiceClient{
		sendReport: connect_go.NewClient[v1.SendReportRequest, v1.SendReportResponse](
			httpClient,
			baseURL+ReportServiceSendReportProcedure,
			opts...,
		),
	}
}

// reportServiceClient implements ReportServiceClient.
type reportServiceClient struct {
	sendReport *connect_go.Client[v1.SendReportRequest, v1.SendReportResponse]
}

// SendReport calls report.v1.ReportService.SendReport.
func (c *reportServiceClient) SendReport(ctx context.Context, req *connect_go.Request[v1.SendReportRequest]) (*connect_go.Response[v1.SendReportResponse], error) {
	return c.sendReport.CallUnary(ctx, req)
}

// ReportServiceHandler is an implementation of the report.v1.ReportService service.
type ReportServiceHandler interface {
	// SendReport
	SendReport(context.Context, *connect_go.Request[v1.SendReportRequest]) (*connect_go.Response[v1.SendReportResponse], error)
}

// NewReportServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReportServiceHandler(svc ReportServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ReportServiceSendReportProcedure, connect_go.NewUnaryHandler(
		ReportServiceSendReportProcedure,
		svc.SendReport,
		opts...,
	))
	return "/report.v1.ReportService/", mux
}

// UnimplementedReportServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReportServiceHandler struct{}

func (UnimplementedReportServiceHandler) SendReport(context.Context, *connect_go.Request[v1.SendReportRequest]) (*connect_go.Response[v1.SendReportResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("report.v1.ReportService.SendReport is not implemented"))
}
