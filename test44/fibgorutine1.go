package main

import (
	"fmt"
	"time"
)

func fib(ch, quit chan int) {
	var first int = 0
	var next int = 1

	for {
		select {
		case ch <- first:
			{
				first = next - first
				next = first + next
				time.Sleep(time.Millisecond)
			}
		case <-quit:
			return
		default: fmt.Println("fdsgdfgsdfg")
		}

	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch) // ждем данные!!! из канала... висим тут, пока они не придут! т.е. здесь аноним горутина заблочена!
		}
		quit <- 100
	}()

	fib(ch, quit)

}
