package main

import "time"
import "fmt"

func main() {
	Cron1()
}

func Cron1() {
	var t *time.Timer

	f := func() {
		fmt.Printf("Expiration time : %v.\n", time.Now())
		fmt.Printf("C`s len: %d\n", len(t.C))
	}

	t = time.AfterFunc(1*time.Second, f)
	//让当前Goroutine 睡眠2s，确保大于内容的完整,不然main会走完
	//这样做原因是，time.AfterFunc的调用不会被阻塞。它会以一部的方式在到期事件来临执行我们自定义函数f。
	time.Sleep(2 * time.Second)
}

func Cron3() {
	//Millisecond 毫秒
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
