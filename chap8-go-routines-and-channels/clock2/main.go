package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8800, "Determine the port listenning to")
	flag.Parse()
	fmt.Println("port", port)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) //handle each connection concurrently

	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //client disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
