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

	rentalBoat := store.NewRentalBoat("Titanic", 500500, 450, true, true, "N/A", "N/A")
	fmt.Println("Name of boat:", rentalBoat.Name, "=", rentalBoat.Boat.Product.Name) // поле продвинуто и не требуется ужасный доступ через все типы
	fmt.Println("Price of boat:", rentalBoat.Name, "=", rentalBoat.Price(0.5)) // тот же эффект касается методов!

	// Магия интерфейса и композиции!
	products := map[string]store.ItemForSale {
        "Kayak": store.NewBoat("Kayak", 279, 1, false),
        "Ball": store.NewProduct("Soccer Ball", "Soccer",19.50),
    }

	for key, val := range products {
		fmt.Println("Key:", key,"Price:", val.Price(0.5))
	}
}