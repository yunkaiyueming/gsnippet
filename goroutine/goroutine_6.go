package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan []string, 5)

	start := time.Now().Unix()
	dates := []int{100000, 200000, 300000, 400000, 500000, 600000, 1000000}
	for _, date := range dates {
		wg.Add(1)
		go worktimeCalc(&wg, ch, date)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	mapCh := make([]string, 0)
	for data := range ch {
		//fmt.Println("receive:", data)
		mapCh = append(mapCh, data...)
	}

	end := time.Now().Unix()
	fmt.Println(end, start, end-start)
}

func worktimeCalc(wg *sync.WaitGroup, ch chan []string, date int) {
	defer wg.Done()
	var date_str []string
	for i := 0; i < date; i++ {
		date_str = append(date_str, strconv.Itoa(date))
	}

	time.Sleep(time.Duration(60 * 20))
	//fmt.Println("send", date_str)
	ch <- date_str
}
