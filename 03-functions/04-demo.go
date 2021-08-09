package main

import "fmt"

func main() {
	fn := func() {
		fmt.Println("fn is invoked")
	}
	exec(fn)

	process(add, 100, 200)
	process(subtract, 200, 500)

}

func process(msg string, fn func(int, int) int, x, y int) {
	fmt.Println("Adding 2 numbers")
	result := fn(x, y)
	fmt.Println("Result:", result)

}

func exec(fn func()) {
	fmt.Println("Executing fn")
	fn()
	fmt.Println("Finished executing fn")
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
