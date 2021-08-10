package main

import (
	"fmt"
	"os"
)

var operations = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	var userChoice int
	var n1, n2 int
	for {
		userChoice = getUserChoice()
		if userChoice > 4 || userChoice < 1 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		n1, n2 = getOperands()
		result := process(userChoice, n1, n2)
		fmt.Println("Result : ", result)
	}
}

func getOperands() (n1 int, n2 int) {
	fmt.Println("Enter the operands :")
	fmt.Scanf("%d %d", &n1, &n2)
	return
}

func process(userChoice, n1, n2 int) (result int) {
	if operation, ok := operations[userChoice]; ok {
		result = operation(n1, n2)
	}
	return result
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Menu")
	fmt.Println("================================================")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exit")
	fmt.Scanf("%d", &userChoice)
	return userChoice
}
