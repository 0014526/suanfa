package main

func main() {
	m:=[][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	spiralOrder(m)
}


func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0 {
		return nil
	}
	step:=0
	size:=len(matrix)*len(matrix[0])
	// 定义四个方向端点
	top,bottom,right,left:=0,len(matrix)-1,len(matrix[0])-1,0
	nums:=make([]int,size)
	for step<size {
		// 从左上到右上
		for i := left; i <=right&&step<size; i++ {
			nums[step]=matrix[top][i]
			step++
		}
		top++
		//  从右上到右下
		for i:=top;i<=bottom&&step<size;i++{
			nums[step]=matrix[i][right]
			step++
		}
		right--
		// 从右下到左下
		for i := right; i >=left&&step<size; i-- {
			nums[step]=matrix[bottom][i]
			step++
		}
		bottom--
		for i := bottom; i >=top&&step<size; i-- {
			nums[step]=matrix[i][left]
			step++
		}
		left++
	}
	return nums
}