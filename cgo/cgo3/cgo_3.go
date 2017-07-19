package main

/*
#include "test.h"
*/
import "C"

func main() {
	C.sayHi() //通过调用c文件来调用C库
	C.sayHello()
}
