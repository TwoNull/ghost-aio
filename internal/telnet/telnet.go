package telnet

import (
	"log"
	"net"
	"time"
)

func Connect(address string) net.Conn {
	connection, err := net.Dial("tcp", address)
	if connection == nil {
		for i := 1; i < 11; i++ {
			time.Sleep(6000 * time.Millisecond)
			connection, err = net.Dial("tcp", address)
			if connection != nil {
				break
			}
		}
	}
	if err != nil {
		log.Fatal("Source Telnet server did not respond within 60 seconds.")
	}
	return connection
}

func Write(w *net.Conn, data string) int {
	buf := []byte(data + "\n")
	n, err := (*w).Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
