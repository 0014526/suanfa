package main

func add(a int, b int) int {
	temp:=0
	for b != 0 {
		temp=(a&b)<<1
		a^=b
		b=temp
	}
	return a
}
