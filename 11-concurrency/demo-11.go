package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	result1 := sum(dataSet1...)
	result2 := sum(dataSet2...)
	result := result1 + result2
	fmt.Println(result)
}

func sum(nos ...int) int {
	var sum int
	for _, n := range nos {
		sum += n
	}
	return sum
}
