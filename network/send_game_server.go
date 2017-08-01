package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	// "encoding/binary"
)

func main() {

	//	if len(os.Args) != 2 {
	//		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	//		os.Exit(1)
	//	}

	// service := os.Args[1]
	// conn,err := net.Dial("tcp", service)
	// checkError(err)
	// defer conn.Close()
	// fmt.Println("connect success")
	// _, err = conn.Write([]byte("1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}"))
	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	// data := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}\r\n`)

	conn, err := net.Dial("tcp", "192.168.7.0:1008")
	checkError(err)
	// _, err = conn.Write([]byte("1 "))
	//res1 := []byte("1 ")
	// s1 := []byte("1 ")
	s2 := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}\r\n\r\n`)
	// s3 := []byte("\r\n\r\n")

	// s := [][]byte{s1,s2,s3}
	// sep1 := []byte("")
	// res := bytes.Join(s,sep1)//你好,世界

	// res := "1 "+'{"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"fixbug.test","params":[],"uid":1000001}'+"\r\n\r\n"

	//res2 : = []byte({"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"fixbug.test","params":[],"uid":1000001}\r\n\r\n`))
	// res := [][]byte{[]byte("1 "),[]byte(`{"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"fixbug.test","params":[],"uid":1000001}\r\n\r\n`)}
	_, err = conn.Write(s2)
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
