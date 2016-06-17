package test

import (
	//"fmt"
	"testing"
)

//go test -v
//单元测试文件名_test.go
//测试函数TestFunc(t *testing.T)
func TestPuls(t *testing.T) {
	if ret := MyPlus(1, 2); ret != 3 {
		t.Error("error")
	} else {
		t.Log("ok")
	}
}

func TestDive(t *testing.T) {
	if ret := MyDive(2, 1); ret != 2 {
		t.Error("error")
	} else {
		t.Log("ok")
	}
}
