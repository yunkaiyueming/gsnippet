package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	MyParseIp()
	MyGetIp()
}

func MyParseIp() {
	name := "192.168.1.1"
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
}

func MyGetIp() {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}
