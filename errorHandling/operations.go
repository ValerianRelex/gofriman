package main

import (
	_ "errors"
	"fmt"
)

// type CategoryError struct {
// 	requestedCategory string
// }

// func (e *CategoryError) Error() string {
// 	return "Category " + e.requestedCategory + " does not exist"
// }

// передача сообщений об ошибке через каналы
type ChannelMessage struct {
	Category string
	Total    float64
	// *CategoryError
	CategoryError error
}

// func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
// 	productCount := 0
// 	for _, p := range slice {
// 		if p.Category == category {
// 			total += p.Price
// 			productCount++
// 		}
// 	}
// 	if productCount == 0 {
// 		err = &CategoryError{requestedCategory: category}
// 	}
// 	return
// }

func (slice ProductSlice) TotalPrice(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		// здесь используем встроенную библиотеку вызова ошибок
		// err = errors.New("Can't find category")

		// или так, чтоб было красиво, кстати, тоже возвращает тип error
		err = fmt.Errorf("Cannot find category: %v", category)
	}
	return
}

func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
