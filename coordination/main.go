package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}

var rwMutex = sync.RWMutex{}

func doSum(count int, val *int) {
	mutex.Lock()
	for i := 0; i < count; i++ {
		*val++
	}
	mutex.Unlock()
	wg.Done()
}

var squares = map[int]int{}

func calculateSquares(iterations, max int) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max) // генерируем случайное число в пределах max

		rwMutex.RLock()
		square, ok := squares[val] // проверим, существует в мапе такое значение по ключу val?
		rwMutex.RUnlock()

		if ok {
			Printfln("Cashed value %v = %v", val, square) // выведем ключ и его значение (квадрат)
		} else {
			rwMutex.Lock()
			/*
							Между освобождением блокировки чтения и
				 получением блокировки записи может быть задержка, в течение
				 которой другие горутины могут получить блокировку записи и внести
				 изменения, поэтому важно убедиться, что состояние данных не
				 изменилось после блокировки записи
			*/
			if _, ok := squares[val]; !ok { // проверим, не имзенилось ли данное, пока брали блокировку
				squares[val] = int(math.Pow(float64(val), 2)) // возведем в степень 2
				Printfln("Added value: %v = %v", val, squares[val])
			} else {
				fmt.Println("-------таки был другая запись-----------------------------------------")
			}
			rwMutex.Unlock()
		}
	}
	wg.Done()
}

func main() {
	counter := 0
	numWG := 5

	wg.Add(numWG)
	for i := 1; i <= numWG; i++ {
		go doSum(5000, &counter)
	}

	wg.Wait()
	Printfln("Total: %v", counter)

	// усложним задачу и поиграем с RWMutex

	wg.Add(numWG)
	for i := 0; i < numWG; i++ {
		go calculateSquares(5, 10)
	}
	wg.Wait()

	Printfln("Cached values: %v", len(squares))
}
