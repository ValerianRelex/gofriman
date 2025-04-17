package main

import (
	"io"
	"strings"
)

func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}
func main() {
	r := strings.NewReader("Kayak")
	processData(r)
}