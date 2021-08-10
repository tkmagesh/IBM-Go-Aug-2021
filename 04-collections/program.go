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

	//Slices
	//var names []string
	var names []string = []string{"Magesh", "Ramesh", "Suresh"}
	fmt.Println(names)

	//names = append(names, "Rajesh")
	//names = append(names, "Rajesh", "Ganesh")
	newNames := []string{"Rajesh", "Ganesh"}
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to the end
		[ : hi] => from the beginning to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] == [lo]
		[:] => copy of the slice
	*/
	fmt.Println("[1:3] => ", names[1:3])
	fmt.Println("[3:] => ", names[3:])
	fmt.Println("[:3] => ", names[:3])
}
