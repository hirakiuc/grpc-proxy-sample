package proxy

import (
	"context"
	"strings"

	pb "github.com/hirakiuc/proxy-sample/proto/proxy"

	grpc_proxy "github.com/mwitkow/grpc-proxy/proxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Config struct {
	Logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	opts := []grpc.ServerOption{
		// https://pkg.go.dev/google.golang.org/grpc#CustomCodec
		// Deprecated: register codecs using encoding.RegisterCodec.
		// The server will automatically use registered codecs based on the incoming requests' headers.
		// See also https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md#using-a-codec.
		// Will be supported throughout 1.x.
		// grpc.CustomCodec(grpc_proxy.Codec()),
	}

	return createServer(opts...), nil
}

func createDirector() grpc_proxy.StreamDirector {
	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		backends := map[string]string{
			"/proxy.BackendA": "backend_a",
			"/proxy.BackendB": "backend_b",
		}

		const Sep = "/"

		parts := strings.Split(fullMethodName, Sep)

		addr, ok := backends[parts[0]]
		if !ok {
			return nil, nil, status.Errorf(codes.Unimplemented, "Unsupported service")
		}

		md, ok := metadata.FromIncomingContext(ctx)

		// Copy the inbound metadata explicitly.
		outCtx, _ := context.WithCancel(ctx)
		outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())

		if ok {
			conn, err := grpc.DialContext(ctx, addr, grpc.WithCodec(grpc_proxy.Codec()))

			// nolint:wrapcheck
			return outCtx, conn, err
		}

		return nil, nil, status.Errorf(codes.Unimplemented, "Unknown method")
	}
}

func registerBackendServices(srv *grpc.Server, director grpc_proxy.StreamDirector, backendDescs []grpc.ServiceDesc) {
	for _, desc := range backendDescs {
		// Extract backend methods
		methods := make([]string, len(desc.Methods))
		for i, v := range desc.Methods {
			methods[i] = v.MethodName
		}

		grpc_proxy.RegisterService(srv, director, desc.ServiceName, methods...)
	}
}

func createServer(opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	director := createDirector()

	backendDescs := []grpc.ServiceDesc{
		pb.BackendA_ServiceDesc,
		pb.BackendB_ServiceDesc,
	}

	registerBackendServices(srv, director, backendDescs)

	return srv
}
