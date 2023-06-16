package main

import "fmt"

func main() {
	day := 7

	switch day {
	case 1:
		fmt.Println("今天是星期一")
	case 2:
		fmt.Println("今天是星期二")
		fallthrough
	case 3:
		fmt.Println("今天是星期三")
		fallthrough
	case 4:
		fmt.Println("今天是星期四")
		fallthrough
	case 5:
		fmt.Println("今天是星期五")
		fallthrough
	case 6:
		fmt.Println("今天是星期六")
	case 7:
		fmt.Println("今天是星期日")
		break
	default:
		fmt.Println("输入无效")
		goto end
	}

	fmt.Println("这句话会被执行吗？")

end:
	fmt.Println("程序结束")
}

func main3() {
	num := 2

	switch num {
	case 1:
		fmt.Println("数字为1")
		fallthrough // 穿透到下一个case,无论下一个case语句条件是否匹配
	case 2:
		fmt.Println("数字为2")
		break
		fallthrough // 穿透到下一个case，无论下一个case语句条件是否匹配
	case 3:
		fmt.Println("数字为3")
	default:
		fmt.Println("数字不在范围内")
	}
}

func main2() {

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
