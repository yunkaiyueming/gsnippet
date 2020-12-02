package main

//来关闭cancel管道向多个Goroutine广播退出的指令。不过这个程序依然不够稳健：当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能保证被完成，因为main线程并没有等待各个工作Goroutine退出工作完成的机制
import (
    "fmt"
    "sync"
    "time"
)

func worker(wg *sync.WaitGroup, cancel chan bool) {
    defer wg.Done()

    for {
        select {
        case <-time.After(time.Second):
            fmt.Println("run time over")
        //default:
            //time.Sleep(time.Second)
            //fmt.Println("hello")
        case <-cancel:
            fmt.Println("worker exit")
            return
        }
    }
}

func main() {
    cancel := make(chan bool)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(&wg, cancel)
    }

    time.Sleep(time.Second*2)
    close(cancel)
    wg.Wait()
}