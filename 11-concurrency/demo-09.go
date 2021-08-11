package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	count := make(chan int)
	wg.Add(1)
	go fn(count)
	fmt.Println("[main] Writing to the count (after 4 seconds)")
	time.Sleep(time.Second * 4)
	count <- 10
	wg.Wait()
	fmt.Println("[main] Exiting")
}

func fn(count chan int) {
	fmt.Println("[fn] Attempting to read the count")
	max := <-count
	fmt.Println("[fn] The count is", max)
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
