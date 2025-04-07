package main

import (
	"fmt"
	"time"
)

func receiveDispathes(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer,
			":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}

func main() {

	// усложним, создадим ограничения для каналов и для переменных канала
	dispatchChannel := make(chan DispatchNotification, 100)

	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	go DispatchOrders(chan <- DispatchNotification(dispatchChannel))
	// receiveDispathes(dispatchChannel)

	productChannel := make(chan *Product)
	go enumerateProducts(productChannel)
	openChannels := 2


	// работа с двумя каналами с помощью select и case
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Dispatch channel has been closed")
				dispatchChannel = nil
				openChannels--
			}
			
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product:", product.Name)
			} else {
				fmt.Println("Product channel has been closed")
				productChannel = nil // если закоментить здесь, то в deafult ветку никогда не попадет и будет infinite loop
				openChannels--
			}
			fmt.Println("from PRODUCT chanel") // быть внимательным, закрытый канал все равно шлет сообщения типа nil 

		default:
			if openChannels == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}

alldone:
	fmt.Println("All values received")

	// dispatchChannel := make(chan DispatchNotification, 100)
	// go DispatchOrders(dispatchChannel)

	// for {
	// 	if details, open := <-dispatchChannel; open {
	// 		fmt.Println("Dispatch to", details.Customer,
	// 			":", details.Quantity,
	// 			"x", details.Product.Name)
	// 	} else {
	// 		fmt.Println("Канал закрыт, данные значит закончились")
	// 		break
	// 	}
	// }
 
	// Цикл for можно использовать с ключевым словом range для перечисления
	//  значений, отправляемых через канал, что упрощает получение значений и
	//  завершает цикл при закрытии канала

	// for details := range dispatchChannel {
	// 	fmt.Println("Dispatch to", details.Customer,
	// 			":", details.Quantity,
	// 			"x", details.Product.Name)
	// }
	// fmt.Println("Channel has been closed")

	// var pauseInSecond int = 20
	// fmt.Println("\nmain function started\n")

	// CalcStoreTotal(Products)

	// time.Sleep(time.Second * time.Duration(pauseInSecond))
	// fmt.Println("\nmain function complete\n")
}
