package Client

import (
	"io"
	"log"
	"net"
	"os"
)

var (
	done = make(chan struct{})
)

func WaitDone(){
	<- done
}

func TerminalOutput(conn net.Conn) {
	io.Copy(os.Stdout,conn)
	log.Println("done")
	done <- struct{}{}
}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
