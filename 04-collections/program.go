package main

import "fmt"

var x int

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iterate over array
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	//range construct
	fmt.Println("using the range construct")
	for _, value := range nos {
		fmt.Println(value)
	}

	//creating a copy of an array
	fmt.Println("creating a copy of an array")
	newNos := nos
	nos[0] = 200
	fmt.Println(nos, newNos)
}
