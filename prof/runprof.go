package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"sync"
)

func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 1000000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c) //频繁申请内存
	}
}

func combine() {
	str := "abcdefghijklmopqrstuvws"
	for i := 0; i < 10000; i++ {
		str += "abcdefghijklmopqrstuvws"
	}
}

func workOnce(wg *sync.WaitGroup) {
	counter()
	combine()
	wg.Done()
}

func main() {
	var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memProfile = flag.String("memprofile", "", "write mem profile to file")
	var traceProfile = flag.String("traceprofile", "", "write trace profile to file")

	flag.Parse()
	//采样cpu运行状态
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	//trace采样
	if *traceProfile != "" {
		f, err := os.Create(*traceProfile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)
		defer f.Close()
		defer trace.Stop()
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}

	wg.Wait()
	//采样memory状态
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
	fmt.Println("finish")
}

//生成：
//go run runprof.go --cpuprofile=cpu.prof
//go run runprof.go --memprofile=mem.prof

//分析：
//go tool pprof cup.ppfor ===> top/web/list
