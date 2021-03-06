package main

import (
	"context"
	"log"

	"github.com/iwita/node-agent/info"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":4242", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := info.NewTasksetAdvisorClient(conn)

	response, err := c.Pin(context.Background(), &info.Info{
		Socket: &info.SocketType{SocketId: 0},
		Pod:    &info.PodType{PodName: "golang"},
	})

	if err != nil {
		log.Fatalf("Error when calling Pin remotely: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
