package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// функция задания минимального и максимального диапазона для случайного числа
func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {
	Printfln("Hello, Math and Sorting")

	val1 := 279.00
	val2 := 48.95
	Printfln("%v -> Abs: %v", val1, math.Abs(val1))
	Printfln("%v -> Ceil: %v", val2, math.Ceil(val2))
	Printfln("%v -> Copysign: %v", val1, math.Copysign(val1, -5))
	Printfln("%v -> Floor: %v", val2, math.Floor(val2))
	Printfln("%v, %v -> Max: %v", val1, val2, math.Max(val1, val2))
	Printfln("%v, %v -> Min: %v", val1, val2, math.Min(val1, val2))
	Printfln("%v, %v -> Mod: %v", val1, val2, math.Mod(val1, val2))
	Printfln("%v -> Pow: %v", val1, math.Pow(val1, 2))
	Printfln("%v -> Round: %v", val2, math.Round(val2))
	Printfln("%v -> RoundToEven: %v", val2, math.RoundToEven(val2))

	// генерация случайного значения
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, IntRange(10, 20))
	}

	// перетасовка, как тасовать - задается пользователем
	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}

	// сортировка
	ints := []int{9, 4, 2, -1, 10}
	Printfln("\n\nInts: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)

	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)

	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}

	notSorted := []int{0, 5, -5, 2, -3, 4, -4, 10, 11, 9, -9, -11, 14, 13, 15}
	sliceSorted := make([]int, len(notSorted))
	copy(sliceSorted, notSorted)

	sort.Ints(sliceSorted)
	fmt.Printf("\nнесортир срез = %v отсортированный срез =%v", notSorted, sliceSorted)

	// пользовательская сортировка, применили композицию типа
	fmt.Println()
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlices(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}
