package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var intMaxVal = math.MaxInt64
	var floatMaxVal = math.MaxFloat64

	fmt.Println(intMaxVal)
	fmt.Println(intMaxVal * 2)
	fmt.Println(floatMaxVal)
	fmt.Println(floatMaxVal * 2)

	fmt.Println("\n\n", math.IsInf((floatMaxVal*2), 0), "\n\n")

	posResult := 3 % 2
	negResult := -3 % 2
	absResult := math.Abs(float64(negResult))
	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)

	first := 100
	second := &first
	third := &first
	alpha := 100
	beta := &alpha
	fmt.Println("second == third", second == third) // true
	fmt.Println("second == beta", second == beta) // разные участки памяти  false

	fmt.Println("first == *beta ---> ", first == *beta) // 100 == 100  true

	kayak := 275        // литеральные значения приводят к типу int
	soccerBall := 19.50 // и к типу float64

	// если не сделать явного преобразования типов, то получим ошибку компиляции
	total := float64(kayak) + soccerBall // # operations .\main.go:41:11: invalid operation: kayak + soccerBall (mismatched types int and float64)
	// total := kayak + int(soccerBall)
	fmt.Println(total)
	fmt.Println(int8(total)) // происходит переполнение и неверный вывод

	fmt.Println("math.Ceil(21.1)", math.Ceil(21.1))   // округление чисел
	fmt.Println("math.Floor(21.1)", math.Floor(21.1)) // округление чисел

	val1 := "true"
	val2 := "false"
	val3 := "not true"
	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	bool3, b3err := strconv.ParseBool(val3)
	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
	fmt.Println("Bool 3", bool3, b3err)

	val := "0"
	valBool1, berr := strconv.ParseBool(val)
	if berr == nil {
		fmt.Println("Parsed value:", valBool1)
	} else {
		fmt.Println("Cannot parse", val)
	}

	// запись ниже более компактная и лаконичная, аналог записи выше
	val = "1"
	if bool1, b1err := strconv.ParseBool(val); b1err == nil {
		fmt.Println("Parsed value:", bool1)
	} else {
		fmt.Println("Cannot parse", val)
	}

	str1 := "100"

	fmt.Println("int1 = ", str1)

	if int1, strErr := strconv.ParseInt(str1, 0, 8); strErr != nil {
		fmt.Println("int1 не удалось распарсить в строку -> ", int1)
	} else {
		fmt.Println("int1 распарсили = ", int1+5) // теперь это не строка. а int и поэтому можно производить мат. операции
	}

	val11 := "100"
	int11, int1err := strconv.ParseInt(val11, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value:", int11)
	} else {
		fmt.Println("Cannot parse", val11, int1err)
	}

	val100 := "100"
	int100, int100err := strconv.ParseInt(val100, 10, 0)
	if int100err == nil {
		var intResult int = int(int100)
		fmt.Println("Parsed value:", intResult+100)
	} else {
		fmt.Println("Cannot parse", val100, int100err)
	}

	// способы задания переменной
	// var valBool bool = true
	// var valBool = true
	// var valBool bool // но оно будет false
	valBool := true

	strBool := strconv.FormatBool(valBool)
	fmt.Println("valBool =", strBool)
	fmt.Println("valBool = " + strBool)

	stringa := "redoctober"

	for index, character := range stringa {
		// выводит последовательность значений rune
		fmt.Println("Index:", index, "char:", string(character)) // если убрать приведение типа, выведется кодировка символов!
	}

	stringAsRuna := "I love GO!"

	for index, char := range stringAsRuna {
		switch char {
		case 'l': fmt.Println("Index of char 'l'", index)
		case 'o': fmt.Println("Index of char 'o'", index)
		}
	}

	emoji := "привет😀 \U0001F600"

	for _, e := range emoji {
		fmt.Println(e, string(e))
	}
}
