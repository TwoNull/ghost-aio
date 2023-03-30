package events

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"

	"github.com/0xdarktwo/ghost-aio/internal/pathing"
)

func EventReader(wg *sync.WaitGroup, conn *net.Conn) {
	defer wg.Done()
	defer (*conn).Close()
	scanner := bufio.NewScanner(*conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
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
	log.Println("Eventreader Terminated")
}

func testGetRadar(conn *net.Conn) {
	time.Sleep(5 * time.Second)
	pathing.GetInGameRadar(conn)
}
