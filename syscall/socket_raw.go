package main

import (
	"bytes"
	"encoding/binary"
	. "fmt"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

type TCPHeader struct {
	SrcPort   uint16
	DstPort   uint16
	SeqNum    uint32
	AckNum    uint32
	Offset    uint8
	Flag      uint8
	Window    uint16
	Checksum  uint16
	UrgentPtr uint16
}

type PsdHeader struct {
	SrcAddr   uint32
	DstAddr   uint32
	Zero      uint8
	ProtoType uint8
	TcpLength uint16
}

func inet_addr(ipaddr string) uint32 {
	var (
		segments []string = strings.Split(ipaddr, ".")
		ip       [4]uint64
		ret      uint64
	)
	for i := 0; i < 4; i++ {
		ip[i], _ = strconv.ParseUint(segments[i], 10, 64)
	}
	ret = ip[3]<<24 + ip[2]<<16 + ip[1]<<8 + ip[0]
	return uint32(ret)
}

func htons(port uint16) uint16 {
	var (
		high uint16 = port >> 8
		ret  uint16 = port<<8 + high
	)
	return ret
}

func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}

func main() {
	var (
		msg       string
		psdheader PsdHeader
		tcpheader TCPHeader
	)

	Printf("Input the content: ")
	Scanf("%s", &msg)

	/*填充TCP伪首部*/
	psdheader.SrcAddr = inet_addr("127.0.0.1")
	psdheader.DstAddr = inet_addr("127.0.0.1")
	psdheader.Zero = 0
	psdheader.ProtoType = syscall.IPPROTO_TCP
	psdheader.TcpLength = uint16(unsafe.Sizeof(TCPHeader{})) + uint16(len(msg))

	/*填充TCP首部*/
	tcpheader.SrcPort = htons(3000)
	tcpheader.DstPort = htons(8080)
	tcpheader.SeqNum = 0
	tcpheader.AckNum = 0
	tcpheader.Offset = uint8(uint16(unsafe.Sizeof(TCPHeader{}))/4) << 4
	tcpheader.Flag = 2 //SYN
	tcpheader.Window = 60000
	tcpheader.Checksum = 0

	/*buffer用来写入两种首部来求得校验和*/
	var (
		buffer bytes.Buffer
	)
	binary.Write(&buffer, binary.BigEndian, psdheader)
	binary.Write(&buffer, binary.BigEndian, tcpheader)
	tcpheader.Checksum = CheckSum(buffer.Bytes())

	/*接下来清空buffer，填充实际要发送的部分*/
	buffer.Reset()
	binary.Write(&buffer, binary.BigEndian, tcpheader)
	binary.Write(&buffer, binary.BigEndian, msg)

	/*下面的操作都是raw socket操作，大家都看得懂*/
	var (
		sockfd int
		addr   syscall.SockaddrInet4
		err    error
	)
	if sockfd, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP); err != nil {
		Println("Socket() error: ", err.Error())
		return
	}
	defer syscall.Shutdown(sockfd, syscall.SHUT_RDWR)
	addr.Addr[0], addr.Addr[1], addr.Addr[2], addr.Addr[3] = 127, 0, 0, 1
	addr.Port = 8080
	if err = syscall.Sendto(sockfd, buffer.Bytes(), 0, &addr); err != nil {
		Println("Sendto() error: ", err.Error())
		return
	}
	Println("Send success!")
}

