package main

import (
	"fmt"
	"time"
)

func main() {
	//在 Go 语言中，如果不加特殊前缀，都是10进制表示，例如：“100”
	//整数可以直接用指数形式，例如：“1E9”，意思是 1 * （10^9)，1 乘以 10 的 9 次方
	//整数的加减法用 + 和 - 号，乘法用 * 号，除法用 / 号， 得到的商是整数，例如 5 / 2 = 2，而 % 号是求余(取模), 例如 5 % 2 = 1
	time.Sleep(1e9)
	fmt.Println("ok")

	time.Sleep(time.Second * 1)
	fmt.Println("ok")
}
