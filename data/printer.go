package main

import "fmt"

func Printfln(str string, vals ...any) {
	fmt.Printf(str + "\n", vals...)
}