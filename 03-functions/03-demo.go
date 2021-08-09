package main

import "fmt"

func main() {

	//functions assigned to variable
	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//anonymous functions
	func() {
		fmt.Println("Anonymous function invoked")
	}()

	func(x, y int) {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}(100, 200)
}
