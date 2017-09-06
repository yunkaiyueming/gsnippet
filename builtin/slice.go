package main

import (
	"fmt"
)

func main() {
	data := []string{"1", "2", "3", "4"}
	changeSlice(&data)
	fmt.Println(data)

	appendNoChangeSlice(data)
	fmt.Println(data)

	change2Slice(data)
	fmt.Println(data)
}

func changeSlice(data *[]string) { //会改变全局变量值
	*data = append(*data, "5")
}

func appendNoChangeSlice(data []string) { //不会改变全局变量的值
	data = append(data, "6")
}

func change2Slice(data []string) { //修改slice中值，会改变
	data[0] = "aaaa"
}
