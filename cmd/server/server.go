package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func rcv_message(conn net.Conn) {
	message, error := bufio.NewReader(conn).ReadString('\n')
	if error == nil {
		fmt.Print("Message received:", string(message))
	} else {
		fmt.Printf("Error %v waiting client:", error.Error())
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Start server...")

	ln, _ := net.Listen("tcp", "127.0.0.1:8000")
	for {
		conn, _ := ln.Accept()
		go rcv_message(conn)
	}
}
