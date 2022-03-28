package main

import (
	"fmt"
)

func main() {
	fmt.Println(numWays(3))
}

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	left,right:=1,1
	for i := 2; i <= n; i++ {
		sum:=left+right
		left=right
		right=sum
	}
	return right
}
