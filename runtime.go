package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("ok")

	fmt.Printf("%d\n", runtime.MemStats.Alloc)
}
