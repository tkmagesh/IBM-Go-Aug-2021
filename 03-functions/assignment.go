package main

import (
	"fmt"
	"os"
)

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
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	return
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
