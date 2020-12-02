package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)


func DialCustom(network, address string, address2 string, timeout time.Duration, localIP []byte, localPort int, )(net.Conn,error) {
	netAddr := &net.TCPAddr{Port:localPort}

	if len(localIP) != 0 {
		netAddr.IP = localIP
	}

	fmt.Println("netAddr:", netAddr)

	d := net.Dialer{Timeout: timeout, LocalAddr: netAddr}

	a,err:=d.Dial(network, address)
	if err != nil {
		fmt.Println("dial failed:", err)
	}

	_,err2:=d.Dial(network, address2)
	if err2 != nil {
		fmt.Println("dial failed2:", err2)
	}
	return a,err
}


func main() {

	serverAddr := "192.168.8.83:15001"
	serverAddr2 := "192.168.8.83:15002"
	// 172.28.0.180
	//localIP := []byte{0xAC, 0x1C, 0, 0xB4}  // 指定IP
	localIP := []byte{} //  any IP，不指定IP
	localPort := 9001   // 指定端口
	conn, err := DialCustom("tcp", serverAddr, serverAddr2, time.Second*10, localIP,localPort)
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// serverAddr2 := "192.168.8.83:15002"
	// conn2, err2 := DialCustom("tcp", serverAddr2, time.Second*10, localIP,localPort)
	// if err2 != nil {
	// 	fmt.Println("dial failed:", err2)
	// 	os.Exit(1)
	// }
	// defer conn2.Close()

	buffer := make([]byte, 512)
	reader := bufio.NewReader(conn)

	n, err2 := reader.Read(buffer)
	if err2 != nil {
		fmt.Println("Read failed:", err2)
		return
	}

	fmt.Println("count:", n, "msg:", string(buffer))

	select{}	
}