package main

import "fmt"
import "os"

func main() {
	//defer是在return之前返回的，return语句不是原子语句，go以栈的形式返回数据
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
}

func TestCreateDefer() {
	f := createFile("E:/www2/GitHub/go_code/src/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

//1
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

//5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

//1
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
