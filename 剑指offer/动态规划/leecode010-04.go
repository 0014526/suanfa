package main

func main() {

}


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloorII1( number int ) int {
	// write code here
	if number==1 {
		return 1
	}
	a,b:=1,2
	for i := 0; i < number-2; i++ {
		b=a+b+1
		a=b-1
	}
	return b
}

func jumpFloorII2( number int ) int {
	// write code here
	return 1<<(number-1)
}