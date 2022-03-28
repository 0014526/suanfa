package main

func main() {

}


func twoSum1(nums []int, target int) []int {
	left:=0
	right:=len(nums)-1

	for left<right {
		if (nums[left]+nums[right])==target {
			return []int{nums[left],nums[right]}
		}
		if (target-nums[left])<=nums[right] {
			right--
		}else{
			left++
		}
		if left>(len(nums)-1)||right<0 {
			break
		}
	}
	return nil
}

// 使用hash处理
func twoSum2(nums []int, target int) []int {
	m:=make(map[int]struct{})
	for _,v:=range nums{
		if _, ok := m[target-v]; ok {
			return []int{target-v,v}
		}else{
			m[v]= struct{}{}
		}
	}
	return nil
}