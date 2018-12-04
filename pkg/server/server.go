package server

import (
	"context"

	"github.com/a-hilaly/lab-grpc/pkg/api/v1"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{Message: "Hello " + in.Name}, nil
}

// New return a new server
func New() *server {
	return &server{}
}
