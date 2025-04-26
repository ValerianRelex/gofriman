package main

import "fmt"

func PlayList(player playerInterface, songs []string) {
	for _, s := range songs {
		player.Play(s)
	}
	player.Stop()
}

type MyType int

// func (t MyType) String() string {
// 	return fmt.Sprintf("%v", t)
// }

// Исправленный метод String()
func (t MyType) String() string {
	// Явно преобразуем к базовому типу int
	return fmt.Sprintf("MyType(%d)", int(t))
}

func main() {
	songsList := []string{"Modern tAlking", "Jessica Jay", "Rastaman"}

	player := TapePlayer{}
	PlayList(player, songsList)

	sliceMyType := []MyType{1, 5, 6, 3, -5}

	fmt.Println(sliceMyType[2])
}
