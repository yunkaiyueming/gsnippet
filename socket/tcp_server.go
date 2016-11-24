package main

import (
	"fmt"
	"io"
	"net"
)

/*
golang提供的socket options接口也是基于上述模型的必要的属性设置。包括

    SetKeepAlive
    SetKeepAlivePeriod
    SetLinger
    SetNoDelay （默认no delay）
    SetWriteBuffer
    SetReadBuffer
golang提供的socket options接口也是基于上述模型的必要的属性设置。包括

    SetKeepAlive
    SetKeepAlivePeriod
    SetLinger
    SetNoDelay （默认no delay）
    SetWriteBuffer
    SetReadBuffer

*/
const RECV_BUF_LEN = 1024

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888") //侦听在6666端口
	if err != nil {
		panic("error listening:" + err.Error())
	}
	fmt.Println("Starting the server")

	for {
		conn, err := listener.Accept() //接受连接

		//设置tcp选项
		tcpConn, ok := conn.(*net.TCPConn)
		if !ok {
			//error handle
		}
		tcpConn.SetNoDelay(true)

		if err != nil {
			panic("Error accept:" + err.Error())
		}
		fmt.Println("Accepted the Connection :", conn.RemoteAddr())
		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			println("C: ", string(buf[0:n]))
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Println("C: ", buf[0:n], err)
			return
		default:
			fmt.Printf("C: %s \n", err)
			return
		}
	}
}
