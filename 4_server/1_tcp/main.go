package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {

			log.Fatal(err)
			return
		}

		// handle multiple connection
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// create new reader from connection
	reader := bufio.NewReader(conn)
	// read from clinet command line
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(conn, "Error reading command: %v\n", err)
		return
	}

	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintf(conn, "Invalid command\n")
		return
	}

	command := parts[0]
	resource := parts[1]
	log.Printf("Received command: %s %s\n", command, resource)

	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintf(conn, "Unknown command: %s\n", command)
	}

}

func handleGet(conn net.Conn, resource string) {
	fmt.Fprintf(conn, "GET command received from %s\n", resource)
}
