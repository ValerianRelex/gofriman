package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2) // создаем канал с буфером
	for category, group := range data {
		// storeTotal += group.TotalPrice(category) // 

		go group.TotalPrice(category, channel) // создаем горутину - добавим к основной горутине еще одну! 
		fmt.Println("Создал горутинку")
	}

	time.Sleep(time.Millisecond * 5000)

	for i := 0; i < len(data); i++ {
		fmt.Println(len(channel), cap(channel))
        fmt.Println("-- channel read pending", len(channel), "items in buffer, size", cap(channel))

		// получение из канала является блокирующей операцией
		storeTotal = storeTotal + <- channel // или кратко storeTotal += <- channel
		fmt.Println("-- channel read complete", storeTotal)
	}


	fmt.Println("Все данные получены и после ручной паузы в 3 секунды, выведется результат! \n")
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		// fmt.Println("Категория:", category, "product:", p.Name) // ключи из карты извлекаются хаотично, поэтому вывод в консоль всегда будет разный, почти
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "channel send complete")
}
