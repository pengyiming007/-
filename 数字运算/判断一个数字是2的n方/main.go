// 使用递归判断一个数字是否为2的n次方
package main

import "fmt"

func main() {
	var num int
	num = 64
	fmt.Println(isPower2(num))

	fmt.Println(isPower(num))
}

// 递归
func isPower2(num int) bool {
	if num == 1 {
		return true
	}
	if num%2 != 0 {
		return false
	}
	n := num >> 1 //左移1位，相当于除2
	return isPower2(n)
}

// 位移法
func isPower(num int) bool {
	i := 1
	for {
		if i == num {
			return true
		}
		if i > num {
			return false
		}
		i = i << 1
	}
}
