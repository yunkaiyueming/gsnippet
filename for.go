package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for i = 0; i < 5; i++ {
		var t int
		fmt.Printf("%d", t) // 0
		t = 10
	}
}
