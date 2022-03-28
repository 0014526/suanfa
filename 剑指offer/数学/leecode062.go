package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param n int整型
 * @param m int整型
 * @return int整型
 */
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (m+lastRemaining(n-1,m))%n
}

func lastRemaining1(n int, m int) int {
	temp:=0
	for i:=2;i<n+1;i++{
		temp=(temp+m)%i
	}
	return temp
}