package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(printNumbers(1))
}

func printNumbers(n int) []int {
	m:=int(math.Pow10(n))
	var ints []int
	for i := 1; i < m; i++ {
		ints = append(ints, i)
	}
	return ints
}