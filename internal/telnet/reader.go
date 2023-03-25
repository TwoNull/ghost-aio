package telnet

import (
	"bufio"
	"log"
	"net"
)

func TelnetWorker(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for {
		scanner.Scan()
		log.Println(scanner.Text())
	}
}
