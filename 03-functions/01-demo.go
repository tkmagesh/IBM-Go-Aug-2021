package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(divide(100, 7))

}

/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/* func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
} */

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return quotient, remainder
}
