package test

import (
	//"fmt"
	"testing"
)

//go test -v
//单元测试文件名_test.go
//测试函数TestFunc(t *testing.T)

/*
测试源文件名必须是_test.go结尾的，go test的时候才会执行到相应的代码
必须import testing包
所有的测试用例函数必须以Test开头
测试用例按照源码中编写的顺序依次执行
测试函数TestXxx()的参数是*testing.T，可以使用该类型来记录错误或者是测试状态
测试格式：func TestXxx (t *testing.T)，Xxx部分可以为任意的字母数字的组合，首字母不能是小写字母[a-z]，例如Testsum是错误的函数名。
函数中通过调用*testing.T的Error，Errorf，FailNow，Fatal，FatalIf方法标注测试不通过，调用Log方法用来记录测试的信息
*/
func TestPuls(t *testing.T) {
	if ret := MyPlus(1, 2); ret != 3 {
		t.Error("error")
	} else {
		t.Log("ok \n")
	}
}

func TestDive(t *testing.T) {
	if ret := MyDive(2, 1); ret != 2 {
		t.Error("error")
	} else {
		t.Log("ok \n ")
	}
}

func TestSum(t *testing.T) {
	if ret := MySub(2, 1); ret != 1 {
		t.Error("error")
	} else {
		t.Log("ok \n")
	}
}
