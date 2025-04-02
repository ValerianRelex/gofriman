package main

import "fmt"

func main() {
	lst := []string{"a", "b", "c", "d"}
	for k, v := range lst {
		if (k ==0) {
			lst = []string{"aa", "bb", "cc", "dd"}
		}
		fmt.Println(v)
	}
}