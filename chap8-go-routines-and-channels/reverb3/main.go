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
func handleCloseConenction(wg *sync.WaitGroup, conn net.Conn) {
	fmt.Fprintln(conn, "Connection closed")
	conn.(*net.TCPConn).CloseRead()
	fmt.Println("read closed")
	go func() {
		wg.Wait()
		conn.(*net.TCPConn).CloseWrite()
		fmt.Println("write closed")
	}()
}
func handleInput(conn net.Conn, reset chan struct{}, wg *sync.WaitGroup, done chan struct{}) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		reset <- struct{}{}
		wg.Add(1)
		go echo(conn, input.Text(), 1*time.Second, wg)
	}
	done <- struct{}{}
}
func handleConn(conn net.Conn) {
	reset := make(chan struct{})
	done := make(chan struct{})
	countdown := 10
	var wg sync.WaitGroup
	tick := time.Tick(1 * time.Second)
	go handleInput(conn, reset, &wg, done)
	for countdown > 0 {
		select {
		case <-reset:
			countdown = 10
		case <-tick:
			countdown = countdown - 1
		case <-done:
			fmt.Println("close by client")
			handleCloseConenction(&wg, conn)
			return
		}
	}
	fmt.Println("close by server")
	handleCloseConenction(&wg, conn)

}
