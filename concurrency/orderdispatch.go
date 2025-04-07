package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

func DispatchOrders(channel chan <- DispatchNotification) { // если нужно ограничить канал на чтение channel chan<- DispatchNotification
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(5) + 5
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		// if i == 1 {
		// 	notification := <- channel // если не запретить чтение из канал, то будет скрытая ошибка
		// 	fmt.Println("Read:", notification.Customer)
		// }
		time.Sleep(time.Millisecond * 750)
	}
	close(channel) // закрытие канала
}
