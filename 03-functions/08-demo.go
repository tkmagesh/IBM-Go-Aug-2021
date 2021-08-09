package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Divide by zero error")

func main() {
	fmt.Println(divide(10, 2))
	result, err := divide(10, 0)
	if err == DivideByZeroError {
		fmt.Println("Divisor cannot be zero!")
	} else if err != nil {
		fmt.Println("Something went wrong!")
	} else {
		fmt.Println(result)
	}
}

func divide(x, y int) (int, error) {
	if y != 0 {
		return x / y, nil
	}
	return 0, DivideByZeroError
}
