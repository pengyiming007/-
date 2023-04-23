package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	permutation(arr, 0)
}

func permutation(arr []int, start int) {
	fmt.Println("start:", start)
	if arr == nil {
		return
	}
	if start == len(arr)-1 {
		fmt.Println(arr)
	} else {
		for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start]
			permutation(arr, start+1)
			arr[start], arr[i] = arr[i], arr[start]
		}
	}
}
