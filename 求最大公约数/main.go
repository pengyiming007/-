package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("请输入第1个数:")
	fmt.Scanln(&num1)
	fmt.Print("请输入第2个数:")
	fmt.Scanln(&num2)
	if maxCommonDivisor(num1, num2) > 0 {
		fmt.Println("最大公约数为:", maxCommonDivisor(num1, num2))
		return
	}
	fmt.Println("这两个数没有最大公约数")
}

func maxCommonDivisor(num1, num2 int) (divisor int) {
	min := num1
	if num1 > num2 {
		min = num2
	}

	for div := min; div > 1; div-- {
		if num1%div == 0 && num2%div == 0 {
			divisor = div
			return
		}
	}
	return
}
