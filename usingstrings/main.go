package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("Error for index %v", index)
	}
	return
}

// TODO: выяснить как работают тип интерфейс в аргументе
func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

func main() {
	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	name, _ := getProductName(1)
	fmt.Println(name)
	_, err := getProductName(10)
	fmt.Println(err.Error())

	Printfln("Бинарный вид числа %d = %b, а сейчас напечатаю руну - %c", 3, 3, []rune(name)[0])

	name = "Kayak"
 	Printfln("Pointer: %p", &name)
 	Printfln("Pointer:", &name)

	// сканирование строк
	var namez string
    var category string
    var price float64

    fmt.Print("Enter text to scan: ")

    n, err := fmt.Scan(&namez, &category, &price) // вот здесь идет считывание данных из консоли
    if (err == nil) {
        Printfln("Scanned %v values", n)
        Printfln("Name: %v, Category: %v, Price: %.2f", namez, category, price)
    } else {
        Printfln("Error: %v", err.Error())
    }
}
