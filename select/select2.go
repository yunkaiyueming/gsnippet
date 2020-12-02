package main

import(
	"fmt"
	"time"
)

//1.对一个关闭的通道再发送值就会导致panic。
//2.对一个关闭的通道进行接收会一直获取值直到通道为空。
//3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//4.关闭一个已经关闭的通道会导致panic。

func main(){
	doChan:=make(chan int)

	go func(){
		time.Sleep(time.Second)
		close(doChan)
	}()

	data,ok:=<-doChan
	fmt.Println(data,ok)
	fmt.Println("over...")
}