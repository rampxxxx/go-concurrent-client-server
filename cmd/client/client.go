package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func snd_message(msg string, sync chan int) {

	conn, error := net.Dial("tcp", "127.0.0.1:8000")
	if conn != nil {
		for {
			//reader := bufio.NewReader(os.Stdin)
			//fmt.Print("Introduce texto a enviar:")
			//text, _ := reader.ReadString('\n')
			// SEND TO SERVER ;-)
			fmt.Fprintf(conn, msg+"\n")
			// Wait reply
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Recibido del servidor:" + message)
			i := <-sync
			fmt.Print("Gorutine %v ends:" + strconv.FormatInt(int64(i), 10))
		}
	} else {
		fmt.Print("Error al connectar :" + error.Error())
	}
}

func main() {
	fmt.Print("Starting client\n")
	synchro := make(chan int, 10)
	for i := 0; i < 10; i++ {
		fmt.Printf("Client %v \n", strconv.FormatInt(int64(i), 10))
		synchro <- i
		go snd_message("Hola"+strconv.FormatInt(int64(i), 10), synchro)
	}
}
