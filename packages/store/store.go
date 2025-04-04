// Package store provides types and methods
// commonly required for online sales
package store

var standartTax = newTaxRate(0.25, 100)

// Product describes an item for sale
type Product struct {
	Name, Category string // Name and type of the product
	price          float64
}


func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return standartTax.calcTax(p)
}

func (pppp *Product) SetPrice(newPrice float64) {
	pppp.price = newPrice
}
