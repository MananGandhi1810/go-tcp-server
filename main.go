package main

import (
	"fmt"
	"net"
)

func main() {
	port := ":1234"
	fmt.Println("Server Starting on port", port)
	_, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println("Error", err)
	}
	
}
