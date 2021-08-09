//deferred functions

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("main completed")
	}()
	f1()
}

func f1() {
	defer func() {
		fmt.Println("f1 completed")
	}()
	fmt.Println("f1 is invoked")
	f2()
}

func f2() {
	/* defer func() {
		fmt.Println("f2 [1] completed")
	}()
	defer func() {
		fmt.Println("f2 [2] completed")
	}()
	defer func() {
		fmt.Println("f3 [3] completed")
	}() */
	defer fmt.Println("f2 [1] completed")
	defer fmt.Println("f2 [2] completed")
	defer fmt.Println("f3 [3] completed")
	fmt.Println("f2 is invoked")
}
