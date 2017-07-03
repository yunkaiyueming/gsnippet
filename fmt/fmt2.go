package main

import (
	"fmt"
	//"log"
	//"regexp"
	//"strings"
)

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	EndColor = "\033[0m"
)

func ColorfulHTML(str string) string {
	return Color(str, Green)
}

func Color(str string, color uint8) string {
	return fmt.Sprintf("%s%s%s", ColorStart(color), str, EndColor)
}

func ColorStart(color uint8) string {
	return fmt.Sprintf("\033[%dm", color)
}

func main() {
	fmt.Println(Color("我们", Red))
	fmt.Println(Color("zheshi", Gray))
	fmt.Println(Color("你的", Green))
	fmt.Println(Color("呵呵", Blue))
	fmt.Println(Color("我们", White))
}
