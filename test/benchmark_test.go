package test

import (
	"testing"
)

/*
基准测试 Benchmark用来检测函数/方法的性能
基准测试用例函数必须以Benchmark开头
go test默认不会执行基准测试的函数，需要加上参数-test.bench，语法:-test.bench="test_name_regex"，
例如go test -test.bench=".*"表示测试全部的基准测试函数
在基准测试用例中，在循环体内使用testing.B.N，使测试可以正常的运行
*/
//go test -test.bench=".*"
//go test -bench .
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		b.Logf("Sum 4  to %d: %d\n", b.N, MyPlus(4, b.N))
	}
}

//每次的平均执行时间是7.80纳秒
func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		MyDive(4, 2)
		b.Log(b.N)
	}
}

func BenchmarkMySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MySub(4, 2)
		b.Log(i, b.N)
	}
}
