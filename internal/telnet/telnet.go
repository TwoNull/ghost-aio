package telnet

import (
	"fmt"
	"log"
	"net"
	"time"
)

func InitConnection(address string) net.Conn {
	conn, err := net.Dial("tcp", address)
	if conn == nil {
		for i := 1; i < 11; i++ {
			time.Sleep(6000 * time.Millisecond)
			log.Println("Connecting | Attempt " + fmt.Sprint(i) + " of 10")
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
	return conn
}
