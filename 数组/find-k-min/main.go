//找一个数组中第k小的数,思路：先排序，再找第K个数。
package main

import "fmt"

func main() {
	arr := []int{1, 0, 3, 2, 4, 2, 1}
	findkMin(3, arr)
}

func findkMin(k int, arr []int) {
	for i := range arr {
		for k := range arr {
			if arr[i] < arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
			}
			fmt.Printf("i:%d,k:%d,arr:%v\n", i, k, arr)
		}
	}

	fmt.Println(arr)
	temp := arr[0]
	index := 1
	for i := range arr {
		if temp < arr[i] {
			index++
		}
		if index == k {
			fmt.Printf("第%d小的数是：%d\n", k, arr[i])
			break
		}
		temp = arr[i]
	}
}



