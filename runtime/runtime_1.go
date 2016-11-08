package main

import (
	"fmt"
	"runtime"
	"time"
)

func test(c chan bool, n int) {
	x := 0
	for i := 0; i < 1000000000; i++ { //10亿
		x += i
	}

	fmt.Println(n, x)
	c <- true
}

//测试runtime使用gomaxprocs的运行速度
func main() {
	//runtime.GOMAXPROCS(1) //设置cpu的核的数量，从而实现并行
	runtime.GOMAXPROCS(runtime.NumCPU()) //默认使用的是最大cpu数，并行运行
	c := make(chan bool)
	t1 := time.Now().Unix()

	for i := 0; i <= 9; i++ {
		go test(c, i)
	}

	var cnt int = 0
	for {
		<-c
		cnt++
		if cnt == 9 {
			break
		}
	}
	t2 := time.Now().Unix()
	fmt.Println("use ", (t2 - t1), " s")
}
