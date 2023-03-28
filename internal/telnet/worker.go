package telnet

import (
	"bufio"
	"log"
	"net"
)

func ReadWorker(conn net.Conn, consoleOut chan string) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for {
		scanner.Scan()
		consoleOut <- scanner.Text()
	}
}

func EventWorker(consoleOut chan string) {
	for line := range consoleOut {
		if len(line) > 12 && line[0:14] == "Connected to " {
			log.Println("Joined Matchmaking Server")
		}
		if len(line) > 16 && line[0:17] == "CCSGO_BlurTarget" {
			log.Println("Team choice dialogue")
		}
		log.Println(line)
	}
}
