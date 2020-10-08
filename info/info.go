package info

import (
	fmt "fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/context"
)

type Server struct {
}

var socketsToCores map[int]string = make(map[int]string, 0)
var dir string

func init() {
	dir, _ := os.Getwd()
	fmt.Printf("Current dir: %v\n", dir)
	//cmd, err := exec.Command("/bin/sh", dir+"/server/getCores.sh").Output()
	cmd, err := exec.Command("/bin/sh", dir+"/server/getCores.sh").Output()

	if err != nil {
		fmt.Printf("A problem occured while reading core info (lscpu): %v\n", err)
	}
	//fmt.Print(string(cmd))
	for _, line := range strings.Split(string(cmd), "\n") {
		//fmt.Println(line)
		sc := strings.Split(line, " ")
		//fmt.Println(len(sc))
		if len(sc) > 1 {
			socketId, _ := strconv.Atoi(sc[0])
			cores := sc[1]
			socketsToCores[socketId] = cores
		}

	}
	//fmt.Println(socketsToCores)
}

// func (s *Server) Pin(ctx context.Context, in *Info) (*AgentMessage, error) {
// 	log.Printf("Agent needs to put %v on Socket %v", in.Pod.PodName, in.Socket.SocketId)
// 	dir, _ := os.Getwd()
// 	fmt.Printf("Current dir: %v\n", dir)
// 	cmdGo := &exec.Cmd{
// 		Path: dir + "/server/pin.sh",
// 		Args: []string{dir + "/server/pin.sh", in.Pod.PodName, socketsToCores[int(in.Socket.SocketId)]},
// 	}

// 	out, err := cmdGo.Output()
// 	if err != nil {
// 		fmt.Printf("A problem occured while pinning: %v", err)
// 	}
// 	return &AgentMessage{Body: string(out)}, nil
// }

func (s *Server) Pin(ctx context.Context, in *Info) (*AgentMessage, error) {
	addPodToMap(in.Pod.PodName, socketsToCores[int(in.Socket.SocketId)])
	return &AgentMessage{Body: "Pod pinning is scheduled"}, nil
}
