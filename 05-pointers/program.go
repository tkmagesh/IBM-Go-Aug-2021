package main

import "fmt"

func main() {
	var x int = 10
	var xPtr *int = &x
	println(xPtr, x)

	//dereferencing a pointer
	y := *xPtr
	// x == *(&x)
	fmt.Println(y)

	fmt.Println("Before incrementing x = ", x)
	increment(&x)
	fmt.Println("After incrementing x = ", x)
	no1, no2 := 10, 20
	fmt.Println("Before swapping ", no1, no2)
	swap(&no1, &no2)
	fmt.Println("After swapping ", no1, no2)
}

func increment(no *int) {
	*no++
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
