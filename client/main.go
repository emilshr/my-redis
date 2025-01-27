package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:6380")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close()

	_, err = conn.Write([]byte("Hello from client"))
	_, err = conn.Write([]byte("Another hello from client"))

	if err != nil {
		log.Fatalf("Error while writing data %s", err.Error())
	}
}
