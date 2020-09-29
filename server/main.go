package main

import (
	"fmt"
	"log"
	"net"

	"github.com/iwita/node-agent/info"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	fmt.Printf("Stef Tsitsipas started Serving on %v\n", lis.Addr().String())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := info.Server{}

	grpcServer := grpc.NewServer()

	info.RegisterTasksetAdvisorServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
