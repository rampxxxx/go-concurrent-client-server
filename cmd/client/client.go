package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, error := net.Dial("tcp", "127.0.0.1:8000")
	if conn != nil {
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Introduce texto a enviar:")
			text, _ := reader.ReadString('\n')
			// SEND TO SERVER ;-)
			fmt.Fprintf(conn, text+"\n")
			// Wait reply
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Recibido del servidor:" + message)
		}
	} else {
		fmt.Print("Error al connectar :" + error.Error())
	}
}
