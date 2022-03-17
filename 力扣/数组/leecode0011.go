package main

import (
	"math"
)

func main() {
	maxArea1([]int{1,8,6,2,5,4,8,3,7})
}

// 重点是在转换方程(左右指针)
func maxArea1(height []int) int {
	right:=len(height)-1
	left:=0
	var max float64
	for left<right {
		max=math.Max((float64(right)-float64(left))*math.Min(float64(height[left]),float64(height[right])),float64(max))
		if height[left]<height[right] {
			left++
		}else{
			right--
		}
	}
	return int(max)
}