package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(5))
}


func fib(n int) int {
	if n < 2 {
		return n
	}
	arr:=make([]int,n+1)
	arr[0]=0
	arr[1]=1
	for i := 2; i <=n; i++ {
		arr[i]=(arr[i-1]+arr[i-2])%1000000007
	}
	return arr[n]
}