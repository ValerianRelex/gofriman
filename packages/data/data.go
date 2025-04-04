package data

import "fmt"

func init() {
	fmt.Println("i am from data.go")
}

// пример того, что при простом импорте пакета, без использования его функций - сама ф. инициализации не будет запущена!
func GetData() []string {
    return []string {"Kayak", "Lifejacket", "Paddle", "SoccerBall"}
 }