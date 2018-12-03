package main

import (
	"flag"
	"log"

	"github.com/a-hilaly/lab-grpc/pkg/cmd"
)

func main() {

	var port string
	flag.StringVar(&port, "port", "55051", "gRPC port to bind")
	flag.Parse()

	if err := cmd.RunServer(port); err != nil {
		log.Fatal(err)
	}

}
