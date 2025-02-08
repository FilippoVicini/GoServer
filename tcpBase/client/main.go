package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() 

	fmt.Fprintf(conn, "GET /index.html HTTP/1.1\r\nHost: localhost\r\n\r\n")

	bs := make([]byte, 1024)

	n, err := conn.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs[:n]))
}
