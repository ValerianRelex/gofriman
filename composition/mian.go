package main

import (
	"fmt"
	"composition/store"
)

func main() {
	fmt.Println("Hi composition!")

	kayak := store.NewProduct("Kayak", "Watersports", 150)
	lifeJacket := &store.Product{ Name: "Lifejacket", Category: "Watresports"}

	for _, p := range []*store.Product{kayak, lifeJacket} {
		fmt.Println("Name: ", p.Name, "Category: ", p.Category, "Price: ", p.Price(0.02))
	}


	boats := []*store.Boat {
		store.NewBoat("Kayak", 159, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}

	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name,
					"Direct:", b.Name, // напрямую используем вложенный тип поля... он встроен в тип Product
					"asdf", b.Price(600)) // ебанул им налогом!
	}

}