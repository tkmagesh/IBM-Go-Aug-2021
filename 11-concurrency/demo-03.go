package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go Print("Hello", wg)

	wg.Add(1)
	go Print("World", wg)
	wg.Wait()
	fmt.Println("Exiting the main fn")
}

func Print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting - ", str)
	wg.Done()
}
