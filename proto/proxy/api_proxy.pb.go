package proxy

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	A "github.com/hirakiuc/grpc-proxy-sample/proto/A"
)

func RegisterBackendAHandlerWithBackendOption(ctx context.Context, srv *grpc.Server, target string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		// nolint:wrapcheck
		return err
	}

	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", target, cerr)
			}

			return
		}

		go func() {
			<-ctx.Done()

			if cerr := conn.Close(); cerr != nil {
				grpclog.Info("Failed to close conn to %s: %v", target, cerr)
			}
		}()
	}()

	return RegisterBackendAHandlers(ctx, srv, conn)
}

func RegisterBackendAHandlers(ctx context.Context, srv *grpc.Server, conn *grpc.ClientConn) error {
	return RegisterBackendAHandlerWithBackend(ctx, srv, A.NewBackendAClient(conn))
}

func RegisterBackendAHandlerWithBackend(ctx context.Context, srv *grpc.Server, client BackendAClient) error {
	backendAServiceDesc := grpc.ServiceDesc{
		ServiceName: "proxy.BackendA",
		HandlerType: (*BackendAServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler:    _BackendA_SayHello_Handler,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "proxy/api.proto",
	}

	handlers := newBackendAHandlersImpl(client)

	srv.RegisterService(&backendAServiceDesc, handlers)

	return nil
}

// handler implementation of a gRPC server for client.
type backendAHandlersImpl struct {
	client BackendAClient

	UnimplementedBackendAServer
}

func newBackendAHandlersImpl(client BackendAClient) BackendAServer {
	return &backendAHandlersImpl{
		client: client,
	}
}

func (srv *backendAHandlersImpl) SayHello(ctx context.Context, req *A.HelloRequest) (*A.HelloReply, error) {
	ret, err := srv.client.SayHello(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch a response from backend: %w", err)
	}

	return ret, nil
}
