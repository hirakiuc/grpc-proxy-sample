package handler

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	pb "github.com/hirakiuc/grpc-proxy-sample/proto/B"
)

type handlerImpl struct {
	log *zap.Logger

	pb.UnimplementedBackendBServer
}

func NewHandler(log *zap.Logger) pb.BackendBServer {
	return &handlerImpl{
		log: log,
	}
}

func (h *handlerImpl) SayBye(ctx context.Context, req *pb.ByeRequest) (*pb.ByeReply, error) {
	return &pb.ByeReply{
		Message: fmt.Sprintf("Bye, %s\n", req.GetName()),
	}, nil
}
