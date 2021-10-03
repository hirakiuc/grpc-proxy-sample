package proxy

import (
	"context"
	"fmt"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	proxy_pb "github.com/hirakiuc/grpc-proxy-sample/proto/proxy"
)

type Config struct {
	Logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_zap.UnaryServerInterceptor(conf.Logger)),
	}

	return createServer(opts...)
}

func createServer(opts ...grpc.ServerOption) (*grpc.Server, error) {
	srv := grpc.NewServer(opts...)

	ctx := context.Background()
	clientOptions := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	err := proxy_pb.RegisterBackendAHandlerWithBackendOption(ctx, srv, "localhost:50051", clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to register handler: %w", err)
	}

	reflection.Register(srv)

	return srv, nil
}
