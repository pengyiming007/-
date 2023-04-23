// 求一个字符串中最长的重复子串，如"banana"的最长重复子串为ana;
package main

import "fmt"

func main() {
	str := "banana"
	// fmt.Println(listSubString(str))

loop:
	for i := 1; i < len(str); i++ {
		// 计算最长的所有子串中，是否存在相等的
		var subStrs []string
		for k := 0; k <= i; k++ {
			temp := string(str[k : len(str)-i+k])
			subStrs = append(subStrs, temp)
		}

		fmt.Println("subStrs:", subStrs)
		subStrMap := make(map[string]bool)
		for _, subStr := range subStrs {
			if subStrMap[subStr] {
				fmt.Println("max length subStr:", subStr)
				break loop
			}
			subStrMap[subStr] = true
		}
	}
}

// 从长到短获取字符串中的所有子串
func listSubString(str string) (subStrs []string) {
	for i := 1; i < len(str); i++ {
		for k := 0; k <= i; k++ {
			temp := string(str[k : len(str)-i+k])
			subStrs = append(subStrs, temp)
		}
	}
	return
}
