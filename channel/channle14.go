package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		wg.Done()
	}()

	timeout := time.Duration(2) * time.Second
	fmt.Printf("Wait for waitgroup (up to %s)\n", timeout)

	for{
	select {
	case <-done:
		fmt.Printf("Wait group finished\n")
	case <-time.After(timeout):
		fmt.Printf("Timed out waiting for wait group\n")
	}
	}

	fmt.Printf("Free at last\n")
}
