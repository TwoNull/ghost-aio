package telnet

import (
	"log"

	tel "github.com/reiver/go-telnet"
)

func Write(w *tel.Conn, data string) int {
	buf := []byte(data + "\n")
	log.Println("Writing to Console")
	n, err := (*w).Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
