package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)

	for no := range ch {
		fmt.Println(no)
	}

	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	stop := time.After(20 * time.Second)
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-stop:
			fmt.Println("Stopped")
			close(ch)
			return
		}
	}

}
