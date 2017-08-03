package main

//粘包问题， 自动读到\r\n就断的问题？？
import (
	"bytes"
	"encoding/binary"
	//"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {
	service := "192.168.8.81:15001"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	res := []byte(`1 {"zoneid":1,"secret":"0d734a1dc94fe5a914185f45197ea846","cmd":"admin.getuserinfo","params":[],"uid":1000001}` + "\r\n")
	_, err = conn.Write(res) //2554
	//_, err = conn.Write(res)
	checkError(err)
	//先假设所有请求返回的长度不超过1024
	initresult := make([]byte, 1024)
	n, _ := conn.Read(initresult)

	//fmt.Println(initresult)
	fmt.Println("header:", string(initresult))
	header := initresult[1:4]
	buf := bytes.NewBuffer(header)
	var x int16
	binary.Read(buf, binary.LittleEndian, &x)
	fmt.Println("header length", x)
	total := int16(n)

	allMsg := string(initresult)
	for {
		if x > total {
			nextLength := x - total
			if nextLength > 0 {
				nextReadRet := make([]byte, nextLength)
				m, _ := conn.Read(nextReadRet)
				fmt.Println(m)
				//initresult = BytesCombine(initresult,newresult)
				//fmt.Println(string(nextReadRet))

				allMsg += string(nextReadRet)

				total += int16(m)
				//fmt.Println(total)
			}
		} else {
			break
		}
	}

	//header不要，对于的字符也不要
	fmt.Println("==========================")
	fmt.Println(allMsg)
	//fmt.Println(allMsg[5:len(allMsg)])
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
