package main

import (
	"fmt"

	"golang.org/x/tour/pic"

	"strings"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := 0; i < dy; i++ {

		pic[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			pic[i][j] = uint8(dx)
		}
	}
	return pic
}

// задание из "тур по гоу" - https://go.dev/tour/moretypes/23
func WordCount(s string) map[string]int {	
	// создадим мапу
	result := make(map[string]int)

	split := strings.Split(s, " ")


	for _, value := range split {
		// result[string(value)] += 1
		result[value]++
	}
	
	return result
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type MyFloat float64

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {


	fmt.Stringer



	lst := []string{"a", "b", "c", "d"}
	for k, v := range lst {
		if k == 0 {
			lst = []string{"aa", "bb", "cc", "dd"}
		}
		fmt.Println(v)
	}

	var fuu MyFloat = 12

	fmt.Println(fuu)

	pic.Show(Pic)


	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
