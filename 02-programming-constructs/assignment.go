package main

import "fmt"

func main() {
	var userChoice int
	var n1, n2, result int
	for {
		fmt.Println("Menu")
		fmt.Println("================================================")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Scanf("%d", &userChoice)
		if userChoice > 4 || userChoice < 1 {
			fmt.Println("Exiting...")
			return
		}
		fmt.Println("Enter the operands :")
		fmt.Scanf("%d %d", &n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result : ", result)
	}
}
