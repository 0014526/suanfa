package main

import (
	"fmt"
)

func main() {
	fmt.Println(hammingWeight1(11))
}

// 普通解法，对于负数存在问题
func hammingWeight1(num uint32) int {
	if num == 0 {
		return 0
	}
	res:=0
	for num!=0 {
		if num&1==1 {
			res++
		}
		num>>=1
	}
	return res
}

/**
现考虑二进制数：val :1101000, val-1: 1100111 那么val & （val-1） : 1100000
如果你会了这个操作，是不是这题就很简单了。
 */
func hammingWeight2(num uint32) int {
	res:=0
	for num!=0 {
		res++
		num&=num-1
	}
	return res
}