package main

import (
	"fmt"
	"io"
	"net"
	_ "time"
)

const RECV_BUF_LEN = 1024

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:6666")
	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, RECV_BUF_LEN)
	for i := 0; i < 5; i++ {
		//准备要发送的字符串
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}
		//fmt.Println("C:", msg)

		//从服务器端收字符串
		n, err = conn.Read(buf)
		switch err {
		case nil:
			println("S:", string(buf[0:n]))
		case io.EOF:
			println("S: close", string(buf[0:n]), err)
		default:
			println("S:", err.Error())
			break
		}

		//等一秒钟
		//time.Sleep(time.Second)
	}

	conn.Close()
	//stop := make(chan int)
	//<-stop
}
