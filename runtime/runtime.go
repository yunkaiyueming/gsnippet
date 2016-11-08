package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go Env()
	go Run()

	Mem()

	time.Sleep(time.Second * 2)
}

func Env() {
	defer fmt.Println("run env defer")
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Version())

	runtime.Goexit()

	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	fmt.Println("----------------------")
}

func Run() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	runtime.Gosched()
	fmt.Println("start run ")
	runtime.GC()
	fmt.Println("----------------------")
}

func Mem() {
	m := new(runtime.MemStats)
	p := fmt.Println
	p(m.Alloc)
	p(m.HeapAlloc)
	fmt.Println("----------------------")
}
