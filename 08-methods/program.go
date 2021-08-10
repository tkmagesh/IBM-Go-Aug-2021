package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

func (p *Product) applyDiscount(discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

//type aliasing
type MyStr string

func (s MyStr) length() int {
	return len(string(s))
}

func main() {
	pen := Product{Id: 100, Name: "Pen", Units: 100, Cost: 5, Category: "Stationary"}
	fmt.Println(pen)
	//applyDiscount(&pen, 10) //apply a discount of 10% on the product
	pen.applyDiscount(10)
	fmt.Println("After applying 10% discount, pen => ", pen)

	//aliasing
	s := MyStr("Hello") //converting the typeof string to MyStr
	fmt.Println(s.length())

}
