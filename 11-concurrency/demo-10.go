package main

import (
	"fmt"
	"time"
)

func main() {
	Print("Hello", 2*time.Second)
	Print("world", 3*time.Second)
}

func Print(str string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		time.Sleep(delay)
		fmt.Println(str)
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
