package telnet

import (
	tel "github.com/aprice/telnet"
	"log"
	"time"
)

var blocked = false

func Write(w *tel.Connection, data string) int {
	for {
		if !IsBlocked() {
			Block()
			buf := []byte(data + "\n")
			log.Println("Writing to Console")
			n, err := (*w).Write(buf)
			if err != nil {
				log.Fatal(err)
			}
			return n
			Unblock()
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func IsBlocked() bool {
	return blocked
}

func Block() {
	blocked = true
}

func Unblock() {
	blocked = false
}
