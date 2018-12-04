package main

import (
	"context"
	"flag"
	"log"

	"github.com/a-hilaly/lab-grpc/pkg/protocol/grpc"
	"github.com/a-hilaly/lab-grpc/pkg/server"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "55051", "gRPC port to bind")
	flag.Parse()

	ctx := context.Background()

	if err := grpc.RunServer(ctx, server.New(), port); err != nil {
		log.Fatal(err)
	}

}
