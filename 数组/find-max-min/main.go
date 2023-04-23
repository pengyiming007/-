//找出数组中最大值和最小值
package main

import "fmt"

func main() {
	numlist := []int{1, 2, 3, 4, 5, 6, 100}
	max, min := findMaxMin(numlist)
	fmt.Println(max, min)
}

func findMaxMin(numList []int) (max, min int) {
	if len(numList) > 0 {
		max, min = numList[0], numList[0]
	}
	for _, num := range numList {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return
}

