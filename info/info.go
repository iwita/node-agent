package info

import (
	fmt "fmt"
	"log"
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
	cmd, err := exec.Command("/bin/sh", "getCores.sh").Output()
	if err != nil {
		fmt.Printf("A problem occured while reading core info (lscpu): %v", err)
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

func (s *Server) Pin(ctx context.Context, in *Info) (*AgentMessage, error) {
	log.Printf("Agent needs to put %v on Socket %v", in.Pod.PodName, in.Socket.SocketId)
	//ListContainer(in.Pod.PodName, int(in.Socket.SocketId))
	dir, _ := os.Getwd()
	//fmt.Printf("Pod name: %v\n", in.Pod.PodName)
	//fmt.Printf("Current dir: %v\n", dir)
	cmdGo := &exec.Cmd{
		//Path: dir + "/server/pin.sh",
		Path: "/server/pin.sh",
//		Args: []string{dir + "/server/pin.sh", in.Pod.PodName, socketsToCores[int(in.Socket.SocketId)]},
		Args: []string{"/server/pin.sh", in.Pod.PodName, socketsToCores[int(in.Socket.SocketId)]},
	}

	//fmt.Println(cmdGo.String())
	out, err := cmdGo.Output()
	//_, err := exec.Command("/bin/sh", dir+"/server/pin.sh"+" "+in.Pod.PodName+" "+socketsToCores[int(in.Socket.SocketId)]).Output()
	if err != nil {
		fmt.Printf("A problem occured while pinning: %v", err)
	}
	return &AgentMessage{Body: string(out)}, nil
}

// func ListContainer(podName string, socketId int) error {
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		panic(err)
// 	}
// 	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	if len(containers) > 0 {
// 		var containerName, containerID string

// 		for _, container := range containers {

// 			// find the correct container
// 			if strings.Contains(container.Names[0], podName) && strings.HasPrefix(container.Names[0], "k8s_server") {
// 				//if strings.Contains(container.Names[0], podName) {
// 				containerName = container.Names[0][1:]
// 				containerID = container.ID
// 				fmt.Printf("Found container: %v, ID: %v\n", containerName, containerID)
// 				break
// 			}
// 		}
// 		// docker update
// 		_, err := cli.ContainerUpdate(context.Background(), containerID, container.UpdateConfig{Resources: container.Resources{CpusetCpus: socketsToCores[socketId]}})
// 		if err != nil {
// 			fmt.Printf("Couldn't apply cpuset in container with ID: %v, %v\n", containerID, err)
// 		}
// 	} else {
// 		fmt.Println("There are no containers running")
// 	}
// 	return nil
// }
