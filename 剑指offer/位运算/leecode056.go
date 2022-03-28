package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int {1,22,22,3,1}))
}

// 可以找单数
func singleNumber(nums []int) int {
	res:=0
	for _, num := range nums {
		res^=num
	}
	return res
}