package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "string"
	x = true
	x = struct{}{}
	x = []int{1, 2, 3}
	fmt.Println(x)

	var val interface{}
	val = 100
	if no, ok := val.(int); ok {
		result := no + 200
		fmt.Println(result)
	} else {
		fmt.Println(val, " is not a number type")
	}

	val = struct{}{}
	switch v := val.(type) {
	case int:
		fmt.Printf("Double of val is %d\n", v*2)
	case string:
		fmt.Printf("Len of string %q is %d\n", v, len(v))
	case []int:
		fmt.Println("List of numbers => ", v)
	default:
		fmt.Println("Unknown type")
	}

}
