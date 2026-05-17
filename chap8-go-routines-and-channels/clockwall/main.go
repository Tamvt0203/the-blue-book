package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func parseMapping(args []string) (map[string]string, error) {
	var err error
	result := make(map[string]string)
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			err = fmt.Errorf("invalid format: %s", arg)
		}
		key := parts[0]
		value := parts[1]
		result[key] = value
	}
	return result, err
}
func main() {
	mapped_args, err := parseMapping(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var done bool
	for {
		if !done {
			for tz, addr := range mapped_args {
				conn, err := net.Dial("tcp", addr)
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println("Successfully connected")
					go handleConn(conn, tz)
				}
			}
			done = true
		}

	}

}
func handleConn(conn net.Conn, tz string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("can not read from connection with error: %s", err)
			return
		}
		fmt.Printf("%s now is: %s", tz, message)
	}
}
