package cmd

import (
	"context"

	"github.com/a-hilaly/lab-grpc/pkg/api/v1"
	"github.com/a-hilaly/lab-grpc/pkg/protocol/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{Message: "Hello " + in.Name}, nil
}

// RunServer runs gRPC server and HTTP gateway
func RunServer(port string) error {
	ctx := context.Background()

	return grpc.RunServer(ctx, &server{}, port)
}
