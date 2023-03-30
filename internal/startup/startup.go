package startup

import (
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/0xdarktwo/ghost-aio/internal/telnet"
	"github.com/0xdarktwo/ghost-aio/internal/window"
)

var wg sync.WaitGroup

func Run(id string, telnetport string, width string, height string) {
	processID := startApp(id, telnetport, width, height)
	conn := telnet.InitConnection("127.0.0.1:" + telnetport)
	consoleOut := make(chan string)
	window.SetWindowBounds(processID)
	wg.Add(1)
	go telnet.ReadWorker(conn, consoleOut)
	go telnet.EventWorker(consoleOut)
	wg.Wait()
}

func startApp(id string, telnetport string, width string, height string) int32 {
	steamdir := os.Getenv("STEAM_PATH")
	procName := os.Getenv("PROCESS_NAME")
	cmd := exec.Command(steamdir, "-applaunch", id, "-windowed", "-novid", "-nojoy", "-noborder", "-w", width, "-h", height, "-x", "0", "-y", "0", "-refresh", "30", "-netconport", telnetport)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error starting app (Steam Not in Default Directory?)")
	}
	processes := window.CheckLaunch(procName)
	if len(processes) == 0 {
		log.Fatal("No CS:GO Instances Found")
	}
	return processes[0]
}

/*func initMM(processID int32) error {
	for {
		initMMStep(processID)
		time.Sleep(1000 * time.Millisecond)
	}
}

func initMMStep(processID int32) error {
	log.Println(window.GetPlayButton(processID))
	return nil
}*/
