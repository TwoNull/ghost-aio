package events

import (
	"bufio"
	"github.com/0xdarktwo/ghost-aio/internal/telnet"
	"io"
	"log"
	"sync"
	"time"

	"github.com/0xdarktwo/ghost-aio/internal/pathing"
	tel "github.com/aprice/telnet"
)

func EventListener(wg *sync.WaitGroup, conn *tel.Connection) {
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
		if line == "Can't use cheat command getpos in multiplayer, unless the server has sv_cheats set to 1." {
			log.Println("Testing Radar")
			testGetRadar(conn)
		}
	}
	log.Println("Event Listener Terminated")
}

func consoleReader(out chan []byte, conn *tel.Connection) {
	rdr := bufio.NewReader(conn)
	var line, cont []byte
	var prefix bool
	var err error
	for {
		if telnet.IsBlocked() {
			break
		}
		if !telnet.IsBlocked() {
			telnet.Block()
			log.Println("Reading from Console")
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
			telnet.Unblock()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func testGetRadar(conn *tel.Connection) {
	time.Sleep(5 * time.Second)
	pathing.GetInGameRadar(conn)
}
