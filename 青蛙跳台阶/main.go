// 青蛙每一次跳1级或者2级台阶，那么跳上N级台阶总共有多少种跳法？
package main

import "fmt"

func main() {
	fmt.Println(Num(30))
}

// 递归方式
func Num(n int) (m int) {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	m = Num(n-1) + Num(n-2)
	return
}
