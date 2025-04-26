package main

import (
	"fmt"
)

func fibd(ch1 chan int, qch <-chan int) {
	x1, x2 := 0, 1

	for {
		select {
		case ch1 <- x1:
			x1, x2 = x2, x2+x1
		case <-qch:
			return
		}
	}
}

func main() {
	chan1 := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println(<-chan1)
		}
		quit <- 100500
		close(chan1)
		close(quit)
	}()

	fibd(chan1, quit)
}
