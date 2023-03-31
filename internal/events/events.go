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
	output := make(chan []byte)
	go consoleReader(output, conn)
	for lineBytes := range output {
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
		if line == "Can't use cheat command getpos in multiplayer, unless the server has sv_cheats set to 1." {
			log.Println("Testing Radar")
			go testGetRadar(conn)
		}
	}
	log.Println("Event Listener Terminated")
}

func consoleReader(output chan []byte, conn *tel.Conn) {
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
			output <- line
		}
		if err == io.EOF {
			break
		}
	}
}

func testGetRadar(conn *tel.Conn) {
	time.Sleep(2 * time.Second)
	pathing.GetInGameRadar(conn)
}
