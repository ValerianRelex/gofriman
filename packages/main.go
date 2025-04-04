package main

import (
	"fmt"
	"packages/store"
	"packages/store/cart"
	_ "packages/data"
	"github.com/fatih/color" // подключили внешние пакеты, для использования их в проекте...
)

func init() {
	color.Green("я функция инициализации, вызываюсь сама по себе перед стартом приложения")
}

func main() {
	product := store.Product {
        Name: "Cereal",
        Category: "Food",
    }

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)

	// ожидаемо!
	var foodProduct *store.Product = &product
	foodProduct.Name = "Buble gum"
	foodProduct.SetPrice(100500)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)

	newProduct := store.NewProduct("Apple", "Food", 150)

	milk := store.NewProduct("Milk", "Drink", 99)

	fmt.Println("newProduct price =", newProduct.Price())
	fmt.Println("milk price =", milk.Price())

	cart1 := cart.Cart {
		CustomerName: "Valeriy",
		Products: []store.Product{*foodProduct},
	}

	fmt.Println(cart1.CustomerName, cart1.GetTotal())

	// поиграем с внешней библиотекой

	color.Red("Name: " + cart1.CustomerName)

}