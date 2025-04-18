package main

import "encoding/json"

type DiscountedProduct struct {
	// *Product
	*Product `json:"fufel"` // Использование тега структуры
	// Discount float64 `json:"-"` // Пропуск поля
	Discount float64 `json:",string"` // Кодировать будет в JSON как строка
}

//Реализация интерфейса Marshaler
func (dp *DiscountedProduct) MarshalJSON() (jsn []byte, err error) {
	if dp.Product != nil {
		m := map[string]interface{}{
			"product": dp.Name,
			"cost":    dp.Price - dp.Discount,
		}
		jsn, err = json.Marshal(m)
	}
	return
}
