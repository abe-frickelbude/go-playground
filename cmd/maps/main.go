package main

import (
	"fmt"
	"strings"
)

func main() {

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	fmt.Println(wordCount("a X quick X brown X fox jumped over a lazy dog"))
}

func wordCount(source string) map[string]int {

	counts := make(map[string]int)
	tokens := strings.Fields(source)
	for _, token := range tokens {
		counts[token]++
	}
	return counts
}
