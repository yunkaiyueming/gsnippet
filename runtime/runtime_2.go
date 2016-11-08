package main

import (
	"fmt"
	"runtime"
)

func main() {
	test()
}

func test_1() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
	// Output:
	// skip = 0, pc = 4198453, file = caller.go, line = 10
	// skip = 1, pc = 4280066, file = $(GOROOT)/src/pkg/runtime/proc.c, line = 220
	// skip = 2, pc = 4289712, file = $(GOROOT)/src/pkg/runtime/proc.c, line = 1394
}

func test() {
	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		n := runtime.Callers(skip, pc)
		if n <= 0 {
			break
		}
		fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
	}
	// Output:
	// skip = 0, pc = [4304486 4198562 4280114 4289760]
	// skip = 1, pc = [4198562 4280114 4289760]
	// skip = 2, pc = [4280114 4289760]
	// skip = 3, pc = [4289760]
}
