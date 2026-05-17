package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// an echo server
func main() {
	listener, err := net.Listen("tcp", "localhost:8800")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error occurred while trying to connect to a client: %s", err)
		}
		go handleConn(conn)
	}
}
func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Println(shout)
	fmt.Fprintln(conn, strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, strings.ToLower(shout))
	time.Sleep(delay)
}
func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 2*time.Second)
	}
	conn.Close()
}
