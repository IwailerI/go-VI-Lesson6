package main

import (
	"fmt"
	"net"
)

// Data ...
type Data struct {
	N  string
	Ch chan int //mutex
}

var dat Data

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:12667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	dat.Ch = make(chan int, 1) // Important for blocking
	dat.Ch <- 1                // at end of execution of goroutine

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleConn(conn)
	}
}

// HandleConn guess what
func HandleConn(conn net.Conn) {
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		conn.Close()
		return
	}
	<-dat.Ch // wait until can acces
	// dat.Mutex.Lock() // (should be) the same thing

	fmt.Println(string(buf[:n]))
	dat.N = string(buf[:n])

	data := []byte("Connection is great!")

	dat.Ch <- 1
	// dat.Mutex.Unlock() // (should be) the same thing

	conn.Write(data)

	conn.Close()
}
