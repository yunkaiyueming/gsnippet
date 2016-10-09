package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	p(time.Now().Unix()) //返回当前的unix时间戳
}
