package main

import (
	"io"
	"net"
)

func main() {
	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", "127.0.0.1:6666")
	if err != nil {
		println(err)
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	for {
		n, addr, err := pc.ReadFrom(buffer)
		switch err {
		case nil:
			println("C:", string(buffer[0:n]))
			pc.WriteTo([]byte("Server收到"), addr)
		case io.EOF:
			pc.Close()
		}
	}
}
