package main

import (
	"fmt"
	"net"
)

var clients []*net.Conn

func handleConnection(c *net.Conn) {
	buffer := make([]byte, 1024)
	defer (*c).Close()
	_, err := (*c).Read(buffer)
	if err != nil {
		fmt.Println("An error occured", err)
	}
	(*c).Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello\r\n"))
	for index, client := range clients {
		if client == c {
			clients[index] = clients[len(clients)-1]
			clients = clients[:len(clients)-1]
		}
	}
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
		clients = append(clients, &c)
		go handleConnection(&c)
		fmt.Println(len(clients))
	}
}
