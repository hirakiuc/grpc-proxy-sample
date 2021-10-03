package proxy

import (
	"context"
	"strings"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_proxy "github.com/mwitkow/grpc-proxy/proxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/hirakiuc/grpc-proxy-sample/proto/proxy"
)

type Config struct {
	Logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	director := createDirector(conf)

	opts := []grpc.ServerOption{
		// https://pkg.go.dev/google.golang.org/grpc#CustomCodec
		// Deprecated: register codecs using encoding.RegisterCodec.
		// The server will automatically use registered codecs based on the incoming requests' headers.
		// See also https://github.com/grpc/grpc-go/blob/master/Documentation/encoding.md#using-a-codec.
		// Will be supported throughout 1.x.
		// grpc.CustomCodec(grpc_proxy.Codec()),

		grpc.UnaryInterceptor(grpc_zap.UnaryServerInterceptor(conf.Logger)),

		grpc.UnknownServiceHandler(grpc_proxy.TransparentHandler(director)),
	}

	return createServer(director, opts...), nil
}

func createDirector(conf *Config) grpc_proxy.StreamDirector {
	log := conf.Logger

	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		backends := map[string]string{
			"proxy.BackendA": "localhost:50051",
			"proxy.BackendB": "localhost:50052",
		}

		const Sep = "/"

		parts := strings.Split(fullMethodName, Sep)

		serviceName := parts[1]
		methodName := parts[2]

		log.Info(
			"Receive a request",
			zap.String("service", serviceName),
			zap.String("method", methodName),
		)

		addr, ok := backends[serviceName]
		if !ok {
			return nil, nil, status.Errorf(codes.Unimplemented, "Unsupported service")
		}

		md, ok := metadata.FromIncomingContext(ctx)

		// Copy the inbound metadata explicitly.
		// nolint:govet
		outCtx, _ := context.WithCancel(ctx)
		// defer cancel()

		outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())

		if ok {
			conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())

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

func createServer(director grpc_proxy.StreamDirector, opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	backendDescs := []grpc.ServiceDesc{
		pb.BackendA_ServiceDesc,
		pb.BackendB_ServiceDesc,
	}

	registerBackendServices(srv, director, backendDescs)
	reflection.Register(srv)

	return srv
}
