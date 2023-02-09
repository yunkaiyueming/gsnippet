package main

import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "Go 语言编程之旅"
	}()

	<-ch
}

//go的两大杀器 pprof + trace 组合，此乃排查好搭档，谁用谁清楚，即使他并不万能。

// go run main.go 2> trace.out
//go tool trace trace.out