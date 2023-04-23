//判断两个字符串是否为换位字符串
package main

import "fmt"

func main() {
	str1 := "aaabbcde"
	str2 := "abcbaadee"
	fmt.Println(compare([]rune(str1), []rune(str2)))
}

func compare(arr1, arr2 []rune) bool {
	var result bool
	strMap := make(map[string]string)
	for _, word := range arr1 {
		strMap[string(word)] = string(word)
	}
	for _, word := range arr2 {
		_, has := strMap[string(word)]
		if !has {
			return has
		}
		delete(strMap, string(word))
	}
	if len(strMap) == 0 {
		result = true
	}
	return result
}
