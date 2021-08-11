package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	c := make(chan int)
	wg.Add(1)
	go add(100, 200, c)
	wg.Wait()
	result := <-c
	fmt.Println(result)
}

func add(x, y int, c chan int) {
	result := x + y
	c <- result
	wg.Done()
}
