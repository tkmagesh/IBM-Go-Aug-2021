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

	fmt.Printf("len(names) => %d, cap(names) => %d\n", len(names), cap(names))
	names = append(names, "x", "y", "z")
	fmt.Printf("names = %v, len(names) => %d, cap(names) => %d\n", names, len(names), cap(names))

	list := make([]int, 5, 100)
	fmt.Println(list, len(list), cap(list))
	list[0] = 0
	list[1] = 10
	list[2] = 20
	list[3] = 30
	list[4] = 40

	//list[5] = 50
	list = append(list, 50)
	fmt.Println(list, len(list), cap(list))
	x := list[0:9]
	fmt.Println(x)

	//Maps
	//key/value pair collection
	fmt.Println("Maps")
	var cityRanks map[string]int = map[string]int{
		"Bengaluru": 4,
		"Udupi":     1,
		"Mangaluru": 2,
		"Mysuru":    3,
	}

	fmt.Println(cityRanks)
	rankOfMysuru := cityRanks["Mysuru"]
	fmt.Println("Rank of Mysuru => ", rankOfMysuru)

	//adding a new key/value pair
	cityRanks["Pune"] = 5
	fmt.Println(cityRanks)

	//check if a key exists
	if rank, ok := cityRanks["Ahmedabad"]; ok {
		fmt.Println("Rank of Ahmedabad => ", rank)
	} else {
		fmt.Println("Ahmedabad is not present in the map")
	}

	//remove a key/value pair
	delete(cityRanks, "Mangaluru")
	fmt.Println("After removing mangaluru, cityRanks => ", cityRanks)

	//iterate over map
	for key, value := range cityRanks {
		fmt.Printf("Value at %s is %d\n", key, value)
	}
}
