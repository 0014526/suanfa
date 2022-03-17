package main

import (
	"sort"
)

/**
找出数组中重复的数字。
在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
示例 1：
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
 
限制：
2 <= n <= 100000
 */
func main() {

}

// 将数据存入map中，如果map里面含有这个元素，就把这个元素返回
func findRepeatNumber1(nums []int) int {
	m:=make(map[int]int)
	for _, value := range nums {
		if _, ok := m[value]; ok {
			// 存在该值，输出该值
			return value
		}else{
			m[value]=1
		}
	}
	return -1
}

// 2.对切片进行排序，与下一个元素重复就返回
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)

	// 遍历切片
	for i,numSize := 0,len(nums); i <numSize; i++ {
		if nums[i]==nums[i+1] {
			return nums[i]
		}
	}
	return -1
}