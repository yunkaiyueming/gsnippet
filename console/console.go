package main

import (
	"fmt"
	_ "os"
	"strings"
	"time"
)

func main() {
	for i := 0; i != 22; i = i + 2 {
		fmt.Printf("%d/20 => %s\n", i, strings.Repeat("==", i))
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("over")
}
