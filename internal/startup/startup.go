package startup

import (
	"log"
	"os/exec"
	"runtime"
	"telnet"
)

func StartApp(id string, steamdir string, telnetport string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd = exec.Command(steamdir, "-applaunch "+id, "-windowed", "-novid", "-nojoy", "-noborder", "-w 960", "-h 540", "-x 0", "-y 0", "-refresh 30", "-netconport "+telnetport)
	} else {
		log.Fatal("Using Unsupported Operating System!")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error starting app (Steam Not in Default Directory?)")
	}
	conn := telnet.InitConnection("localhost:" + telnetport)

}
