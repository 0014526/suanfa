package main

func main() {

}

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution( numbers []int ) int {
	// write code here

	result:=make(map[int]int)
	for _, number := range numbers {
		if _, ok := result[number]; ok {
			result[number]=result[number]+1
		}else{
			result[number]=1
		}
	}
	for key, value := range result {
		if value>len(numbers)/2 {
			return key
		}
	}
	return 0
}
