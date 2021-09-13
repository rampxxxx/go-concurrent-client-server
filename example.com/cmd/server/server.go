package main

import (
	"fmt"
	"log"
	"net"
)

func rcv_message(conn net.Conn, syn chan bool) {
	//vertex := datos.Vertex{}
	var netword_buffer = make([]byte, 1024)
	//dec := gob.NewDecoder(&netword_buffer)
	defer conn.Close()
	count, err := conn.Read(netword_buffer)
	if err == nil {
		fmt.Printf("Message count (%v) received:(%v)\n", count, netword_buffer)
	} else {
		fmt.Printf("Error receiving (%v)\n", err.Error())
	}
	//err = dec.Decode(&netword_buffer)
	//message, err := bufio.NewReader(conn).ReadString('\n')
	if err == nil {
		fmt.Printf("Message decoded:(%v)\n", netword_buffer)
		fmt.Fprintf(conn, "Desde el servidor (%v)\n", netword_buffer)
	} else {
		fmt.Printf("Error %v waiting client\n", err.Error())
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
