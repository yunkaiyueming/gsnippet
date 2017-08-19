package main

import (
	"fmt"
)

func main() {
	uInfos := make([]int, 1010)
	end, t, splitNum := 0, 0, 300
	for i := 0; i < len(uInfos); i = end {
		end = (t + 1) * splitNum
		if end > len(uInfos) {
			end = len(uInfos)
		}

		fmt.Println("job add", i, end)
		pieceUsers := uInfos[t*splitNum : end]
		fmt.Println(pieceUsers)
		t++
	}
}
