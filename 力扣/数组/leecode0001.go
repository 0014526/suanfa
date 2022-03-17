package main

func main()  {
	twoSum2([]int{2,7,11,15}, 9)
}

// 爆破
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i]+nums[j]==target {
				return []int{i,j}
			}
		}
	}
	return []int{0,0}
}


// 使用一个字典将数据存入
func twoSum2(nums []int, target int) []int {
	hashTable:=make(map[int]int)
	for i, num := range nums {
		if j,ok:=hashTable[target-num];ok {
			return []int{i,j}
		}
		hashTable[num]=i
	}
	return nil
}