package main

import "fmt"

func main() {
	/* implement the functions without using any package level variables */
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4

	fmt.Println(decrement()) //=> 3
	fmt.Println(decrement()) //=> 2
	fmt.Println(decrement()) //=> 1
	fmt.Println(decrement()) //=> 0
	fmt.Println(decrement()) //=> -1
}
