package main

import "fmt"

func fib(a, b chan int) {
	x, y := 0, 1

	for {
		select {
		case a <- x: // пишет в канал, здесь блочится до тех пор, пока кто-нить не прочитает!
			x, y = y, x + y 
		case <-b: // здесь сработает сразу же, как только из канала придет какая то инфа! любой int
			return
		}
	}
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			printfln("Число фибоначи %d", <-a)
		}
		b <- 100
		close(b)
	}()

	fib(a, b)

}

func printfln(template string, verb ...interface{}) {
	fmt.Printf(template + "\n", verb...)
}
