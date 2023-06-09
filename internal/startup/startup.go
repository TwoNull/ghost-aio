package startup

import (
	"github.com/0xdarktwo/ghost-aio/internal/events"
	"github.com/0xdarktwo/ghost-aio/internal/sourceio"
	tel "github.com/reiver/go-telnet"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/0xdarktwo/ghost-aio/internal/window"
)

var wg sync.WaitGroup
var sourceCaller struct{}

func Run(id, port, width, height string) {
	log.Println("Started")
	processID := startApp(id, width, port, height)
	go establishConnection("127.0.0.1:" + port)
	window.SetWindowBounds(processID)
	wg.Add(1)
	go events.EventListener(&wg)
	wg.Wait()
}

func startApp(id, width, port, height string) int32 {
	steamdir := os.Getenv("STEAM_PATH")
	procName := os.Getenv("PROCESS_NAME")
	cmd := exec.Command(steamdir, "-applaunch", id, "-windowed", "-novid", "-nojoy", "-noborder", "-w", width, "-h", height, "-x", "0", "-y", "0", "-refresh", "30", "-netconport", port)
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

func establishConnection(address string) {
	caller := sourceio.SourceCaller
	err := tel.DialToAndCall(address, caller)
	if err != nil {
		log.Fatal(err)
	}
}
