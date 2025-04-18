package main

import "fmt"

func fib(a, exit chan int) {
	x, y := 0, 1

	for {
		select {
		case a <- x:
			{
				x = y - x
				y = y + x
			}
		case <-exit:
			return
		}
	}
}

func main() {
	a, exit := make(chan int), make(chan int)

	go func ()  {
		for i := 0; i < 10; i++ {
			fmt.Printf("Число фибоначи %v \n", <- a)
		}
		close(exit)
	}()

	fib(a, exit)

}
