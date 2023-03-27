package telnet

import (
	"bufio"
	"net"
)

func ReadWorker(conn net.Conn, consoleOut chan string) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for {
		scanner.Scan()
		consoleOut <- scanner.Text()
	}
}
