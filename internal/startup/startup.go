package startup

import (
	"log"
	"os/exec"
	"sync"
	"telnet"
	"window"
)

var wg sync.WaitGroup

func Run(id string, steamdir string, processName string, telnetport string) {
	startApp(id, steamdir, telnetport)
	processes := window.CheckLaunch(processName)
	log.Print(processes)
	conn := telnet.InitConnection("127.0.0.1:2121")
	consoleOut := make(chan string)
	wg.Add(1)
	go telnet.ReadWorker(conn, consoleOut)
	wg.Wait()
}

func startApp(id string, steamdir string, telnetport string) {
	cmd := exec.Command(steamdir, "-applaunch", id, "-windowed", "-novid", "-nojoy", "-noborder", "-w", "1280", "-h", "720", "-x", "0", "-y", "0", "-refresh", "30", "-netconport", telnetport)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error starting app (Steam Not in Default Directory?)")
	}
}
