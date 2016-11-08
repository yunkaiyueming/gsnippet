package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	TestCmd()
}

func TestCmd() {
	//cmd := exec.Command("netstat", "-ano")
	//cmd := exec.Command("tasklist")
	cmd := exec.Command("cmd", "-dir")
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

//根据操作系统获取不同命令解析器
func GetOsCmd(cmd string) string {
	var bash_path string
	if runtime.GOOS == "windows" {
		bash_path = "C:/Windows/System32/cmd.exe"
	} else {
		//all_cmd = "/bin/bash", "-c"
		bash_path = "/bin/bash"
	}

	fmt.Println(bash_path, cmd)
	return cmd
}
