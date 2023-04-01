package sourceio

import (
	"bufio"
	"bytes"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"io"
	"log"
	"time"
)

var Input = make([]string, 0)
var Output = make(chan string)

var SourceCaller telnet.Caller = SourceCallerType{}

type SourceCallerType struct{}

func (caller SourceCallerType) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	callTelnet(Input, Output, w, r)
}

func QueueMessage(message string) {
	Input = append(Input, message)
}

func callTelnet(i []string, o chan string, w telnet.Writer, r telnet.Reader) {
	var buffer bytes.Buffer
	var p []byte
	var crlfBuffer [2]byte = [2]byte{'\r', '\n'}
	var crlf = crlfBuffer[:]

	var rdr = bufio.NewReader(r)
	var line, cont []byte
	var prefix bool
	var err error
	log.Println("Calling Telnet")
	for {
		line, prefix, err = rdr.ReadLine()
		log.Println("line read")
		for prefix && err == nil {
			cont, prefix, err = rdr.ReadLine()
			line = append(line, cont...)
		}
		if line != nil {
			o <- string(line)
		}
		if err == io.EOF {
			break
		}

		for len(i) > 0 {
			buffer.Write([]byte(i[0]))
			buffer.Write(crlf)
			p = buffer.Bytes()

			log.Println(p)
			n, err := oi.LongWrite(w, p)
			if nil != err {
				break
			}
			if expected, actual := int64(len(p)), n; expected != actual {
				log.Fatalf("Transmission problem: tried sending %d bytes, but actually only sent %d bytes.", expected, actual)
			}
			log.Println("Write Success")
			buffer.Reset()
			i = i[1:]
		}
		time.Sleep(3 * time.Millisecond)
	}
	log.Println("EOF")
}
