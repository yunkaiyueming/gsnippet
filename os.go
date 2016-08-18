package main

import (
	"fmt"
	"os/exec"
)

func main() {
	TestCmd()
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
