package main

import (
	"math"
	"sort"
)

func main() {
	findMedianSortedArrays2([]int{}, []int{1})
}


// 基础解法,用轮子
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2==0 {
		return (float64(nums1[len(nums1)/2-1] )+ float64(nums1[len(nums1)/2])) / 2
	}else{
		return float64(nums1[len(nums1)/2])
	}
}

// 使用快慢指针做
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 快慢指针
	length:=len(nums1)+len(nums2)
	var left,right,index int
	arr:=make([]int,length)
	if len(nums1)==0 {
		arr=nums2
	}
	if len(nums2) == 0 {
		arr=nums1
	}
	for index < length&&len(nums1)!=0&&len(nums2)!=0 {
		if left>=len(nums1) {
			left--
			nums1[left]=math.MaxInt
		}
		if right>=len(nums2) {
			right--
			nums2[right]=math.MaxInt
		}
		if nums1[left]<nums2[right] {
			arr[index]=nums1[left]
			if left<len(nums1) {
				left++
			}
		}else if nums1[left]>=nums2[right]{
			arr[index]=nums2[right]
			if right<len(nums2) {
				right++
			}
		}
		index++
	}
	if len(arr)%2==0 {
		return (float64(arr[len(arr)/2-1] )+ float64(arr[len(arr)/2])) / 2
	}else{
		return float64(arr[len(arr)/2])
	}

}