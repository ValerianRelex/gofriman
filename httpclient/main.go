package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	// Printfln("Starting HTTP Server")
	// http.ListenAndServe(":5000", nil)

	// Сервак HTTP
	go http.ListenAndServe(":5000", nil)

	time.Sleep(time.Second)

	// клиент HTTP
	response, err := http.Get("http://localhost:5000/echo") // варианты: /json  и /html
	if err == nil {
		response.Write(os.Stdout)
	} else {
		Printfln("Error: %v", err.Error())
	}

	// левый сервер
	response, err = http.Get("https://search.worldbank.org/api/v3/wds?format=json&qterm=energy&display_title=water&fl=display_title&rows=2&os=201")
	if err == nil {
		response.Write(os.Stdout)
		Printfln("\nStatus code = %v", response.StatusCode)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
