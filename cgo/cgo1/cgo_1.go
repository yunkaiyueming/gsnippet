package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
    printf("%s\n", str);
}

void echo(){
	print("cgo");
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func TestCPrint() {
	s := "hello"
	cs := C.CString(s) //c源码嵌入go中，go通过C调用c标准库，可以调用通过
	defer C.free(unsafe.Pointer(cs))
	C.print(cs)
}

func main() {
	fmt.Println("======")
	C.echo()
	TestCPrint()
	fmt.Println("======")
}

/*output:
======
======
cgo
hello
*/
