//反转字符串

package main

import "fmt"

func main() {
	str := "abcde"
	revert([]rune(str))
}

func revert(arr []rune) {
	newArray := make([]rune, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		newArray = append(newArray, arr[i])
	}
	fmt.Println(string(newArray))
}
