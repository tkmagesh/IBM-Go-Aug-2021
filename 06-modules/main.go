package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calculator.AddClient(100, 200))
	fmt.Println(calculator.AddClient(10, 20))
	fmt.Println(calculator.AddClient(1, 20))
	fmt.Println(calculator.GetCounter())
	fmt.Println(utils.IsOdd(101))
	color.Blue("A blue string")
	color.Red("A red string")
}
