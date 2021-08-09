package main

func main() {
	println(sum(10, 20))
	println(sum(10, 20, 30))
	println(sum(10, 20, 30, 40))
}

func sum(nos ...int) int {
	//nos => slice (collection) of int
	result := 0
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
