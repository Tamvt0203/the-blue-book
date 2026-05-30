package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
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
func echo(conn net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
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
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(conn, input.Text(), 1*time.Second, &wg)
	}
	conn.(*net.TCPConn).CloseRead()
	for {
		wg.Wait()
		conn.(*net.TCPConn).CloseWrite()
	}
}
