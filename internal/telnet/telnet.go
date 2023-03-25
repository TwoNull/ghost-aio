package telnet

import (
	"fmt"
	"log"
	"time"

	"github.com/reiver/go-telnet"
)

func InitConnection(address string) *telnet.Conn {
	time.Sleep(50000 * time.Millisecond)
	conn, err := telnet.DialTo(address)
	if err != nil {
		for i := 1; err == nil || i < 11; i++ {
			time.Sleep(4000 * time.Millisecond)
			log.Println("Attempting Connection | Try " + fmt.Sprint(i) + " of 10")
			conn, err = telnet.DialTo(address)
		}
	}
	if err != nil {
		log.Fatal("Source Telnet server did not respond within 90 seconds.")
	}
	return conn
}

func GetNextMessage(conn *telnet.Conn) (string, error) {
	var res []byte
	_, err := conn.Read(res)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
