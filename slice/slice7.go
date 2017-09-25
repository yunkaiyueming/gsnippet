package main

import (
	"fmt"
)

func main() {
	data := []string{"0", "1", "2", "3"}
	deleteIndex := 3

	//tmp := make([]string, 0)
	tmp := append(data[0:deleteIndex], data[deleteIndex+1:]...)
	fmt.Println(deleteIndex, tmp)

}
