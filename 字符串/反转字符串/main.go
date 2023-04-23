package main

import "fmt"

func main() {
	str := "我是中国人"
	fmt.Println(revert(str))
}

// 1. 将字符串转为rune类型的数组；2. 依次将首(start)与尾(end)字符换位置，直接到start>=end
func revert(input string) (output string) {
	s := []rune(input)
	if s == nil {
		return
	}
	start := 0
	end := len(s) - 1

	for {
		if start >= end {
			break
		}
		
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

	output = string(s)
	return
}
