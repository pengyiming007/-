package main

import "fmt"

//找到数组中唯一重复的元素
func main() {
	numList := []int{1, 2, 3, 4, 5, 6}
	duplicateNum := findDuplicate(numList)
	fmt.Println(duplicateNum)
}

func findDuplicate(numList []int) (n int) {
f:
	for i, num := range numList {
		for i2, num2 := range numList {
			if num == num2 && i != i2 {
				n = num
				break f
			}
		}
	}
	return
}
