package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wait := sync.WaitGroup{}
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wait.Done()
			wait.Add(1)

			cond.L.Lock()
			fmt.Println("Waiting start...")
			cond.Wait() //Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L
			fmt.Println("Waiting end...")
			cond.L.Unlock()

			fmt.Println("Goroutine run. Number:", i)
		}(i)
	}

	//OneNotice(cond)
	AllNotice(cond)

	wait.Wait()
	fmt.Println("end")
}

func OneNotice(cond *sync.Cond) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		cond.Signal() //Signal 唤醒一个再此cond 对象上等待的goroutine
	}
}

func AllNotice(cond *sync.Cond) {
	time.Sleep(time.Second * 2)
	cond.Broadcast() //Brocast 唤醒所有在这个cond 对象上等待的 goroutine；
}
