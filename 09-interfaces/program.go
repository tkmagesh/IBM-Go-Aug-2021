package main

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c *Circle) Perimeter() float64 {
	return c.Radius * 2 * math.Pi
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return r.Height*2 + r.Width*2
}

/* func PrintAreaOfCircle(c *Circle) {
	println("Area = ", c.Area())
}

func PrintAreaOfRectangle(r *Rectangle) {
	println("Area = ", r.Area())
}
*/

type ShapeWithArea interface {
	Area() float64
}

func PrintArea(s ShapeWithArea) {
	println("Area = ", s.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(s ShapeWithPerimeter) {
	println("Perimeter = ", s.Perimeter())
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}

func main() {
	c := &Circle{Radius: 3}
	//PrintAreaOfCircle(c)
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)
	r := &Rectangle{Height: 3, Width: 4}
	//PrintAreaOfRectangle(r)
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
