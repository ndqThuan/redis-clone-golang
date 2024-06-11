package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		// Read user input and put it inside buffer
		_, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// Out of input -> break
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}
}
