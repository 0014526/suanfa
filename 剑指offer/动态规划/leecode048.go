package main

func lengthOfLongestSubstring(s string) int {
	m:=make(map[byte]struct{})
	max:=0
	for left,right := 0,0; right < len(s);{
		if _,ok:=m[s[right]];ok {
			delete(m,s[left])
			left++
		}else{
			m[s[right]]= struct{}{}
			right++
		}
		if right-left>max {
			max=right-left
		}
	}
	return max
}