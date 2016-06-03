package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	num = 10000000
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	GoChannel()
}

func GoFetchUrl() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			// Fetch the URL.
			resp, _ := http.Get(url)
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func GoChannel() {
	st := time.Now().UnixNano()
	TestChan()
	fmt.Printf("task cost %d \r\n", (time.Now().UnixNano()-st)/int64(time.Millisecond))
}

func TestChan() {
	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(1)

	go func() {
		for _ = range c {
		}
		wg.Done()
	}()

	for i := 0; i < num; i++ {
		c <- "123"
	}

	close(c)
	wg.Wait()
}
