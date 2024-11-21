package main

import (
	"fmt"
	"net"
)

func handleConnection(c *net.Conn) {
	buffer := make([]byte, 1024)
	defer (*c).Close()
	_, err := (*c).Read(buffer)
	if err != nil {
		fmt.Println("An error occured", err)
	}
	data := string(buffer)
	fmt.Println(data)
	(*c).Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello\r\n"))
}

func main() {
	port := ":1234"
	fmt.Println("Server Starting on port", port)
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer listener.Close()
	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(&c)
	}
}
