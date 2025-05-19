package main

import (
	"fmt"
	"reflect"
	"strings"
)

func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i < elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, 	fmt.Sprintf("%v: %v", fieldName, fieldVal))
			}
			PrinterLn("%v: %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			PrinterLn("%v: %v", elemType.Name(), elemValue)
			if (reflect.TypeOf(false) == elemValue.Type()) {
				fmt.Println("asdfasdfasdfasdfasdf==========================")
			}
			PrinterLn("Тип совпадает с преобразованием = %v:  %v", elemType.AssignableTo(reflect.TypeOf(elemValue)), elemValue)
		}
	}
}

func printEnotherDetails(values ...interface{}) {
for _, elem := range values {
	elemValue := reflect.ValueOf(elem)
	switch elemValue.Kind() {
		case reflect.Bool:
			var val bool = elemValue.Bool()
			PrinterLn("Bool: %v", val)
		case reflect.Int:
			var val int64 = elemValue.Int()
			PrinterLn("Int: %v", val)
		case reflect.Float32, reflect.Float64:
			var val float64 = elemValue.Float()
			PrinterLn("Float: %v", val)
		case reflect.String:
			var val string = elemValue.String()
			PrinterLn("String: %v", val)
		case reflect.Ptr:
			var val reflect.Value = elemValue.Elem()
			if (val.Kind() == reflect.Int) {
				PrinterLn("Pointer to Int: %v", val.Int())
			}
		default:
			PrinterLn("Other: %v", elemValue.String())
	}
}
}

type Payment struct {
	Currency string
	Amount   float64
}

func main() {
	product := Product{Name: "Kayak", Category: "Watersports", Price: 279,}
	customer := Customer{Name: "Alice", City: "New York"}
	payment := Payment { Currency: "USD", Amount: 100.50 }

    printDetails(product, customer, payment, 10, true)

	number := 100
    printEnotherDetails(true, 10, 23.30, "Alice", &number, product)
}
