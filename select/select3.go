package main

import (
	"fmt"
	"time"
)

func sendwork(work chan int){
	time.Sleep(2*time.Second)
	for i := 1; i <= 5; i++ {
		work<-i
		// time.Sleep(2*time.Second)
	}
	close(work)
}

func main(){

	// var work chan int
	work:=make(chan int,100)
	go sendwork(work)

	count:=0
	for{
		select{
		case <-time.After(time.Second):
			fmt.Println("time second over")
			// return
		case value,ok:=<-work:
			fmt.Println("get work value",value,ok)
			break  //break不退出for循环
			fmt.Println("break after")
		// default:
		// 	fmt.Println("default run....")
		}
		count++
		if count>20 {break}
	}

	// for x:=range work {
	// 	fmt.Println("get work value", x)
	// }
	
}

//go version 1.19.4
//无缓存的channel无值，取不到； 有值，取值； 关闭后取到对应类型默认零值
//有缓存的channel无值，取不到； 有值，取值； 关闭后取到对应类型默认零值