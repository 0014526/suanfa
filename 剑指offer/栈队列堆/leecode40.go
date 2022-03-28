package main

import (
	"sort"
)

func main() {

}


func getLeastNumbers1(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k+1]
}

func getLeastNumbers2(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k+1]
}