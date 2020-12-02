package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := "192.168.8.83:15004"
	conn, err := net.Dial("tcp", service)
	checkError(err)

	for{

	//res := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"fixbug.test","params":[],"uid":1000001}` + "\r\n")
	res := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"user.sync","params":[],"uid":1005083}` + "\r\n")
	_, err = conn.Write(res)
	checkError(err)
	//先取消息的header,header长度我们事先知道，所以这个长度大于header的长度即可
	const RECV_BUF_LEN = 5
	initheader := make([]byte, RECV_BUF_LEN)
	n, _ := conn.Read(initheader)
	header := initheader[1:4]

	//获取请求的实际长度 len
	buf := bytes.NewBuffer(header)
	var lenth int16
	binary.Read(buf, binary.LittleEndian, &lenth)
	fmt.Println(lenth)

	total := int16(n)
	//fmt.Println(total)

	//allMsg := string(initheader[RECV_BUF_LEN:n])
	var allMsg string
	//fmt.Println(allMsg)
	for {
		if lenth > total {
			nextLength := lenth - total
			if nextLength > int16(1024) {
				//fmt.Println("haha")
				nextLength = int16(1024)
			}
			if nextLength > 0 {
				nextReadRet := make([]byte, nextLength)
				m, _ := conn.Read(nextReadRet)
				fmt.Println(m, "===>", string(nextReadRet))
				allMsg += string(nextReadRet[0:m])
				total += int16(m)
			}
		} else {
			break
		}
	}

	//fmt.Println("==========================")
	//fmt.Println(allMsg)
	HandleMsg(allMsg)
	//var r interface{}
	//json.Unmarshal([]byte(allMsg), &r)
	//checkError(err)
	//game, _:= r.(map[string]interface{})
	//fmt.Println(game["cmd"])

	time.Sleep(time.Second)
	}

}

func HandleMsg(msg string) {
	t := make(map[string]interface{})
	fmt.Println("========解析start===========")
	err := json.Unmarshal([]byte(msg), &t)
	if err != nil {
		fmt.Println("map err", err.Error())
	}

	fmt.Println(t)
	fmt.Println("=========解析end==========")
	//fmt.Println(t["cmd"])
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
