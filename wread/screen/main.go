package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	color := robotgo.GetPixelColor(100, 200)
	fmt.Println(color)
	fmt.Println(robotgo.GetPixelColor(400, 10))
	fmt.Println(robotgo.GetScreenSize())

	saveScreen()
}

func saveScreen() {
	bitMap := robotgo.CaptureScreen(10, 10, 500, 500)
	robotgo.SaveBitmap(bitMap, "./out.png")
}
