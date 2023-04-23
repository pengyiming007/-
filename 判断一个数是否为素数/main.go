package main

import "fmt"

func main() {
	var num int
	fmt.Print("请求输入一个数:")
	fmt.Scanln(&num)
	if isPrime(num) {
		fmt.Printf("%d 是素数", num)
		return
	}
	fmt.Printf("%d 不是素数", num)
}

func isPrime(num int) (ok bool) {
	ok = true
	var n int
	if num%2 == 0 {
		n = num / 2
	} else {
		n = (num + 1) / 2
	}
	for i := 2; i < n; i++ {
		if num%i == 0 {
			fmt.Printf("%d能整除%d\n", num, i)
			ok = false
			return
		}
	}
	return
}
