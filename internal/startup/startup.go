package startup

import (
	"log"
	"os/exec"
	"sync"
	"telnet"
)

var wg sync.WaitGroup

func StartApp(id string, steamdir string, telnetport string) {
	cmd := exec.Command(steamdir, "-applaunch", id, "-windowed", "-novid", "-nojoy", "-noborder", "-w", "1280", "-h", "720", "-x", "0", "-y", "0", "-refresh", "30", "-netconport", telnetport)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error starting app (Steam Not in Default Directory?)")
	}
	conn := telnet.InitConnection("127.0.0.1:2121")
	log.Println("Connected to CS:GO Console")
	wg.Add(1)
	go telnet.TelnetWorker(conn)
	wg.Wait()
}
