package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
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
		vertex := datos.Vertex{X: 6, Y: 9}
		var netword_buffer bytes.Buffer
		enc := gob.NewEncoder(&netword_buffer)
		err := enc.Encode(vertex)
		fmt.Printf("Datos pre-encoded : (%v) \n", vertex)
		fmt.Printf("Datos encoded : (%v) \n", netword_buffer.Bytes())
		if err == nil {
			fmt.Printf("Exito codificando Len:(%v) Cap:(%v)\n", netword_buffer.Len(), netword_buffer.Cap())
		} else {
			fmt.Printf("Error al codificar (%v)", err.Error())
		}
		_, err = conn.Write(netword_buffer.Bytes())
		//_, err = conn.Write([]byte("Hola desde cliente"))
		if err == nil {
			fmt.Printf("Exito enviando ")
		} else {
			fmt.Printf("Error al enviar (%v)", err.Error())
		}
		//fmt.Fprintf(conn, msg+"\n")
		//fmt.Fprintf(conn, strconv.FormatInt(int64(vertex.X), 10))
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
	for i := 0; i < datos.MAX_THREADS; i++ {
		fmt.Printf("Calling gorutine %v \n", strconv.FormatInt(int64(i), 10))
		//synchro <- i
		go snd_message("Hola"+strconv.FormatInt(int64(i), 10) /*synchro,*/, i, done)
	}

	for i := 0; i < datos.MAX_THREADS; i++ {
		<-done
	}
}
