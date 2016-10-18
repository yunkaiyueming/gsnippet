package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var total_tickets int = 100000
var mutex = &sync.Mutex{}

func sell_tickets(i int) {
	for total_tickets > 0 {
		mutex.Lock()
		if total_tickets > 0 { //如果有票就卖
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets-- //卖一张票

			fmt.Println("id:", i, "  ticket:", total_tickets)

			string_i := strconv.Itoa(i)
			string_ti := strconv.Itoa(total_tickets)
			msg := "id:" + string_i + "   " + "total_tickets:" + string_ti + "\n"
			append_write_file(msg)
		}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix()) //生成随机种子

	for i := 0; i < 10; i++ { //并发5个goroutine来卖票
		fmt.Println(i)
		go sell_tickets(i)
	}

	//等待线程执行完
	var input string
	fmt.Scanln(&input)

}

func append_write_file(msg string) {
	fil, err := os.OpenFile("E:/www2/GitHub/go_code/src/2.log", os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer fil.Close()
	fil.WriteString(msg)
}
