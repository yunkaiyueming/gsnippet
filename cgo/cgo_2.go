package main

/*
#include <stdio.h>

void sayHi() {
    printf("Hi");
}
*/
import "C"

func main() {
	C.sayHi()
}
