package main

//粘包问题， 自动读到\r\n就断的问题？？
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var CLIENT_RECV_LEN = 2
var CLIEN_RECV_DEFAULT_LEN int16 = 20
var recevMsgCh = make(chan string)

func main() {
	service := "127.0.0.1:8082"
	conn, err := net.Dial("tcp", service)
	checkError(err)
	res := ReadByIoutil("E:/GO_PATH/src/gsnippet/network/client.json")
	go handleServerMsg()

	_, err = conn.Write(getSendByte(res))
	checkError(err)

	headerByte := make([]byte, CLIENT_RECV_LEN)
	conn.Read(headerByte)

	buf := bytes.NewBuffer(headerByte)
	var headerLen int16
	binary.Read(buf, binary.LittleEndian, &headerLen)
	fmt.Println("header length:", headerLen)

	var total int16
	oneMsg := ""
	var nextLength int16
	for {
		if total < headerLen {
			if headerLen > CLIEN_RECV_DEFAULT_LEN {
				nextLength = CLIEN_RECV_DEFAULT_LEN
			} else {
				nextLength = headerLen
			}
			nextReadRet := make([]byte, nextLength)
			m, _ := conn.Read(nextReadRet)
			oneMsg += string(nextLength[0:m])
			total += int16(m)
			fmt.Println(total)
		} else {
			if oneMsg != "" {
				recevMsgCh <- oneMsg
				oneMsg = ""
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
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

func handleServerMsg() {
	for {
		data := <-recevMsgCh
		fmt.Println("S:", data)
	}
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
