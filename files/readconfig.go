package main

import "os"

func LoadConfig() (err error) {
	data, err := os.ReadFile("config.json")
	if err == nil {
		Printfln(string(data))
	}
	return
}
func init() {
	err := LoadConfig()
	if err != nil {
		Printfln("Error Loading Config: %v", err.Error())
	}
}