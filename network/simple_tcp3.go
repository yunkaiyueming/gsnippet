package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

//tcp的readfull版本
func main() {
	service := "192.168.8.81:15001"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	res := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}` + "\r\n")
	_, err = conn.Write(res)
	checkError(err)
	//先假设所有请求返回的长度不超过1024
	headerRes := make([]byte, 5)
	n, _ := conn.Read(headerRes)
	fmt.Println(n, string(headerRes))

	buf := bytes.NewBuffer(headerRes[1:4])
	var headerLength int16
	binary.Read(buf, binary.LittleEndian, &headerLength)

	fmt.Println(headerLength)
	data := make([]byte, headerLength-5)
	io.ReadFull(conn, data)

	var paserData map[string]interface{}
	json.Unmarshal(data, &paserData)
	fmt.Println(paserData)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
