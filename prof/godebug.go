package main

import("sync")
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}

//查看垃圾回收（GC）信息
//GODEBUG=gctrace=1 go run godebug.go

//查看调度器信息
//GODEBUG=schedtrace=1000 go run godebug.go 