package main

import "os"
import "fmt"

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

func tips() {
	fmt.Println("ok")
	i := true
	if i {
		fmt.Println("aa")
	}
	fmt.Println("wrong")
}
