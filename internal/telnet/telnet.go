package telnet

import (
	"fmt"
	"log"
	"net"
	"time"
)

func InitConnection(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		for i := 1; err == nil || i < 11; i++ {
			time.Sleep(3000 * time.Millisecond)
			log.Println("Retryimg Connection | Attempt " + fmt.Sprint(i) + " of 10")
			conn, err = net.Dial("tcp", address)
		}
	}
	if err != nil {
		log.Fatal("Source Telnet server did not respond within 30 seconds.")
	}
	log.Println("Connected to CS:GO Console")
	return conn
}
