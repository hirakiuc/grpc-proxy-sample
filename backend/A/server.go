package a

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hirakiuc/grpc-proxy-sample/backend/A/handler"
	pb "github.com/hirakiuc/grpc-proxy-sample/proto/A"
)

type Config struct {
	Logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_zap.UnaryServerInterceptor(conf.Logger)),
	}

	handlers := handler.NewHandler(conf.Logger)

	return createServer(handlers, opts...), nil
}

func createServer(handlers pb.BackendAServer, opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	pb.RegisterBackendAServer(srv, handlers)
	reflection.Register(srv)

	return srv
}
