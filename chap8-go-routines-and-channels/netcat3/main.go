package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8800")
	if err != nil {
		log.Fatal(err)
	}
	doneRead := make(chan struct{})
	doneWrite := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		conn.(*net.TCPConn).CloseRead()
		fmt.Println("close read")
		doneRead <- struct{}{} //signal the main go routine
	}()
	go func() {
		mustCopy(conn, os.Stdin)
		conn.(*net.TCPConn).CloseWrite()
		fmt.Println("close write")
		doneWrite <- struct{}{}
	}()
	select {
	case <-doneRead:
		conn.(*net.TCPConn).CloseWrite()
		fmt.Println("close write")
	case <-doneWrite:
		<-doneRead
	}
}
func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
