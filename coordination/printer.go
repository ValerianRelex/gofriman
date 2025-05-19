package main

import "fmt"

func Printfln(templates string, values ...any) {
	fmt.Printf(templates + "\n", values...)
}