//从有序不重复的数组中找出相等的数
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(findNum2(arr, 3, 0, 7))
}
func findNum(arr []int, num int) (index int) {
	length := len(arr)
	var middle int
	if length%2 == 0 {
		middle = length / 2
	} else {
		middle = (length + 1) / 2
	}

	for i := middle; i > 0 && i < length; {
		if num > arr[i] {
			i++
			continue
		}
		if num < arr[i] {
			i--
			continue
		}
		if num == arr[i] {
			return i
		}
	}
	return
}

// 递归取中间数比较
func findNum2(arr []int, num, start, end int) (index int) {
	if start == end || start <= 0 || end >= len(arr) {
		return
	}
	length := len(arr)
	var middle int
	if length%2 == 0 {
		middle = length / 2
	} else {
		middle = (length + 1) / 2
	}

	if num == arr[middle] {
		return middle
	}

	if num > arr[middle] {
		start = middle
	}
	if num < arr[middle] {
		end = middle + 1
	}

	index = findNum2(arr[start:end], num, start, end)
	return
}
