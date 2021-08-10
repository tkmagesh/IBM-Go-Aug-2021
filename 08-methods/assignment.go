package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type Products []Product

func main() {
	products := Products{
		{105, "Pen", 5, 50, "Stationary"},
		{107, "Pencil", 2, 100, "Stationary"},
		{103, "Marker", 50, 20, "Utencil"},
		{102, "Stove", 5000, 5, "Utencil"},
		{101, "Kettle", 2500, 10, "Utencil"},
		{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	dummyProduct := Product{Name: "Dummy"}
	fmt.Println("IndexOf Function")
	fmt.Println("Index of Marker => ", products.IndexOf(Product{103, "Marker", 50, 20, "Utencil"}))
	fmt.Println("Index of Dummy => ", products.IndexOf(dummyProduct))

	fmt.Println()
	fmt.Println("Includes Function")
	fmt.Println("Includes Marker ? => ", products.Includes(Product{103, "Marker", 50, 20, "Utencil"}))
	fmt.Println("Includes Dummy ? => ", products.Includes(dummyProduct))

	fmt.Println()
	fmt.Println("Filter function")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println("Stationary Products => ")
	fmt.Println(stationaryProducts.Format())

	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 100
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println("Costly Products => ")
	fmt.Println(costlyProducts.Format())

	fmt.Println()
	fmt.Println("Any function")
	// Are there any costly product?
	fmt.Println("Are there any stationary products ? => ", products.Any(stationaryProductPredicate))
	// Are thery any stationary products?
	fmt.Println("Are there any costly products ? => ", products.Any(costlyProductPredicate))

	fmt.Println()
	fmt.Println("All function")
	//Are all the products in the collection are costly products?
	fmt.Println("Are all the products in the collection are costly products ? => ", products.All(costlyProductPredicate))

	//Are all the products in the collection are stationary products?
	fmt.Println("Are all the products in the collection are stationary products ? => ", products.All(stationaryProductPredicate))
}

//Method of product
func (product *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//Method of []Product
func (products Products) Format() string {
	var result string
	for _, p := range products {
		result += p.Format() + "\n"
	}
	return result
}

//Method of []Product
//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)
func (products Products) IndexOf(product Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

//Method of []Product
//Includes => return true/false based on the existence of the product in the products collection
func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

//Filter => returns a new collection of products that match the given criteria
//use cases

//filter costly products (cost > 100)
//Method of []Product
func (products Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

//Method of []Product
//Any => return true/false based on the existence of the given product that satisfies the given criteria in the products collection
func (products Products) Any(predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

//Use Cases

//Method of []Product
//All => return true/false based on the existence of all the given products that satisfy the given criteria in the products collection
func (products Products) All(predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

//Use Cases
