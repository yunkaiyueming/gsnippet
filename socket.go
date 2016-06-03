package main

import (
	"fmt"
	"io"
	"net"
)

/**
 * Starting the server
 * Accept the connection:  127.0.0.1:14071
 * Warning: End of data EOF
 */
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		panic("error listen: " + err.Error())
	}
	fmt.Println("Starting the server")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error accept: " + err.Error())
		}
		fmt.Println("Accept the connection: ", conn.RemoteAddr())
		go echoServer(conn)
	}
}

func echoServer(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Printf("Warning: End of data %s\n", err)
			return
		default:
			fmt.Printf("Error: read data %s\n", err)
			return
		}
	}
}
