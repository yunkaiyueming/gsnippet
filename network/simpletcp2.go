package main

import (
	"net"
	"fmt"
	"encoding/json"
	"bytes"
	"os"
	"encoding/binary"
)

func main() {
	service := "192.168.8.81:15001"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	res := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}`+"\r\n")
	_, err = conn.Write(res)
	checkError(err)
	//先假设所有请求返回的长度不超过1024
	initresult := make([]byte,1024*2*2)
	n,_ := conn.Read(initresult)
	fmt.Println("hehe")

	fmt.Println(n)
	fmt.Println(initresult)
	fmt.Println(string(initresult))
	header := initresult[1:4]
	buf := bytes.NewBuffer(header)
	var x int16
	binary.Read(buf, binary.LittleEndian, &x)
	fmt.Println(x)
	fmt.Println("hahhaa")
	total := int16(n)
	fmt.Println(total)

	for {
		if x > total {
			//newresult := make([]byte,8024)
			m,_ := conn.Read(initresult)
			fmt.Println(m)
			//initresult = BytesCombine(initresult,newresult)
			fmt.Println(string(initresult))
			total += int16(m)
			fmt.Println(total)
		}
		if total >= x {
			break
		}
	}

	//header不要，对于的字符也不要
	res = initresult[5:n]
	// fmt.Println(res)
	var r interface{}
	json.Unmarshal(res, &r)
	checkError(err)
	game, _:= r.(map[string]interface{})
	// fmt.Println(game)
	fmt.Println(game["cmd"])
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}




