package telnet

import (
	"bufio"
	"log"
	"net"

	"github.com/0xdarktwo/ghost-aio/internal/pathing"
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
		length := len(line)
		if length > 11 && line[length-11:length] == " connected." {
			continue
		}
		if length > 5 && line[0:5] == "Map: " {
			log.Println("Connected to Server on " + line[5:length])
			go pathing.InitPathing(line[5:length])
		}
		if len(line) > 16 && line[0:17] == "CCSGO_BlurTarget" {
			log.Println("Team choice dialogue")
		}
		log.Println(line)
	}
}
