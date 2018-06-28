package main

import (
	//	"fmt"

	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
	err = beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		panic(err)
	}

	//	err = beeep.Alert("Title", "Message body", "assets/warning.png")
	//	if err != nil {
	//		panic(err)
	//	}
}
