package main

func main()  {
	
}

/**
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func reOrderArray( array []int ) []int {
	for i,j := 0,0; i < len(array); i++ {
		if array[i]&i==1 {
			array[i],array[j]=array[j],array[i]
			j++
		}
	}
	return array
}
