package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println(divide(100, 0))
}

func divide(x, y int) int {
	if y == 0 {
		panic("Unable to divide by zero")
	}
	return x / y
}
