package main

import (
	"fmt"
)

func main() {
	go Print("Hello")
	go Print("World")
	//time.Sleep(100 * time.Millisecond)
	var input string
	fmt.Scanln(&input)
}

func Print(str string) {
	fmt.Println(str)
}
