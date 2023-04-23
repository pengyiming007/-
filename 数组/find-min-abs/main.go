// 从升序数组中（有负数、正数）找出绝对值最小的数
package main

import "fmt"

func main() {
	arr := []int{-10, -9, -5, -2, -1, 2, 5}
	min := find(arr)
	fmt.Println(min)
}

func find(arr []int) (min int) {
	for i, num := range arr {
		if num == 0 {
			return num
		}
		if num > 0 {
			if num+arr[i-1] > 0 {
				return arr[i-1]
			} else {
				return num
			}
		}
	}
	return
}
