package main

func main() {
	
}

func countDigitOne(n int) int {
	// 这里digitNum为digit所在位
	digitNum,sum:=1,0
	high,cur,low:=n/10,n%10,0
	for high!=0||cur!=0 {
		if cur<1 {
			sum+=high*digitNum
		} else if cur==1{
			sum+=high*digitNum+low+1
		} else{
			sum+=(high+1)*digitNum
		}
		// 换下一位，更新高低位
		low=low+cur*digitNum
		high,cur=high/10,high%10
		digitNum=digitNum*10
	}
	return sum
}