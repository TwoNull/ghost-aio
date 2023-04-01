package events

import (
	"github.com/0xdarktwo/ghost-aio/internal/sourceio"
	"log"
	"sync"
	"time"
)

func EventListener(wg *sync.WaitGroup) {
	go sendInput()
	for outByte := range sourceio.Output {
		log.Println("     BYTE: " + string(outByte))
	}
	log.Println("Event Listener Terminated")
}

func sendInput() {
	time.Sleep(1000 * time.Millisecond)
	sourceio.QueueMessage("name")
}
