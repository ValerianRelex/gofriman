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
	fmt.Println("Name:", product.name, "Category:", product.category, "Price", product.calcTax(0.2, 100))
}

// метод с параметрами и возвращаемым значением
func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

// определили всевдоним для среза с типом Product
type ProductList []Product

// а это метод, принимает указатель на срез с певдонимом ProductList и возвращает словарь ключ - стринг, значение - сумма
func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

// определим интерфейс
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	products := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price) // функция
		p.printDetails()                                                        // метод
		// Product.printDetails(p) // так тоже можно, если бы ожидаемый аргумент не был бы указателем
	}

	fmt.Println("Name:", products[0].name,
		"Category:", products[0].category,
		"Price", products[0].price) // функция

	// работа с мапой, через рандж, возвращает ключ и значение
	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

	prods := ProductList(getProducts()) // приведения типа к псевдониму

	// и тогда мы сможем обратиться к методу prods.calcCategoryTotals() определенному для псевдонима
	for category, total := range prods.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

	// используем данные из одного пакета, но из разных файлов
	kayak := Products{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}
	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))

	// щаз разрыв башки, определили срез интерфейса.... а, походу мы можем засунуть сюда эти типы данных, т.к. они все реализуют
	// интерфейс Expense, блядь, пиздец... догадался наконец-то. А еще характерно, что внутри них не написано, что они
	// реализуют интерфейс Expense

	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	fmt.Println("Total: ", calcTotal(expenses))

	product := Product{"Alishev", "Baul", 500.0}

	var expense Expense = &product // а без амперсанда создастся новое значение

	product.price = 100

	fmt.Println("product.price =", product.price, "expense.price =", expense.getCost(true))

	var total float64 = 0
	for _, item := range expenses {
		total += item.getCost(false)
		
		// УТВЕРЖДЕНИЕ ТИПА с помощью swith case
		switch value := item.(type) {
		case Service: fmt.Println("Type Service", value.description)
		case Product: fmt.Println("Type Product", value.category)
		default: fmt.Print("Not identity")
		}

	}
}



func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(false)
		if s, ok := item.(Service); ok { // УТВЕРЖДЕНИЕ ТИПА, сужаем тип, и выводим только тип Service
			fmt.Println(s.description, s.durationMonths, s.monthlyFee)
		}

		// тоже самое можно сделать с помощью swith case
	}
	return
}