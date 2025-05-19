package main

import "fmt"

func PrinterLn(template string, val ...any) {
	fmt.Printf(template + "\n", val...)
}