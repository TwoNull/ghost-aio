package startup

import (
	"github.com/0xdarktwo/ghost-aio/internal/events"
	tel "github.com/aprice/telnet"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/0xdarktwo/ghost-aio/internal/window"
)

var wg sync.WaitGroup

func Run(id, port, width, height string) {
	processID := startApp(id, width, port, height)
	conn, err := tel.Dial("127.0.0.1:" + port)
	if conn == nil {
		for i := 1; i < 11; i++ {
			time.Sleep(6000 * time.Millisecond)
			conn, err = tel.Dial("127.0.0.1:" + port)
			if conn != nil {
				break
			}
		}
	}
	if err != nil {
		log.Fatal("Source Telnet server did not respond within 60 seconds.")
	}
	window.SetWindowBounds(processID)
	wg.Add(1)
	go events.EventListener(&wg, conn)
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
