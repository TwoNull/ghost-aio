package startup

import (
	"errors"
	"github.com/0xdarktwo/ghost-aio/internal/telnet"
	"github.com/0xdarktwo/ghost-aio/internal/window"
	"log"
	"os/exec"
	"sync"
)

var wg sync.WaitGroup

func Run(id string, steamdir string, processName string, telnetport string) {
	processID, err := startApp(id, steamdir, telnetport, processName)
	if err != nil {
		log.Fatal(err)
	}
	conn := telnet.InitConnection("127.0.0.1:" + telnetport)
	consoleOut := make(chan string)
	err = window.SetWindowBounds(processID)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go telnet.ReadWorker(conn, consoleOut)
	err = initMM(processID)
	wg.Wait()
}

func initMM(processID int32) error {
	log.Println(window.GetTopPixel(processID))
	return nil
}

/*func initMMStep(processID int32) error {

}*/

func startApp(id string, steamdir string, telnetport string, processName string) (int32, error) {
	cmd := exec.Command(steamdir, "-applaunch", id, "-windowed", "-novid", "-nojoy", "-noborder", "-w", "1280", "-h", "720", "-x", "0", "-y", "0", "-refresh", "30", "-netconport", telnetport)
	err := cmd.Run()
	if err != nil {
		return 0, errors.New("Error starting app (Steam Not in Default Directory?)")
	}
	processes := window.CheckLaunch(processName)
	if len(processes) >= 1 {
		return processes[0], nil
	}
	return 0, errors.New("No CS:GO Instances Found")
}
