//在不排序的情况下，找出数组的中位数
package main

import "fmt"

func main() {
	arr := []int{7, 5, 3, 1, 11, 9}
	fmt.Println(getMid(arr))
}

var pos = 0

func partition(arr []int, low, high int) {
	key := arr[low]
	for low < high {
		for low < high && arr[high] > key {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] < key {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = key
	pos = low
}

func getMid(arr []int) int {
	low := 0
	n := len(arr)
	high := n - 1
	mid := (low + high) / 2
	for true {
		partition(arr, low, high)
		if pos == mid {
			break
		} else if pos > mid {
			high = pos - 1
		} else {
			low = pos + 1
		}
	}
	fmt.Println(arr)
	if (n % 2) != 0 {
		return arr[mid]
	} else {
		return (arr[mid] + arr[mid+1]) / 2
	}
}
