package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	message = flag.String("message", "Dies ist ein Test", "echo-Nachricht")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", "localhost:56789")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	const maxBufferSize = 1024
	l := len(*message)
	if l > maxBufferSize {
		l = maxBufferSize
	}
	n, err := conn.Write([]byte(*message)[:l])
	if err != nil {
		log.Fatalf("failed to write to server %w", err)
	}
	recv := make([]byte, n)
	_, err = conn.Read(recv)
	if err != nil {
		log.Printf("error receiving form server %w", err)
	}
	fmt.Printf("Received: %s\n", string(recv))
}
