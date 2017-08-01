package main

import (
	//"bytes"
	"fmt"
	"io"
	"net"
	//"strings"
	"sync"
	_ "time"
)

const RECV_BUF_LEN = 1024 * 2

var group sync.WaitGroup

func main() {
	conn, err := net.Dial("tcp", "192.168.8.81:15001")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("conn:", conn.RemoteAddr())

	for i := 0; i < 1; i++ {
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

	//s1 := []byte("")
	s2 := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"fixbug.test","params":[],"uid":1000001}` + "\r\n")
	//s3 := []byte("\r\n")

	//s := [][]byte{s2, s3}
	//sep1 := []byte("")
	//res := bytes.Join(s, sep1) //你好,世界

	_, err := conn.Write(s2)
	if err != nil {
		println("Write Buffer Error:", err.Error())
		return
	}
	//fmt.Println("C:", msg)
}

func GetMsg(conn net.Conn, i int) {
	//从服务器端收字符串
	buf := make([]byte, RECV_BUF_LEN)
	n, err := conn.Read(buf)
	switch err {
	case nil:
		fmt.Println(i, " S:", string(buf[:n]))
		return
	case io.EOF:
		fmt.Println(i, " S: close", string(buf[:n]), err)
		return
	default:
		fmt.Println(i, " S:", err.Error())
		return
	}

	group.Done()
}
