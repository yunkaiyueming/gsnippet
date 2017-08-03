package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

const SERVER_DEFAULT_CONN_RECV = 512
const SERVER_HEADER_RECV = 2

var connMap = make(map[string]net.Conn)
var recvMsgCh = make(chan string, 10)
var sendMsgCh = make(chan string, 10)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("conn err:", err.Error())
	} else {
		fmt.Println("server start:", ln.Addr().String())
	}

	go RouteRecvMsg()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("conn err:", err.Error())
			conn.Close()
			continue
		} else {
			fmt.Println("connected:", conn.RemoteAddr().String())

			clientId := fmt.Sprintf("%s", conn.RemoteAddr().String())
			connMap[clientId] = conn
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	var headerLength int16
	readLength := SERVER_HEADER_RECV
	oneMsg := ""
	isReadHeader := true
	var total int16

	for {
		readMsg := make([]byte, readLength)
		n, err := conn.Read(readMsg)
		if err != nil {
			fmt.Println("C: err", err.Error())
			continue
		} else if err == io.EOF {
			fmt.Println("C closed:")
			conn.Close()
		} else {
			fmt.Println("C:", string(readMsg[:n]))

			if readLength == SERVER_HEADER_RECV && isReadHeader {
				buf := bytes.NewBuffer(readMsg[:n])
				err := binary.Read(buf, binary.LittleEndian, &headerLength)
				if err != nil {
					fmt.Println("read heder err", err.Error())
				} else {
					fmt.Println("header length:", headerLength)
				}

				isReadHeader = false
			} else {
				oneMsg += string(readMsg[:n])
				total += int16(n)

				if total < headerLength { // not finish
					if headerLength-total < SERVER_DEFAULT_CONN_RECV {
						readLength = int(headerLength - total)
					} else {
						readLength = SERVER_DEFAULT_CONN_RECV
					}
					fmt.Println("=============readLength:", readLength, "=====")
				} else { //over
					fmt.Println("read over")
					recvMsgCh <- oneMsg
					readLength = SERVER_HEADER_RECV
					oneMsg = ""
				}
			}
		}
	}
}

func RouteRecvMsg() {
	mapData := make(map[string]interface{})
	for {
		msg := <-recvMsgCh
		err := json.Unmarshal([]byte(msg), &mapData)
		if err != nil {
			fmt.Println("parse msg err:", err.Error())
		} else {
			fmt.Println("parse json ok", mapData)
		}

		switch mapData["action"] {
		case "user":
			fmt.Println("route user")

		case "fight":
			fmt.Println("route fight")

		case "chat":
			fmt.Println("route chat")

		case "broadcast":
			fmt.Println("broad cast")
			for clientId, conn := range connMap {
				sendJson := `{"name":"aa","chat":"true","ret":"success", "msg":"hello everyone"}`
				conn.Write(getSendByte(sendJson))
				fmt.Println(clientId, "broadcast ok")
			}
		}
	}
}

func RouteSendMsg() {
	mapData := make(map[string]interface{})
	for {
		msg := <-sendMsgCh
		err := json.Unmarshal([]byte(msg), &mapData)
		if err != nil {
			fmt.Println("parse msg err:", err.Error())
		}

		switch mapData["action"] {
		case "user":
			fmt.Println("route user")

		case "fight":
			fmt.Println("route fight")

		case "chat":
			fmt.Println("route chat")
		}
	}
}

func getSendByte(msg string) []byte {
	headerLen := len([]byte(msg))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int16(headerLen))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.String())

	fmt.Println(fmt.Sprintf("%s%s", buf.String(), msg))
	return []byte(fmt.Sprintf("%s%s", buf.String(), msg))
}
