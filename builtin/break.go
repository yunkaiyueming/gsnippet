package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}

	test_break()
}

func test_break() {
	i := 10
loop:
	for {
		i++
		fmt.Printf("%d\n", i)
		if i > 20 {
			break loop
		}
	}
}
