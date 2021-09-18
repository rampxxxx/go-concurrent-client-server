package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"example.com/socket/cmd/pkg/datos"
)

func rcv_message(conn net.Conn, syn chan bool) {
	vertex := datos.Vertex{}
	//var netword_buffer = make([]byte, 100)
	//var netword_buffer bytes.Buffer
	//netword_buffer := new(bytes.Buffer) // 'new' return a pointer
	netword_buffer := bytes.NewBuffer(make([]byte, 100))
	netword_buffer.Grow(100)
	//var read_buffer = make([]byte, 100)
	var read_buffer = netword_buffer.Bytes()
	defer conn.Close()
	count, err := conn.Read(read_buffer)
	if err == nil {
		fmt.Printf("Message count (%v) received:(%v)\n", count, read_buffer)
	} else {
		fmt.Printf("Error receiving (%v)\n", err.Error())
	}
	//dec := gob.NewDecoder(&netword_buffer)
	dec := gob.NewDecoder(netword_buffer)
	err = dec.Decode(&vertex)
	//message, err := bufio.NewReader(conn).ReadString('\n')
	if err == nil {
		fmt.Printf("Message decoded:(%v)\n", vertex)
		fmt.Fprintf(conn, "Desde el servidor (%v)\n", vertex)
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
