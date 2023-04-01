package sourceio

import (
	"bytes"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"log"
	"time"
)

var Input = make([]string, 0)
var Output = make(chan byte)

var SourceCaller telnet.Caller = SourceCallerType{}

type SourceCallerType struct{}

func (caller SourceCallerType) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	callTelnet(&Input, Output, w, r)
}

func QueueMessage(message string) {
	Input = append(Input, message)
}

func callTelnet(i *[]string, o chan byte, w telnet.Writer, r telnet.Reader) {
	var readBuffer [1]byte
	var readP = readBuffer[:]

	var writeBuffer bytes.Buffer
	var writeP []byte
	var crlfBuffer [2]byte = [2]byte{'\r', '\n'}
	var crlf = crlfBuffer[:]
	var n int
	var err error
	log.Println("Calling Telnet")
	for {
		n, err = r.Read(readP)
		if n <= 0 && nil != err {
			break
		}
		if n > 0 {
			o <- readP[0]
		}

		log.Println(len(*i))

		for len(*i) > 0 {
			log.Println("WRITING MESSAGE")
			writeBuffer.Write([]byte((*i)[0]))
			writeBuffer.Write(crlf)
			writeP = writeBuffer.Bytes()

			log.Println(writeP)
			n, err := oi.LongWrite(w, writeP)
			if nil != err {
				break
			}
			if expected, actual := int64(len(writeP)), n; expected != actual {
				log.Fatalf("Transmission problem: tried sending %d bytes, but actually only sent %d bytes.", expected, actual)
			}
			log.Println("Write Success")
			writeBuffer.Reset()
			(*i) = (*i)[1:]
		}
		time.Sleep(3 * time.Millisecond)
	}
	log.Println("EOF")
}
