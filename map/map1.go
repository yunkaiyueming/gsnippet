package main

import (
	"fmt"
	"strconv"
)

var mapData = make(map[string]string)

func main() {
	go writeA()
	go writeB()
	// go readA()
	// go readB()
	for {
	}
}

func writeB() {
	for i := 1; i <= 1000; i++ {
		id := "B" + strconv.Itoa(i)
		mapData[id] = id
	}
}

func writeA() {
	for i := 1; i <= 1000; i++ {
		id := "A" + strconv.Itoa(i)
		mapData[id] = id
	}
}

func readA() {
	for i := 1; i <= 1000; i++ {
		id := "A" + strconv.Itoa(i)
		fmt.Println(mapData[id])
	}
}

func readB() {
	for i := 1; i <= 1000; i++ {
		id := "B" + strconv.Itoa(i)
		fmt.Println(mapData[id])
	}
}
