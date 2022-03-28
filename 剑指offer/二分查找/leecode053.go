package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetNumberOfK([]int{1,2,3,3,3,3,4,5},3))
}

func GetNumberOfK( data []int ,  k int ) int {
	leftIndex,rightIndex:=0,0
	lleft,lright:=0,len(data)
	for lleft < lright {
		mid:=lleft+(lright-lleft)/2
		if data[mid] < k {
			lleft=mid+1
		}else{
			lright=mid
		}
	}
	leftIndex=lleft
	rleft,rright:=0,len(data)
	for rleft < rright {
		mid:=rleft+(rright-rleft)/2
		if data[mid] <= k {
			rleft=mid+1
		}else{
			rright=mid
		}
	}
	rightIndex=rleft
	return rightIndex-leftIndex

}