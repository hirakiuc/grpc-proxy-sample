package proxy

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Logger *zap.Logger
}

type Server struct {
	logger *zap.Logger
}

func NewServer(conf *Config) (*grpc.Server, error) {
	opts := grpc.ServerOption{}

	return createServer(opts...)
}

func createServer(opts ...grpc.ServerOptions) (*grpc.Server, error) {
	return grpc.NewServer(opts...)
}
