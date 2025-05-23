package main

import "sort"

type Product struct {
	Name  string
	Price float64
}
type ProductSlice []Product

// композиция типа?
func ProductSlices(p []Product) {
	sort.Sort(ProductSlice(p))
}
func ProductSlicesAreSorted(p []Product) {
	sort.IsSorted(ProductSlice(p))
}
func (products ProductSlice) Len() int {
	return len(products)
}
func (products ProductSlice) Less(i, j int) bool {
	return products[i].Price < products[j].Price
}
func (products ProductSlice) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}
