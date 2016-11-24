package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//获取线程ID
func GoRoutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	GoroutineId, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}

	return GoroutineId

}

func GetProcessId() (int, int) {
	return os.Getpid(), os.Getppid()
}

func main() {
	pid, ppid := GetProcessId()
	rid := GoRoutineID()
	println(ppid, pid, rid)
}
