package telnet

import (
	"fmt"
	"log"
	"net"
	"time"
)

func InitConnection(address string) net.Conn {
	time.Sleep(30000 * time.Millisecond)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		for i := 1; err == nil || i < 11; i++ {
			time.Sleep(4000 * time.Millisecond)
			log.Println("Retryimg Connection | Attempt " + fmt.Sprint(i) + " of 10")
			conn, err = net.Dial("tcp", address)
		}
	}
	if err != nil {
		log.Fatal("Source Telnet server did not respond within 70 seconds.")
	}
	return conn
}
