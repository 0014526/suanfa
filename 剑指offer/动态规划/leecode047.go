package main

func maxValue(grid [][]int) int {
	dp:=make([][]int,len(grid)+1)
	for i := 0; i < len(grid); i++ {
		dp[i]=make([]int,len(grid[i])+1)
	}
	max:= func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j]=dp[i-1][j-1]+max(dp[i-1][j],dp[i][j-1])
		}
	}
	return dp[len(grid)][len(grid[0])]
}
