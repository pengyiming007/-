// 将字符串(1,(2,3),4,(5,6))变为(1,2,3,4,5,6)
package main

import "fmt"

func main() {
	str := "(1,(2,3),4,(5,6))"
	fmt.Println(transform(str))
}

func transform(input string) (output string) {
	str := []rune(input)
	fmt.Println("start:", string(str[0]))
	if string(str[0]) != "(" || string(str[len(str)-1]) != ")" {
		fmt.Println("string format error")
		return
	}
	for i := 1; i < len(str)-1; i++ {
		if string(str[i]) == "(" || string(str[i]) == ")" {
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}

	return string(str)
}
