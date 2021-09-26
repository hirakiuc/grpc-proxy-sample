package handler

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	pb "github.com/hirakiuc/grpc-proxy-sample/proto/A"
)

type handlerImpl struct {
	log *zap.Logger

	pb.UnimplementedBackendAServer
}

func NewHandler(log *zap.Logger) pb.BackendAServer {
	return &handlerImpl{
		log: log,
	}
}

func (h *handlerImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s\n", req.GetName()),
	}, nil
}
