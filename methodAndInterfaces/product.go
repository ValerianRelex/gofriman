package main

type Products struct {
	name, category string
	price          float64
}

// реализации интерфейса Expense
func (p Product) getName() string {
	return p.name
}


// реализации интерфейса Expense
func (p Product) getCost(_ bool) float64 {
	return p.price
}
