package main

import "fmt"

func main() {
	/* implement the functions without using any package level variables */
	increment, decrement := getCounter()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4

	fmt.Println(decrement()) //=> 3
	fmt.Println(decrement()) //=> 2
	fmt.Println(decrement()) //=> 1
	fmt.Println(decrement()) //=> 0
	fmt.Println(decrement()) //=> -1
}

func getCounter() (func() int, func() int) {
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	decrement := func() int {
		counter--
		return counter
	}
	return increment, decrement
}
