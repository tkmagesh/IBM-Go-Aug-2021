package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

type OpCounter struct {
	opCount    int
	sync.Mutex // guards
}

func (c *OpCounter) Increment() {
	c.Lock()
	{
		c.opCount++
	}
	c.Unlock()
}

func (c *OpCounter) GetCount() int {
	return c.opCount
}

var opCounter = &OpCounter{}

func main() {
	wg.Add(2)
	go add(100, 200)
	go add(10, 20)
	wg.Wait()
	fmt.Println(opCounter.GetCount())
}

func add(x, y int) {
	result := x + y
	fmt.Println(x, y, result)
	opCounter.Increment()
	wg.Done()
}
