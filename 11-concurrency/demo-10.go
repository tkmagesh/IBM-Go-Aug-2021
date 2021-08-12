package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Print("Hello", 500*time.Millisecond, ch1, ch2)
	go Print("world", 750*time.Millisecond, ch2, ch1)
	ch1 <- 1

	/*
		var input string
		fmt.Scanln(&input)
	*/
}

func Print(str string, delay time.Duration, ch1, ch2 chan int) {
	for i := 0; i < 5; i++ {
		<-ch1
		time.Sleep(delay)
		fmt.Println(str)
		ch2 <- 1
	}
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
