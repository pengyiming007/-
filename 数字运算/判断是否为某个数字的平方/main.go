// 使用递归求一个数字是否为某个数的平方，如：64是8的平方
package main

import "fmt"

func main() {
	var num int
	num = 81
	fmt.Println(isPower(num, 1, num))
}

func isPower(num, low, high int) bool {
	if high-low < 2 {
		return false
	}
	mid := (low + high) / 2
	if num == mid*mid {
		return true
	}
	if num < mid*mid {
		high = mid
		return isPower(num, low, high)
	}
	low = mid
	return isPower(num, low, high)
}
