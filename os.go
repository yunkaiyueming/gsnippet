package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	GetOs()
}

func TestCmd() {
	arg := []string{"Hello", "World!"}
	cmd := exec.Command("echo", arg...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	//test in ArchLinux
	fmt.Printf("The output is: %s\n", output) //The output is: Hello World!
}

func GetOs() {
	//runtime.GOARCH 返回当前的系统架构；runtime.GOOS 返回当前的操作系统。
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
