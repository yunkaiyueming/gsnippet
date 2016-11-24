package main

import (
	"fmt"
	"net"
	"os"
	_ "time"
)

const (
	MAX_CONN_NUM = 5
)

//echo server Goroutine
func EchoFunc(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			//println("Error reading:", err.Error())
			return
		}
		//send reply
		println("C:", string(buf[:n])) //线程安全的读和写，所有对conn的Read和Write操作都是有fdMutex互斥的
		//多个Goroutine对同一conn的并发读不会出现读出内容重叠的情况，但内容断点是依 runtime调度来随机确定的。存在一个业务包数据，1/3内容被goroutine-1读走，另外2/3被另外一个goroutine-2读 走的情况。比如一个完整包：world，当goroutine的read slice size < 5时，存在可能：一个goroutine读到 “worl”,另外一个goroutine读出”d”。
		_, err = conn.Write(buf[:n])
		if err != nil {
			//println("Error send reply:", err.Error())
			return
		}
	}
}

//initial listener and run
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("running ...\n")

	var cur_conn_num int = 0
	conn_chan := make(chan net.Conn)
	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	/*
		go func() {
			for _ = range time.Tick(1e8) {
				fmt.Printf("cur conn num: %f\n", cur_conn_num)
			}
		}()
	*/

	//最多routines数
	for i := 0; i < MAX_CONN_NUM; i++ {
		go func(i int) {
			fmt.Println("i:", i)
			for conn := range conn_chan {
				println("i:", i, "conn get")
				ch_conn_change <- 1
				EchoFunc(conn)
				ch_conn_change <- -1
			}
		}(i)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accept:", err.Error())
			return
		}
		conn_chan <- conn
	}
}
