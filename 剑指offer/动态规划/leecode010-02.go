package main

import (
	"fmt"
)

func main() {
	fmt.Println(rectCover( 4 ))
}
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
// 斐波那契数列变种
func rectCover( number int ) int {
	// write code here
	if number<2 {
		return number
	}
	res:=make([]int,number+1)
	res[0]=0
	res[1]=1
	res[2]=2
	for i := 3; i <=number; i++ {
		res[i]=(res[i-1]+res[i-2])%1000000007
	}
	return res[number]

}