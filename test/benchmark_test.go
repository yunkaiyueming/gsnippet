package test

import (
	"testing"
)

//go test -test.bench=".*"
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		MyDive(4, 2)
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
	}
}
