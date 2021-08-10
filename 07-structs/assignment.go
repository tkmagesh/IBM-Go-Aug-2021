package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	dummyProduct := Product{Name: "Dummy"}
	fmt.Println("IndexOf Function")
	fmt.Println("Index of Marker => ", IndexOf(products, Product{103, "Marker", 50, 20, "Utencil"}))
	fmt.Println("Index of Dummy => ", IndexOf(products, dummyProduct))

	fmt.Println()
	fmt.Println("Includes Function")
	fmt.Println("Includes Marker ? => ", Includes(products, Product{103, "Marker", 50, 20, "Utencil"}))
	fmt.Println("Includes Dummy ? => ", Includes(products, dummyProduct))

	fmt.Println()
	fmt.Println("Filter function")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := Filter(products, stationaryProductPredicate)
	fmt.Println("Stationary Products => ")
	fmt.Println(FormatProducts(stationaryProducts))

	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 100
	}
	costlyProducts := Filter(products, costlyProductPredicate)
	fmt.Println("Costly Products => ")
	fmt.Println(FormatProducts(costlyProducts))

	fmt.Println()
	fmt.Println("Any function")
	// Are there any costly product?
	fmt.Println("Are there any stationary products ? => ", Any(products, stationaryProductPredicate))
	// Are thery any stationary products?
	fmt.Println("Are there any costly products ? => ", Any(products, costlyProductPredicate))

	fmt.Println()
	fmt.Println("All function")
	//Are all the products in the collection are costly products?
	fmt.Println("Are all the products in the collection are costly products ? => ", All(products, costlyProductPredicate))

	//Are all the products in the collection are stationary products?
	fmt.Println("Are all the products in the collection are stationary products ? => ", All(products, stationaryProductPredicate))
}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func FormatProducts(products []Product) string {
	var result string
	for _, p := range products {
		result += Format(p) + "\n"
	}
	return result
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)
func IndexOf(products []Product, product Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

//Includes => return true/false based on the existence of the product in the products collection
func Includes(products []Product, product Product) bool {
	return IndexOf(products, product) != -1
}

//Filter => returns a new collection of products that match the given criteria
//use cases

//filter costly products (cost > 100)

func Filter(products []Product, predicate func(Product) bool) []Product {
	var result []Product
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

//Any => return true/false based on the existence of the given product that satisfies the given criteria in the products collection
func Any(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

//Use Cases

//All => return true/false based on the existence of all the given products that satisfy the given criteria in the products collection
func All(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

//Use Cases
