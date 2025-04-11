package main

import (
	"fmt"
	//"strings"
	"regexp"
)

func main() {
	description := "A boat for one person"
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}
}