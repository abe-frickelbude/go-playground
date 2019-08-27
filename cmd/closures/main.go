package main

import (
	"fmt"
)

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func fibonacci() func() int {

	prevPrevValue := 0 // F(n-2)
	prevValue := 1     // F(n-1)
	return func() int {
		value := prevValue + prevPrevValue
		// update F(n-2) and F(n-1)
		prevPrevValue = prevValue
		prevValue = value
		return value
	}
}
