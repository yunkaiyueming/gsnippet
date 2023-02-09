package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func workForever() {
	for {
		go counter()
		time.Sleep(1 * time.Second)
	}
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	counter()
}

func main() {
	go workForever()
	http.HandleFunc("/get", httpGet)
	http.ListenAndServe("localhost:8000", nil)
}

//外网web访问地址 http://localhost:8000/debug/pprof/

//终端获取
// go tool pprof http://127.0.0.1:8000/debug/pprof/profile 排查cpu占用
// go tool pprof http://127.0.0.1:8000/debug/pprof/heap 排查内存占用
//go tool pprof http://localhost:8000/debug/pprof/allocs 排查gc回收
//go tool pprof http://localhost:6060/debug/pprof/goroutine 排查gorountine分配

//go tool pprof -http=:8081 /Users/ray/pprof/pprof.goroutine.003.pb.gz 查看火焰图 该命令将在所指定的端口号运行一个 PProf 的分析用的站点。
