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

	var sum = func(nos ...int) int {
		result := 0
		for idx := 0; idx < len(nos); idx++ {
			result += nos[idx]
		}
		return result
	}
	fmt.Println(sum(10, 20, 30, 40, 50))
}
