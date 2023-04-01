package sourceio

import (
	"bufio"
	"bytes"
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"io"
	"log"
	"os"
	"time"
)

var Input = make(chan string)
var Output = make(chan string)

var SourceCaller telnet.Caller = SourceCallerType{}

type SourceCallerType struct{}

func (caller SourceCallerType) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	callTelnet(Input, Output, os.Stderr, w, r)
}

func callTelnet(i chan string, o chan string, stderr io.WriteCloser, w telnet.Writer, r telnet.Reader) {
	var buffer bytes.Buffer
	var p []byte
	var crlfBuffer [2]byte = [2]byte{'\r', '\n'}
	var crlf = crlfBuffer[:]

	var rdr = bufio.NewReader(r)
	var line, cont []byte
	var prefix bool
	var err error
	for {
		line, prefix, err = rdr.ReadLine()
		for prefix && err == nil {
			cont, prefix, err = rdr.ReadLine()
			line = append(line, cont...)
		}
		if line != nil {
			Output <- string(line)
		}
		if err == io.EOF {
			break
		}
		for message := range i {
			buffer.Write([]byte(message))
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
		}
		Input = make(chan string)
		time.Sleep(3 * time.Millisecond)
	}
}
