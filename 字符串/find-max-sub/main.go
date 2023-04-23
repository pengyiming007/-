// 找出两个字符串中最大公共子串
package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "dbcca"
	str2 := "dgccadede"
	fmt.Println(string(findMaxSub([]rune(str1), []rune(str2))))

	fmt.Println(findMaxSubString([]rune(str1), []rune(str2)))
}

func findMaxSub(arr1, arr2 []rune) (subArr []rune) {
	max := 0
	for i := 0; i < len(arr1); i++ {
		for j := i; j < len(arr1); j++ {
			sub1 := arr1[i : j+1]

			for k1 := 0; k1 < len(arr2); k1++ {
				for k2 := k1 + 1; k2 < len(arr2); k2++ {
					sub2 := arr2[k1 : k2+1]
					if string(sub1) == string(sub2) {
						if len(sub1) > max {
							max = len(sub1)
							subArr = sub1
						}
					}
				}
			}
		}
	}
	return
}

func findMaxSubString(arr1, arr2 []rune) (sub string) {
	short := arr1
	long := arr2
	if len(arr2) < len(arr1) {
		short = arr2
		long = arr1
	}
	l := len(short)
	for i := 0; i < len(short); i++ {
		for j := 0; j <= i; j++ {
			sub := short[j : l-i+j]
			if strings.Contains(string(long), string(sub)) {
				return string(sub)
			}
		}
	}
	return
}
