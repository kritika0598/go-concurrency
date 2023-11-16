package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080") // create a TCP Server
	if err != nil {
		log.Fatalf("Could not create a listener: %v", err)
	}

	var connections int32
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue // continue as we want to server other connections
		}
		connections++

		go func() {
			defer func() {
				_ = conn.Close()
				atomic.AddInt32(&connections, -1) // cannot accept multiple connections at the same tome
			}()
			// serve connections
			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			time.Sleep(time.Second)
			_, err := conn.Write([]byte("success"))
			if err != nil {
				log.Fatalf("Could not write to connection: %v", err)
			}
		}()
	}
}
