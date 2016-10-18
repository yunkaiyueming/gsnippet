package main

import (
	"fmt"
	"runtime"
	_ "time"
)

var (
	flag = false
	str  string
)

func xiaorui() {
	flag = true
	str = "setup flag to true"
}

func main() {
	runtime.GOMAXPROCS(4)
	go xiaorui()
	//time.Sleep(1 * time.Second)
	// 理论来说，当我在xiaorui()把flag 改为true后，后面的逻辑会退出.
	for {
		if flag {
			break
		}
	}
	fmt.Println(str)
}
