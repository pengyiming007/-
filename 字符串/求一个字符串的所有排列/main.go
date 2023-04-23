// 打印出字符串的所有排列组合，如abc,acb,bac,bca,cab,cba
package main

import "fmt"

func main() {
	str := "abc"
	permutation([]rune(str), 0)
}

// 递归实现的步骤：1. 寻找递归，找出大规模与小规模之间的联系，通过递归将规模不断缩小；2. 定义好递归结束的条件，避免无限循环；3.定义函数、与参数，确定递归要做什么
// 先将第一个元素作为start，先全排start+1及以后的元素；再分别对换start和以后元素的位置，再全排start+1以后的元素；再将start和后面的元素换回来
func permutation(str []rune, start int) {
	if len(str) == 0 {
		return
	}
	if start == len(str)-1 {
		fmt.Println(string(str))
		return
	}

	permutation(str, start+1)
	for i := start + 1; i < len(str); i++ {
		str[start], str[i] = str[i], str[start]
		permutation(str, start+1)
		str[start], str[i] = str[i], str[start]
	}
}
