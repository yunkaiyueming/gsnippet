package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.ShowAlert("ceshi", "hello world")
	fmt.Println(robotgo.GetPID())
}
