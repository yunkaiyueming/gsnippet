package main

import (
	"fmt"
)

func main() {
	int_a := 1
	string_b := "afdfdas"
	slice_c := make([]string, 2, 10)
	slice_f := make([]string, 1)
	map_d := make(map[string]string)
	channel_e := make(chan int)

	slice_c = []string{"aa", "bb", "cc"}
	map_d = map[string]string{"id": "123", "name": "aaa", "age": "20"}

	//内建函数的使用
	println(int_a)
	println(string_b)
	fmt.Println(slice_c)

	println(cap(slice_c))
	println(len(slice_c))
	println(len(string_b))

	copy(slice_f, slice_c)
	fmt.Println(slice_f)

	slice_f = append(slice_f, "aaa")
	fmt.Println(slice_f)

	fmt.Println(map_d)
	delete(map_d, "id")
	fmt.Println(map_d)

	channel_e <- 1
	<-channel_e

	close(channel_e)
}
