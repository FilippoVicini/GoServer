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
	}
	defer ln.Close() 
	log.Println("Server is listening on port 8080...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) 	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintln(conn, "Invalid request format")
		return
	}

	command := parts[0]
	resource := parts[1]
	log.Printf("Received command: %s %s", command, resource)

	switch command {
	case "GET":
		handleGet(conn, resource)
	default:
		fmt.Fprintln(conn, "Unknown command")
	}
}

func handleGet(conn net.Conn, resource string) {
	fmt.Fprintf(conn, "Connection received for resource: %s\n", resource)
}
