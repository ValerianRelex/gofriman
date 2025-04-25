package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type intSlice []int

func (islice intSlice) String() string {
	stringSlice := make([]string, len(islice))

	for i, v := range islice {
		stringSlice[i] = fmt.Sprintf("%#v", v)
		fmt.Printf(" val = %v, type = %T, data = %d \n", v, v, v)
	}

	return strings.Join(stringSlice, ",")
}

func main() {

	notSortedSlice := []int{-1, 5, 3, 10, 9, 0, -5}

	// var sortedSlice intSlice
	sortedSlice := make(intSlice, len(notSortedSlice))

	copy(sortedSlice, notSortedSlice)
	sort.Ints(sortedSlice)

	fmt.Println("\n\n Отсортированный массив через запятую, композиция в действии? \n", sortedSlice)

	i := 100
	var j *int = &i

	fmt.Println("i = ", i, "j = ", j)
	i++
	fmt.Println("i = ", i, "j = ", j)

	first := 10
	second := &first
	first++
	fmt.Println("first = ", first, "second = ", *second)

	var nilPointer *int
	fmt.Println(nilPointer == nil) // true
	pointerNilPointer := &nilPointer
	fmt.Println("nilPointer = ", nilPointer, "pointerNilPointer = ", pointerNilPointer)

	nilPointer = &first
	fmt.Println("nilPointer = ", *nilPointer, "*pointerNilPointer указывает на nilPointer, а там хранится адрес! поэтому =  ", *pointerNilPointer)

	array := [3]string{"bob", "anna", "ceren"}

	secondValue := &array[1]

	fmt.Println(*secondValue) // anna
	sort.Strings(array[:])
	fmt.Println(*secondValue) // bob

	// constExample()
	// loopPrintRandom()
}

func loopPrintRandom() {
	for i := 0; i < 990056053552318569; i++ {
		fmt.Print(rand.Int())
	}
}

func constExample() {
	const price, tax, isdirty = 125.5, 10, true
	const min, max, ishidden float32 = 125.5, 10, 0

	fmt.Println(price*tax, "isdirty = ", isdirty == false)
	fmt.Println(min*max, ishidden+14)
}
