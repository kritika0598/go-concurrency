package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":1729") // create a TCP Server
	if err != nil {
		log.Fatalf("Could not create a listener: %v", err)
	}

	fmt.Println("Hello world", listener)

	for {
		conn, err := listener.Accept() // blocking call
		if err != nil {
			log.Fatalf("Could not accept the connection: %v", err)
		}
		fmt.Println("Connection: ", conn)
		go do(conn)
	}
}

func do(conn net.Conn) {
	var buf []byte = make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// some processing
	time.Sleep(5 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n Hello world!\r\n"))
	conn.Close()
}
