package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FindShort(s string) int {
	arr := strings.Split(s, " ")

	var min int = len(arr[0])

	for i := 1; i < len(arr); i++ {
		if min > len(arr[i]) {
			min = len(arr[i])
		}

	}
	return min
}

func FindShort2(s string) (c int) {
	for _, word := range strings.Fields(s) {
		if c == 0 || len(word) < c {
			c = len(word)
		}
	}
	return c
}

// Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
// Return your answer as a number.
func SumMix(arr []any) (sum int) {
	for _, v := range arr {
		if strVal, ok := v.(string); ok {
			if val, err := strconv.Atoi(strVal); err == nil {
				sum += val
			}
		} else {
			sum += v.(int)
		}
	}
	return
}

func SumMix2(arr []any) int {
  
	sum := 0
	
	for _, val := range arr {
	  
	 switch val := val.(type){
	   case int:
		sum += val
	   case string:
		k,_ := strconv.Atoi(val)
		sum += k
	 }
	 
	}
	
	return sum
  }

  func SumMix3(arr []any) int {
	sum := 0
	for _, v := range arr {
		iv, _ := strconv.Atoi(fmt.Sprintf("%v", v))
		sum += iv
	}
	return sum
  }


  func Points(games []string) (result int) {
	for _, v := range games {
		xy := strings.Split(v, ":")

		if xy[0] > xy[1] {
			result += 3
		} else if xy[0] == xy[1] {
			result += 1
		}
	}
	return
  }

  // clever solution 
  func Points2(games []string) (score int) {
	var x, y int
	for _, g := range games {
		fmt.Sscanf(g, "%d:%d", &x, &y)
		if x > y {
			score += 3
		}
		if x == y {
			score += 1
		}
	}

	return
}

func main() {

	gameResults := []string{"1:0","2:0","3:0","4:0","2:1","3:1","4:1","3:2","4:2","4:3"}
	fmt.Printf("Сумма очков во всех играх = %+v\n", Points(gameResults))


	fmt.Printf("\n Сумма чисел = %d \n\n", SumMix([]any{9, 3, "7", "3"}))


	str := "This is test string"
	fmt.Printf("Самое короткое слово это - %d", FindShort(str))

	// 	strconv.Itoa()
	for i := 0; i < 2; i++ {
		fmt.Printf("\n %v в 16-ой системе счисления = %s", i, strconv.FormatInt(int64(i), 16))
		fmt.Printf("\n %v в 16-ой системе счисления = %d", i, strconv.Itoa(i))
	}
}
