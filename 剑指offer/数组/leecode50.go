package main

func main()  {
	
}

// 遍历一次存到数组中，然后找到数组
func firstUniqChar(s string) byte {
	var list [26]int
	length:=len(s)
	for i := 0; i < length; i++ {
		list[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if list[s[i]-'a']==1 {
			return s[i]
		}
	}
	return ' '
}
