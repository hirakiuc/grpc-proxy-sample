package B

import (
	"github.com/hirakiuc/proxy-sample/backend/B/handler"

	pb "github.com/hirakiuc/proxy-sample/proto/B"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	opts := []grpc.ServerOption{
		// TODO
	}

	handlers := handler.NewHandler(conf.Logger)

	return createServer(handlers, opts...), nil
}

func createServer(handlers pb.BackendBServer, opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	pb.RegisterBackendBServer(srv, handlers)
	reflection.Register(srv)

	return srv
}
