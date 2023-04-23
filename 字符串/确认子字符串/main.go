// 确认一个字符串是否是另一个字符串的子串，并求其第一次出现的下标
package main

import "fmt"

func main() {
	str1 := "abcdefgh"
	str2 := "z"
	fmt.Println(subString(str2, str1))
}

// 该方法中存在bug，当短串中第一个字符在长串出现多次，而第1次出现并不是子串
// 修改获取start的方法，如果后面匹配不相等，则再次继续往后遍历长串，并获取后面与短串第1个字符相等的下标
func subString(str1, str2 string) (start int, isSub bool) {
	// 先分别取长、短字符串
	long := []rune(str1)
	short := []rune(str2)
	if len(str1) < len(str2) {
		long = []rune(str2)
		short = []rune(str1)
	}
	// 将短串的第1个字符去遍历匹配长串中的字符，如果相等，则取其下标
	for i := range long {
		if short[0] == long[i] {
			start = i
			break
		}
	}

	// 遍历短串，分别与长串中的字符比较
	isSub = true
	for i := range short {
		if short[i] != long[start+i] {
			start = 0
			isSub = false
		}
	}
	return
}
