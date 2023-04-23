// 获取一个字符串的最大回文子串，回文串即为左右对称的字符串

// 解题思路：从大到小依次取字符串的子串，再较验这些子串是否为回文串
package main

import "fmt"

func main() {
	str := "abcbad"
	s := []rune(str)
	for i := 0; i < len(str); i++ { //总共截取i
		for j := 0; j <= i; j++ { // 取子串时，左边截取j,那么右边截取i-j，子串即s[j:len-(i-j)]
			subStr := s[j : len(str)-i+j]
			if IsSymmetry(subStr) {
				fmt.Printf("最长子串的长度是：%d, 子串是：%s", len(str)-i, string(subStr))
				return
			}
		}
	}
}

func IsSymmetry(str []rune) bool {
	start := 0
	end := len(str) - 1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}
