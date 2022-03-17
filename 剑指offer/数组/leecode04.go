package main

import (
	"sort"
)

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

 

示例:

现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
 
限制：
0 <= n <= 1000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func main() {

}


// 爆破，笨蛋法
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j< len(matrix[0]); j++ {
			if matrix[i][j]==target{
				return true
			}
		}
	}
	return false
}

// 找规律
// 直接取左下或者右上的元素
func findNumberIn2DArray2(matrix [][]int, target int) bool  {
	// 以左下角为原点
	i:=len(matrix)-1
	j:=0
	for i > -1 && j<len(matrix[0]){
		if matrix[i][j]>target {
			i--
		}else if matrix[i][j]<target {
			j++
		}else{
			return true
		}
	}
	return false
}

// 使用特殊方法,数组切片
func findNumberIn2DArray3(matrix [][]int, target int) bool  {
	// 以左下角为原点
	for _,value:=range matrix{
		i:=sort.SearchInts(value,target)
		if i<len(value)&&target==value[i] {
			return true
		}
	}
	return false
}