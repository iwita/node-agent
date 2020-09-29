package info

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
}

func (s *Server) Pin(ctx context.Context, in *Info) (*AgentMessage, error) {
	log.Printf("Agent needs to put %v on Socket %v", in.Pod.PodName, in.Socket.SocketId)
	return &AgentMessage{Body: "Everything ok", Date: timestamppb.New(time.Now())}, nil
}
