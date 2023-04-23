// 求字符串中连续出现相同字符的最大值，如字符串"aaabbc"中"a"的最大值为3
package main

import "fmt"

func main() {
	str := "aaabccccc"
	fmt.Println(getMaxDupWord(str, 0, 1, 1))
	fmt.Println(getMaxDupWord2(str))
}

// 递归方法
func getMaxDupWord(s string, start, currentLen, maxLen int) int {
	if start == len(s)-1 {
		return max(currentLen, maxLen)
	}

	if s[start] == s[start+1] {
		return getMaxDupWord(s, start+1, currentLen+1, maxLen)
	} //else {
	return getMaxDupWord(s, start+1, 1, max(currentLen, maxLen))
	//}
}
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

// 非递归
func getMaxDupWord2(str string) int {
	currentLen := 1
	maxLen := 1
	var word string
	for i := 0; i < len(str); i++ {
		if string(str[i]) == word {
			currentLen++

			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			word = string(str[i])
			currentLen = 1
		}
	}
	return maxLen
}
