package main

import (
	"encoding/json"
	//"strings"
	"fmt"
	"io"
	"net"
	"time"
)

const RECV_BUF_LEN = 1024 * 4

var msgCh = make(chan string)

func main() {
	conn, err := net.Dial("tcp", "192.168.8.81:15001")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("conn:", conn.RemoteAddr())

	go SendMsg(conn)
	go GetMsg(conn)
	//go BeatHeart(conn) //定时心跳
	go HandleMsg()

	stop := make(chan int)
	<-stop
}

func SendMsg(conn net.Conn) {
	//准备要发送的字符串
	s := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}` + "\r\n")
	_, err := conn.Write(s)
	if err != nil {
		println("Write Buffer Error:", err.Error())
		return
	}
}

func BeatHeart(conn net.Conn) {
	tick := time.NewTicker(5 * time.Second)
	for {
		<-tick.C
		fmt.Println("send")
		SendMsg(conn)
	}
}

func GetMsg(conn net.Conn) {
	//从服务器端收字符串
	buf := make([]byte, RECV_BUF_LEN)
	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			fmt.Println(" S:", string(buf[:n]))
			msgCh <- string(buf[:n])

		case io.EOF:
			fmt.Println(" S: close", string(buf[:n]), err)

		default:
			fmt.Println(" S:", err.Error())
		}
	}
}

func HandleMsg() {
	t := make(map[string]interface{})
	for { //如果这里处理数据比较耗时，可以使用1个协程处理1个消息
		msg := <-msgCh
		//header
		//3
		msg = msg[5:len(msg)]
		fmt.Println("========解析start===========")
		json.Unmarshal([]byte(msg), &t)
		fmt.Println(t)

		fmt.Println("=========解析end==========")
	}
}
