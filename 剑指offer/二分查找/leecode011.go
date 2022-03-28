package main

import (
	"fmt"
)

func main() {
	fmt.Println(minArray1([]int{2,2,2,0,1}))
}

// 这有啥意义吗
func minArray(numbers []int) int {
	min:=numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i]<min {
			min=numbers[i]
		}
	}
	return min
}

// 二分法
func minArray1(numbers []int) int {
	left:=0
	right:=len(numbers)-1
	for left<right {
		mid:=left+(right-left)/2
		if numbers[mid]<numbers[right] {
			right=mid
		}else if numbers[mid]>numbers[right] {
			left=mid+1
		}else{
			right--
		}
	}
	return numbers[left]
}
