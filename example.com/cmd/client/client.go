package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"

	"example.com/socket/cmd/pkg/datos"
)

func snd_message(msg string /*sync chan int,*/, i int, done chan bool) {

	fmt.Printf("Connectando %v\n", i)
	conn, error := net.Dial("tcp", "127.0.0.1:8000")
	if error == nil {
		fmt.Printf("Exito Connectando %v!\n", i)
		defer conn.Close()
		//reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Introduce texto a enviar:")
		//text, _ := reader.ReadString('\n')
		// SEND TO SERVER ;-)
		vertex := datos.Vertex{X: 1, Y: 2}
		fmt.Fprintf(conn, msg+"\n")
		fmt.Fprintf(conn, strconv.FormatInt(int64(vertex.X), 10))
		// Wait reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Recibido del servidor (%v): %v\n", i, message)
		//i := <-sync
		fmt.Printf("Gorutine %v ends\n", strconv.FormatInt(int64(i), 10))
	} else {
		fmt.Printf("Error al connectar (%v)", error.Error())
	}
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Print("Starting client\n")
	//synchro := make(chan int, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("Calling gorutine %v \n", strconv.FormatInt(int64(i), 10))
		//synchro <- i
		go snd_message("Hola"+strconv.FormatInt(int64(i), 10) /*synchro,*/, i, done)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
