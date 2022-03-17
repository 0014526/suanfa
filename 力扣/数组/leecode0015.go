package main

import (
	"fmt"
	"sort"
)

func main() {
	threeSum([]int{-1,0,1,0})
}

// 可以爆破但是不提倡
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	left:=0
	right:=len(nums)-1
	reMap:=make(map[int]int)
	for i, num := range nums {
		reMap[num]=i
	}
	result:=make([][]int,0)
	for left<right {

		target:=0-(nums[left]+nums[right])
		for i := left+1; i <right; i++ {
			if nums[i]==target {
				fmt.Println([]int{nums[left],nums[right],nums[i]})
				for nums[left]==nums[left+1]&&(left+1)<right {
					left++
				}
				for nums[right]==nums[right-1]&&left<(right-1) {
					right--
				}
				result = append(result, []int{nums[left],nums[right],nums[i]})
				if right==left+1 {
					break
				}
			}
		}

		if (nums[left]+nums[right])<=0 {
			left++
		}else if (nums[left]+nums[right])>0{
			right--
		}
	}
	return result
}