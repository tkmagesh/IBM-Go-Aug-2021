package main

import "fmt"

//var appName string = "My App"
var appName = "My App"

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//using type inference
	/*
		var msg = "Hello World!"
	*/

	//the following is applicable only in a function (not at the package level)
	msg := "Hello World!"
	fmt.Println(msg)

	/*
		var label string
		var x int
		var y int
	*/

	/*
		var label string
		var x, y int
	*/

	/*
		var label string
		var (
			x int
			y int
		)
	*/
	/*
		var (
			label string
			x     int
			y     int
		)
	*/

	/*
		var (
			label string
			x, y  int
		)

		label = "Add Result :"
		x = 100
		y = 200
	*/

	/*
		var (
			label string = "Add Result :"
			x     int    = 100
			y     int    = 200
		)
	*/

	/*
		var (
			label string = "Add Result :"
			x, y  int    = 100, 200
		)
	*/

	/*
		var (
			label, x, y = "Add Result :", 100, 200
		)
	*/

	/*
		var label, x, y = "Add Result :", 100, 200
	*/

	label, x, y := "Add Result :", 100, 200
	fmt.Println(label, x+y)

	fmt.Println(appName)
}
