package main

import (
	"fmt"
	"net"
)

func main() {
	// Establish connection
	conn, err := net.Dial("tcp", "127.0.0.1:15395")
	if err != nil {
		panic(err)
	}
	// defer conn.Close() // no need for this

	// Create json request
	conn.Write([]byte(""))

	// Recieve responce
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf[:n])
}
