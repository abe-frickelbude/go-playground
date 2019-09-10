package main

import (
	"fmt"
	"math"
	"time"
)

const (
	// PI stuff
	PI      float32 = 3.14159
	EPSILON         = 1.19e-07
)

func main() {

	defer fmt.Println("done!")

	fmt.Println("Hello, 世界")
	fmt.Println("The time is", time.Now())
	// for(1..10)
	//	fmt.Println(rand.Intn(100))
	x, y := test(34, 45)
	fmt.Println(x, y)

	const z complex64 = 2 + 1i
	fmt.Println(z)

	fmt.Printf("%T %v\n", z, z)
	fmt.Printf("%T %v\n", PI, PI)

	f := float64(x)
	fmt.Println(f)

	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i * i)
	// }

	fmt.Println(pow(2, 3, 9))
	fmt.Println(pow(2, 3, 7))

	fmt.Println("own: ", sqrt(2.0))
	fmt.Println("reference: ", math.Sqrt(2.0))

	xptr := &x
	fmt.Printf("%T %v\n\n", xptr, xptr)
	fmt.Println(*xptr)

	switchStuff()
	slices()

	compSignal := make(chan int)
	for i := 0; i < 10; i++ {
		announce("Timeout reached", 1e9, compSignal)
		<-compSignal // wait for announce() to finish,  discard received value
	}
}

func announce(message string, delay time.Duration, c chan int) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
		c <- 1
	}() // Note the parentheses - must call the function.
}

func test(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return a, b
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sqrt(x float64) float64 {
	z := x
	zPrev := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-zPrev) <= EPSILON {
			break
		}
		zPrev = z
	}
	return z
}

func switchStuff() {
	t := time.Now()
	// switch with no condition can be used to replace a typical ugly if-then-else-if-.... chain
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func slices() {

	a := make([]int, 5)
	for i := range a {
		a[i] = i
	}

	printSlice("a", a)

	b := make([]int, 3, 5)
	for i := range b {
		b[i] = i
	}

	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}


