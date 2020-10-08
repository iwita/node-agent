package info

import (
	fmt "fmt"
	"os"
	"os/exec"
	"time"
)

type podInfo struct {
	Cores   string
	Started time.Time
}

var CheckInterval time.Duration
var Timeout time.Duration
var podToInfo map[string]podInfo

func init() {
	podToInfo = make(map[string]podInfo)
	CheckInterval = time.Duration(10) * time.Second
	Timeout = time.Duration(100) * time.Second
}

func addPodToMap(pod, cores string) {
	podToInfo[pod] = podInfo{
		Cores:   cores,
		Started: time.Now(),
	}
	return
}

func CheckMap() {
	for k, v := range podToInfo {

		// check for timeout
		if time.Since(v.Started) > Timeout {
			fmt.Println("Runned out of time")
			delete(podToInfo, k)
		}

		// try to pin pod 'k' to cores 'v'
		ok, err := tryToPin(k, v.Cores)
		if err != nil {
			fmt.Printf("An error occured in pin.sh script: %v\n", err)
			fmt.Printf("Remaining time: %v\n", Timeout-time.Since(v.Started))
		}
		if !ok {
			fmt.Printf("Not able to pin %v to %v\n", k, v.Cores)
			// decrease the timeout limit
			//v.Timeout -= time.Second * CheckInterval
		} else {
			// if pinning is successful, remove the pod from the map
			fmt.Printf("Successfully pinned %v to %v (time remaining: %v)\n", k, v.Cores, Timeout-time.Since(v.Started))
			delete(podToInfo, k)
		}

	}
	fmt.Printf("Remaining map: %v\n", podToInfo)
}

func tryToPin(pod, cores string) (bool, error) {
	dir, _ := os.Getwd()
	fmt.Printf("Current dir: %v\n", dir)
	cmdGo := &exec.Cmd{
		Path: dir + "/server/pin.sh",
		Args: []string{dir + "/server/pin.sh", pod, cores},
	}

	out, err := cmdGo.Output()
	if err != nil {
		fmt.Printf("A problem occured while pinning: %v\n", err)
		return false, err
	}
	fmt.Println(string(out))
	return true, nil
}

func removePod(pod string) {
	delete(podToInfo, pod)
}
