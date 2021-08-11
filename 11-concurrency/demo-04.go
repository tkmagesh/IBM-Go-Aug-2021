package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

//communicate by sharing memory [NOT ADVICABLE in GOLANG]
var opCount = 0

var mutex = &sync.Mutex{}

func main() {
	wg.Add(2)
	go add(100, 200)
	go add(10, 20)
	wg.Wait()
	fmt.Println(opCount)
}

func add(x, y int) {
	result := x + y
	fmt.Println(x, y, result)
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	wg.Done()
}
