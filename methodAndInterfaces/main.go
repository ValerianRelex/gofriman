package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

// функция
func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:",
		product.category,
		"Price", product.price)
}

// метод
func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:",
		product.category,
		"Price", product.price)
}

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price) // функция
		p.printDetails() // метод
	}
}
