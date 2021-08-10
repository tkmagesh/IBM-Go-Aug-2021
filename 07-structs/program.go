package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

func main() {
	//var pen Product = Product{}
	//var pen Product = Product{100, "Pen", 100, 5, "Stationary"}
	//var pen Product = Product{Id: 100, Name: "Pen", Units: 100, Cost: 5, Category: "Stationary"}
	pen := Product{Id: 100, Name: "Pen", Units: 100, Cost: 5, Category: "Stationary"}
	fmt.Println(pen)

	fmt.Printf("Product %s costs %v\n", pen.Name, pen.Cost)

	applyDiscount(&pen, 10) //apply a discount of 10% on the product
	fmt.Println("After applying 10% discount, pen => ", pen)
}

func applyDiscount(p *Product, discount float64) {
	/* 	(*p).Cost = (*p).Cost * ((100 - discount) / 100) */
	p.Cost = p.Cost * (100 - discount/100)
}
