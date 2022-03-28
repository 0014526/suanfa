package main

import (
	"fmt"
	"sort"
	"strings"
)

func PrintMinNumber( numbers []int ) string {
	sort.Slice(numbers, func(i, j int) bool {
		return compareNum(numbers[i],numbers[j])
	})
	var str strings.Builder
	for i := 0; i < len(numbers); i++ {
		str.WriteString(fmt.Sprintf("%d",numbers[i]))
	}
	return str.String()
}

func compareNum(i, j int) bool {
	str1:=fmt.Sprintf("%d%d",i,j)
	str2:=fmt.Sprintf("%d%d",j,i)
	if str1>str2 {
		return true
	}
	return false
}