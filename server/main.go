package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/iwita/node-agent/info"
	"google.golang.org/grpc"
)

func main() {

	// initialize a timer that will check
	timer := time.NewTicker(info.CheckInterval)
	fmt.Printf("Timer was set to %v\n", info.CheckInterval)
	go func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Time to check the map")
				info.CheckMap()
			}
		}
	}()

	lis, err := net.Listen("tcp", ":4242")
	//lis, err := net.Listen("tcp", ":3000")
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
