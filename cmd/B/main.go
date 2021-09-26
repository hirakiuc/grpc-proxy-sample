package main

import (
	"fmt"
	"net"

	backend "github.com/hirakiuc/proxy-sample/backend/B"
	"go.uber.org/zap"
)

const (
	port = "localhost:50052"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		// nolint:forbidigo
		fmt.Printf("failed to create logger: %v", err)

		return
	}

	// nolint:errcheck
	defer logger.Sync()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error(
			"failed to listen a port",
			zap.Error(err),
			zap.String("port", port),
		)

		return
	}

	c := &backend.Config{
		Logger: logger,
	}

	srv, err := backend.NewServer(c)
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))

		return
	}

	logger.Info("start listening", zap.String("port", port))

	err = srv.Serve(lis)
	if err != nil {
		logger.Error("failed to serve gRPC service", zap.Error(err))

		return
	}
}
