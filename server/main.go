package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6380")
	if err != nil {
		panic(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error encountered", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection received from ", conn.RemoteAddr().String())

	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error while reading buffer %s\n", err.Error())
		return
	}
	fmt.Println("Received some data ", string(buffer[:]))
}
