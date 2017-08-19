package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("err dialing:", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimedInput := strings.Trim(input, "\r\n")
		if trimedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimedInput))
		if err != nil {
			fmt.Println("err conn.write:", err)
			return
		}
	}
}
