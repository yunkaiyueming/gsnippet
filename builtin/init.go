package main

import (
	"fmt"
)

func prepare() {
	fmt.Println(1e9)
	fmt.Println("prepare")
}

func init() {
	prepare()
	fmt.Println("init")
}

func main() {
	fmt.Println("main func start")
}
