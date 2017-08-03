package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str1 := ReadByIoutil("E:/GO_PATH/src/gsnippet/network/client.json")
	fmt.Println(len(str1), len([]byte(str1)))
	getSendByte(str1)

	str2 := "abcde"
	fmt.Println(len(str2), len([]byte(str2)))
	getSendByte(str2)

	str3 := "12我们ab"
	fmt.Println(len(str3), len([]byte(str3)))
	getSendByte(str3)
}

func getSendByte(msg string) []byte {
	headerLen := len([]byte(msg))

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int16(headerLen))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//fmt.Printf("% x", buf.String())
	//fmt.Println("% x", buf.Bytes())

	fmt.Printf("% x, %d", buf.Bytes(), buf.Len())
	fmt.Println("===================================")
	return []byte(fmt.Sprintf("%s%s", buf.String(), msg))
}

func ReadByIoutil(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}
