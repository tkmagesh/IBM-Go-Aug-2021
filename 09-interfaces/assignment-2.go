package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type Products []Product

func (p Products) Len() int { return len(p) }
func (p Products) Less(i, j int) bool {
	return p[i].Id < p[j].Id
}
func (p Products) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p Products) Sort() {
	sort.Sort(p)
}

type ByName struct {
	Products
}

func (p ByName) Less(i, j int) bool {
	return p.Products[i].Name < p.Products[j].Name
}

func (p Products) SortByName() {
	sort.Sort(ByName{p})
}

func main() {
	products := Products{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Default sort")
	products.Sort()
	fmt.Println(products)

	fmt.Println()
	fmt.Println("Sort By Name")
	products.SortByName()
	fmt.Println(products)
	//Sort the products by any attribute
	//IMPORTANT : Do NOT implement your own sorting logic. Use "sort" package functions
	//HINT : look at sort.Sort() function

}

//Method of product
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//Method of []Product
func (products Products) String() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%s\n", p)
	}
	return result
}
