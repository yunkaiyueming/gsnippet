package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	ListenKey()
}

func ldemo() {
	mleft := robotgo.AddEvent("mleft") //(鼠标参数: mleft, mright, wheelDown, wheelUp, wheelLeft, wheelRight)
	if mleft == 0 {
		fmt.Println("left mouse click")
	}

	k := robotgo.AddEvent("k") //(鼠标参数: mleft, mright, wheelDown, wheelUp, wheelLeft, wheelRight)
	if k == 0 {
		fmt.Println("k ")
	}

	mup := robotgo.AddEvent("wheelUp") //(鼠标参数: mleft, mright, wheelDown, wheelUp, wheelLeft, wheelRight)fdjjlk
	if mup == 0 {
		fmt.Println("mouse up")
	}
}

func ListenKey() {
	keys := "qwertyuiopasdfghjklzxcvbnm"
	for _, v := range keys {
		listen(v)
	}
}

func listen(v rune) {
	fmt.Println(string(v))
	if isOk := robotgo.AddEvent(string(v)); isOk == 0 {
		fmt.Println("press key:" + string(v))
		return
	}
}
