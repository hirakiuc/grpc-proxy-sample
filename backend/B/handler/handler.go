package handler

import (
	"context"
	"fmt"

	pb "github.com/hirakiuc/grpc-proxy-sample/proto/B"

	"go.uber.org/zap"
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
