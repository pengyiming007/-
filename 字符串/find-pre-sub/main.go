//找出字符串数组中的最大前缀子串

package main

import "fmt"

func main() {
	strs := []string{"flow", "fled", "flyi"}
	findMaxPre(strs)
}

func findMaxPre(strs []string) {
	str1 := strs[0]
	s1 := []rune(str1)
	var subStr []rune
	for i, char := range s1 {
		var has bool
		for _, str := range strs {
			ss := []rune(str)
			if char == ss[i] {
				has = true
			} else {
				has = false
				break
			}
		}
		if has {
			subStr = append(subStr, char)
		} else {
			break
		}
	}
	fmt.Println(string(subStr))
}
