package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg.Add(2)
	go Print("Hello", 500*time.Millisecond, ch1, ch2, wg)
	go Print("world", 750*time.Millisecond, ch2, ch1, wg)
	ch1 <- 1
	wg.Wait()
	/*
		var input string
		fmt.Scanln(&input)
	*/
}

func Print(str string, delay time.Duration, ch1, ch2 chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-ch1
		time.Sleep(delay)
		fmt.Println(str)
		ch2 <- 1
	}
	wg.Done()
}

/*
Output:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
