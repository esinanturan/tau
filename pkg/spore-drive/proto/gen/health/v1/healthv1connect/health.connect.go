// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: health/v1/health.proto

package healthv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/taubyte/tau/pkg/spore-drive/proto/gen/health/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// HealthServiceName is the fully-qualified name of the HealthService service.
	HealthServiceName = "health.v1.HealthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthServicePingProcedure is the fully-qualified name of the HealthService's Ping RPC.
	HealthServicePingProcedure = "/health.v1.HealthService/Ping"
	// HealthServiceSupportsProcedure is the fully-qualified name of the HealthService's Supports RPC.
	HealthServiceSupportsProcedure = "/health.v1.HealthService/Supports"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	healthServiceServiceDescriptor        = v1.File_health_v1_health_proto.Services().ByName("HealthService")
	healthServicePingMethodDescriptor     = healthServiceServiceDescriptor.Methods().ByName("Ping")
	healthServiceSupportsMethodDescriptor = healthServiceServiceDescriptor.Methods().ByName("Supports")
)

// HealthServiceClient is a client for the health.v1.HealthService service.
type HealthServiceClient interface {
	Ping(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error)
	Supports(context.Context, *connect.Request[v1.SupportsRequest]) (*connect.Response[v1.Empty], error)
}

// NewHealthServiceClient constructs a client for the health.v1.HealthService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) HealthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &healthServiceClient{
		ping: connect.NewClient[v1.Empty, v1.Empty](
			httpClient,
			baseURL+HealthServicePingProcedure,
			connect.WithSchema(healthServicePingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		supports: connect.NewClient[v1.SupportsRequest, v1.Empty](
			httpClient,
			baseURL+HealthServiceSupportsProcedure,
			connect.WithSchema(healthServiceSupportsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// healthServiceClient implements HealthServiceClient.
type healthServiceClient struct {
	ping     *connect.Client[v1.Empty, v1.Empty]
	supports *connect.Client[v1.SupportsRequest, v1.Empty]
}

// Ping calls health.v1.HealthService.Ping.
func (c *healthServiceClient) Ping(ctx context.Context, req *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error) {
	return c.ping.CallUnary(ctx, req)
}

// Supports calls health.v1.HealthService.Supports.
func (c *healthServiceClient) Supports(ctx context.Context, req *connect.Request[v1.SupportsRequest]) (*connect.Response[v1.Empty], error) {
	return c.supports.CallUnary(ctx, req)
}

// HealthServiceHandler is an implementation of the health.v1.HealthService service.
type HealthServiceHandler interface {
	Ping(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error)
	Supports(context.Context, *connect.Request[v1.SupportsRequest]) (*connect.Response[v1.Empty], error)
}

// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthServiceHandler(svc HealthServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	healthServicePingHandler := connect.NewUnaryHandler(
		HealthServicePingProcedure,
		svc.Ping,
		connect.WithSchema(healthServicePingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	healthServiceSupportsHandler := connect.NewUnaryHandler(
		HealthServiceSupportsProcedure,
		svc.Supports,
		connect.WithSchema(healthServiceSupportsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/health.v1.HealthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthServicePingProcedure:
			healthServicePingHandler.ServeHTTP(w, r)
		case HealthServiceSupportsProcedure:
			healthServiceSupportsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthServiceHandler struct{}

func (UnimplementedHealthServiceHandler) Ping(context.Context, *connect.Request[v1.Empty]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("health.v1.HealthService.Ping is not implemented"))
}

func (UnimplementedHealthServiceHandler) Supports(context.Context, *connect.Request[v1.SupportsRequest]) (*connect.Response[v1.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("health.v1.HealthService.Supports is not implemented"))
}