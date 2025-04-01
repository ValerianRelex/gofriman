package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, Collections")

	var array1 [3]string
	array2 := [3]string{}

	array1[0] = "java"
	array1[1] = "go"
	array1[2] = "rust"

	fmt.Println(array1)
	fmt.Println(array2)
	//

	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	var otherArray [3]string = names

	notArray := &names
	names[0] = "RedOctober"
	fmt.Println(names)
	fmt.Println(otherArray)
	fmt.Println("notArray", *notArray)

	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}

	names[0] = "Kayak"
	same := names == moreNames
	fmt.Println("arrays is identity", same)

	// срезы
	fmt.Println("... slice")

	srez := make([]string, 3)
	fmt.Println(srez)

	slice1 := make([]string, 3, 6)
	slice1[0] = "Oklahoma"
	slice1[1] = "Nevada"
	slice1[2] = "Gorgia"

	slice2 := append(slice1, "omen", "perdel", "murka") // из за того, что емкость позволяет вместить еще элементы, ссылка будет на один и тот же массив
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "Кутузов!" // изменив начение в slice1 поменяется значение и в slice2
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice4 := make([]string, 2, 6)
	slice4[0] = "aibaba"
	slice4[1] = "kongo"

	sliceMoreName := []string{"love", "me", "tender"}
	enotherSliceMoreNames := append(slice4, sliceMoreName...)
	fmt.Println(enotherSliceMoreNames)

	enotherSliceMoreNames[0] = "fuck"
	fmt.Println(enotherSliceMoreNames)
	fmt.Println(slice4)
	fmt.Println(sliceMoreName)

	newSlice := array1[1:2] // создать новый срез из массива, начиная с индекса 1 и размером 2 - 1 = 1 элемент
	newSlice2 := array1[:]
	fmt.Println("array1", array1)
	fmt.Println("newSlice", newSlice)
	fmt.Println("newSlice2", newSlice2, "размер =", len(newSlice2), "вместимость =", cap(newSlice2))

	array1[1] = "changed!!!" // и вот чудо GO, я то ожидал другое. Но, два новых массива поддерживаются одним и тем же массивом!!! array1
	fmt.Println("array1", array1)
	fmt.Println("newSlice", newSlice)
	fmt.Println("newSlice2", newSlice2, "длина =", len(newSlice2), "вместимость =", cap(newSlice2))
	fmt.Println("newSlice", newSlice, "длина =", len(newSlice), "вместимость =", cap(newSlice))
	//

	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"} // этот массив поддерживает оба среза заданных ниже
	someNames := products[1:3]
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))

	fmt.Println("products", products)
	fmt.Println("products len", len(products), "cap:", cap(products))

	someNames = append(someNames, "anotherArray") // добавив новый элемент, превысилил размер поддерживающего массива
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames)) // создался новый массив, обьемом 6 элементов
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames)) // прежний массив поддерживает allNames
	//
	//  копирование срезов

	fmt.Println("копирование срезов")
	products2 := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames2 := products2[1:]      // "Lifejacket", "Paddle", "Hat"
	someNames2 := make([]string, 2) //
	copy(someNames2, allNames2)
	fmt.Println("someNames2:", someNames2)
	fmt.Println("allNames2", allNames2)

	apples := []string{"green", "red", "white", "blue", "yellow"}
	basket := []string{"big", "small"}

	copy(basket, apples)
	fmt.Println("basket", basket)
	fmt.Println("apples", apples)

	// операция удаления с помощьью функции append
	deleted := append(apples[:2], apples[3:]...)
	fmt.Println(deleted)

	//  Перечисление срезов
	for i, val := range deleted[2:] {
		fmt.Println("index", i, "value", val)
	}

	// срезы между собой сравнить нельзя, только с nil
	// invalid operation: deleted == apples (slice can only be compared to nil) (exit status 1)
	// var b = deleted == apples;
	// fmt.Println(b)

	// получение массива, лежащего в основе среза
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[3]string)(p1) // явное преобразование типа среза, указали еще и длину, меньше может быть, а больше - ошибка
	array := *arrayPtr           // следуем за указателем, чтобы получить значение массива
	fmt.Println(array, "емоксть массива, или это срез? =========", cap(array), len(array))

	// получение массива, лежащего в основе среза. Второй способ, по мнение deepseek невероятный
	p2 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtrUnreal := ([3]string)(p2) // явное преобразование типа среза, указали еще и длину, меньше может быть, а больше - ошибка
	arrayUnreal := arrayPtrUnreal     // следуем за указателем, чтобы получить значение массива
	fmt.Println(arrayUnreal, "емоксть массива, или это срез? =========", cap(arrayUnreal), len(arrayUnreal))

	// TODO: НЕПОНЯТНОЕ ПОВЕДЕНИЕ, РАЗОБРАТЬСЯ!!!
	fmt.Println("\n\n")
	// pInt := 100
	nuar := &p1
	fmt.Println(nuar)   // печать с префиксом & --- а я ожидаю область памяти где начинается срез
	fmt.Println(*nuar)  // а тут верное ожидание, я же следую за указателем
	fmt.Println(&p1[0]) // адрес начала поддерживающего массива, срез структура, содержит

	// КАРТЫ!!!
	//

	productes := make(map[string]float64, 10)
	productes["Kayak"] = 279
	productes["Lifejacket"] = 48.95

	fmt.Println("Map size:", len(productes))
	fmt.Println("Price:", productes["Kayak"])
	fmt.Println("Price:", productes["Lifejacket"])

	productes["Kayak"] = 0
	fmt.Println("new Price:", productes["Kayak"])
	fmt.Println("Map size:", len(productes))
	fmt.Println("productes =", productes)

	// определение мапы с помощьью литералов
	products44 := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}

	fmt.Println("products44 =", products44)

	// Решаем проблему, вовзращения нуля при отсутствующем ключе в map
	pro := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0, // чтобы было понятнее, просто закомить эту строчку
	}
	value, ok := pro["Hat"]
	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	// упростим с помощью оператора инициализации
	if value, ok := pro["figvam"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	// удаление обьектов с карты

	delete(pro, "Hat")

	// перечислим карту

	for key, value := range productes {
		fmt.Println("key =", key, "value =", value)
	}

	// перечислить по порядку мапу надо так, советует автор книги. Сначало получить срез, потом взять ключи у среза и по ключам вывести мапу

	mapa := map[string]int{
		"b": 5,
		"v": 3,
		"l": 7,
		"a": 2,
		"c": 1,
	}

	// выведет че попало!
	for k, v := range mapa {
		fmt.Println("k =", k, "v =", v)
	}

	// создаем срез
	srezzz := make([]string, 0, len(mapa))

	// заполним срез значениями из мапы
	for key, _ := range mapa {
		srezzz = append(srezzz, key)
	}

	// срез содержит ключи мапы, но в хаотичном порядке, поэтому мы это дело отсортируем
	sort.Strings(srezzz)

	// vualia! здесь значения среза и есть ключи мапы и таким образом получили отсортированный по ключам вывод мапы
	for _, v := range srezzz {
		fmt.Println("key =", v, "value =", mapa[v])
	}

	// понимание двойной природы строк

	var price string = "$48.95"
	var currency byte = price[0]
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency тип byte:", currency)
	// явное преобразование 
	fmt.Println("Currency преобразованные :", string(currency))

	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
	fmt.Println("amountString", amountString)
	
}
