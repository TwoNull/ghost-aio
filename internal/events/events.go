package events

import (
	"bufio"
	"io"
	"log"
	"sync"
	"time"

	tel "github.com/reiver/go-telnet"

	"github.com/0xdarktwo/ghost-aio/internal/pathing"
)

func EventListener(wg *sync.WaitGroup, conn *tel.Conn) {
	out := make(chan []byte)
	go consoleReader(out, conn)
	for lineBytes := range out {
		line := string(lineBytes)
		log.Println(line)
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
	log.Println("Event Listener Terminated")
}

func consoleReader(out chan []byte, conn *tel.Conn) {
	rdr := bufio.NewReader(conn)
	var line, cont []byte
	var prefix bool
	var err error
	for {
		line, prefix, err = rdr.ReadLine()
		for prefix && err == nil {
			cont, prefix, err = rdr.ReadLine()
			line = append(line, cont...)
		}
		if line != nil {
			out <- line
		}
		if err == io.EOF {
			break
		}
	}
}

func testGetRadar(conn *tel.Conn) {
	time.Sleep(5 * time.Second)
	pathing.GetInGameRadar(conn)
}
