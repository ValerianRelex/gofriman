package main

import p "fmt" // заюзали псевдоним!

func main() {

	recoveryFunc := func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok { // в скобочках - это происходит утверждение типа
				p.Println("Error:", err.Error())
			} else if str, ok := arg.(string); ok { // в скобочках - это происходит утверждение типа
				p.Println("Message:", str)
			} else {
				p.Println("Panic recovered")
			}
		}
	}
	defer recoveryFunc()

	categories := []string{"Watersports", "Chess", "Snowboard"}

	for _, cat := range categories {
		total, err := Products.TotalPrice(cat)
		if err == nil {
			p.Println(cat, "Total:", ToCurrency(total))
		} else {
			p.Println("Ошибка, сэр!", err.Error(), "это же отсутствующая категория!") // todo: непонимаю разницу в вызове err и err.Error() - работает одинаково!
		}
	}

	// А ну как, запилим асинхронно!
	catchan := make(chan ChannelMessage, 10)

	go Products.TotalPriceAsync(categories, catchan) // создали горутину!

	for date := range catchan {
		if date.CategoryError == nil {
			p.Println("вывод из канала:", date.Category, "=", date.Total)
		} else {
			panic(date.CategoryError)
			p.Println("Ошибка:", date.CategoryError)
		}
	}

}
