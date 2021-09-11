package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func rcv_message(conn net.Conn, syn chan bool) {
	defer conn.Close()
	message, error := bufio.NewReader(conn).ReadString('\n')
	if error == nil {
		fmt.Printf("Message received:(%v)\n", string(message))
		fmt.Fprintf(conn, "Desde el servidor\n"+string(message))
	} else {
		fmt.Printf("Error %v waiting client\n", error.Error())
		//os.Exit(1)
	}
	syn <- true
}

func main() {
	synchro := make(chan bool) // one at a time, qos? :-)
	fmt.Println("Start server...")

	fmt.Println("Listening ...")
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("Error listening ... (%s)", err)
	}
	defer listen.Close()
	for {
		fmt.Println("Locking channel ...")
		fmt.Println("Accepting connection ...")
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalf("Error accept ... (%s)", err)
			continue
		}
		fmt.Println("Handling connection ...")
		go rcv_message(conn, synchro)
		<-synchro
	}
}
