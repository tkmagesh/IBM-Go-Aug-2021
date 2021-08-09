package main

import "fmt"

func main() {
	fn := func() {
		fmt.Println("fn is invoked")
	}
	exec(fn)
	addProcess := wrapProcess("Adding numbers", add)
	addProcess(100, 200)
	addProcess(10, 20)
	addProcess(30, 40)
	addProcess(20, 50)

	subtractProcess := wrapProcess("Subtracting numbers", subtract)
	subtractProcess(200, 500)

}

func wrapProcess(msg string, fn func(int, int) int) func(x, y int) {
	return func(x, y int) {
		fmt.Println(msg)
		result := fn(x, y)
		fmt.Println("Result:", result)
	}
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
