package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Units    int
	Cost     float64
	Category string
}

/* type PerishableProduct struct {
	Prod   Product
	Expiry string
} */

type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, units int, cost float64, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, units, cost, category}, expiry}
}

func main() {
	/*
		grapes := PerishableProduct{Prod: Product{201, "Grapes", 5, 60, "Food"}, Expiry: "2 Days"}
		fmt.Println(grapes)
		fmt.Println("Cost of grapes => ", grapes.Prod.Cost)
	*/

	//grapes := PerishableProduct{Product: Product{201, "Grapes", 5, 60, "Food"}, Expiry: "2 Days"}
	grapes := NewPerishableProduct(201, "Grapes", 5, 60, "Food", "2 Days")
	fmt.Println(grapes)
	fmt.Println("Cost of grapes => ", grapes.Product.Cost)

	/*
		grapes := PerishableProduct{Product{201, "Grapes", 5, 60, "Food"}, "2 Days"}
		fmt.Println(grapes)
		fmt.Println("Cost of grapes => ", grapes.Cost)
	*/
}
