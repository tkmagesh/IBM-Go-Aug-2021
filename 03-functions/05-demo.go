package main

import "fmt"

func main() {
	add := fn()
	fmt.Println(add(1200, 200))
}

func fn() func(int, int) int {
	fmt.Println("returning a function")
	return func(x, y int) int {
		return x + y
	}
}
