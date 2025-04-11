package main

import (
	"fmt"
	//"strings"
	"regexp"
)

func main() {
	description := "A boat for one person"

	// матчим регулярные выражения, иными словами - ищем совпадения
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	// вариант другой
	pattern, compileErr := regexp.Compile("[A-z]oat")

	question := "Is that a goat?"
	preference := "I like oats" // !
	if compileErr == nil {
		fmt.Println("Description:",
			pattern.MatchString(description))
		fmt.Println("Question:",
			pattern.MatchString(question))
		fmt.Println("Preference:",
			pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}

	// совпадение по статичной части и изменяющейся
	pattern2 := regexp.MustCompile("A [A-z]* for [A-z]* person")

	str := pattern2.FindString(description)
	fmt.Println("Match:", str)
}
