package A

import (
	"github.com/hirakiuc/proxy-sample/backend/A/handler"

	pb "github.com/hirakiuc/proxy-sample/proto/A"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	opts := []grpc.ServerOption{}

	handlers := handler.NewHandler(conf.Logger)

	return createServer(handlers, opts...), nil
}

func createServer(handlers pb.BackendAServer, opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	pb.RegisterBackendAServer(srv, handlers)
	reflection.Register(srv)

	return srv
}
