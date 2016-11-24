package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	_ "time"
)

const RECV_BUF_LEN = 1024

var group sync.WaitGroup

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < 5; i++ {
		go SendMsg(conn, i)
		go GetMsg(conn, i)
		group.Add(1)
	}

	group.Wait()
	conn.Close()
	stop := make(chan int)
	<-stop
}

func SendMsg(conn net.Conn, i int) {
	//准备要发送的字符串
	msg := fmt.Sprintf("%d Hello World \n", i)
	_, err := conn.Write([]byte(msg))
	if err != nil {
		println("Write Buffer Error:", err.Error())
		return
	}
	fmt.Println("C:", msg)
}

func GetMsg(conn net.Conn, i int) {
	//从服务器端收字符串
	buf := make([]byte, RECV_BUF_LEN)
	n, err := conn.Read(buf)
	switch err {
	case nil:
		println(i, " S:", string(buf[:n]))
		return
	case io.EOF:
		println(i, " S: close", string(buf[:n]), err)
		return
	default:
		println(i, " S:", err.Error())
		return
	}

	group.Done()
}
