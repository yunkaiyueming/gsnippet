package main

import (
	"fmt"
	"os"
	_ "strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	hsot_name, _ := os.Hostname()
	fmt.Println(hsot_name)

	fmt.Println(os.Getwd())

	for _, e := range os.Environ() {
		//pair := strings.Split(e, "=")
		//fmt.Println(pair[0], pair[1])
		fmt.Println(e)
	}
}
