package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// insert next fib() value to channel, will block until Println() below retrieves a value
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	// zero-buffer-length channels are effectively "exchange" points - producer(s) will block until
	// a channel's value is retrieved, consumer(s) will block until a value becomes available
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
