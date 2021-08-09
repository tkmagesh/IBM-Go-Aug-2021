//package declaration
package main

//import dependent packages
import "fmt"

//package level variables

//package init function

//main function
func main() {
	var name string
	fmt.Printf("Hello, what is your name?\n")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hi %s, Have a nice day!\n", name)
	n1, n2 := 10, 20
	fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
}

//other functions
