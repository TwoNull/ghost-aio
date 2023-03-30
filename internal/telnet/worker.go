package telnet

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/0xdarktwo/ghost-aio/internal/pathing"
)

var wg sync.WaitGroup

var conn net.Conn
var consoleOut = make(chan string)
var terminated bool = false

func Init(address string) {
	conn, err := net.Dial("tcp", address)
	if conn == nil {
		for i := 1; i < 11; i++ {
			time.Sleep(6000 * time.Millisecond)
			log.Println("Connecting to Client | Attempt " + fmt.Sprint(i) + " of 10")
			conn, err = net.Dial("tcp", address)
			if conn != nil {
				break
			}
		}
	}
	if err != nil {
		log.Fatal("Source Telnet server did not respond within 60 seconds.")
	}
	log.Println("Connected to CS:GO Console")
}

func EventWorker() {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for !terminated {
		scanner.Scan()
		line := scanner.Text()
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
	}
	defer wg.Done()
}

func Write(message string) {

}
