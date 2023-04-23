// 将一个包含大小写的字符串，将所有小写字母排在大写字母的前面，不要求对大小写字母之间排序

package main

import "fmt"

func main() {
	str := "acfeBAcZef"
	fmt.Println(sortStr(str))
}

// i=0,j=len(str),开始遍历字符串，如果i>=j结束
// 如果第str[0]为小写，i++; 如果str[i]为大写,则判断str[j]，
// 如果str[j]为大写则j--，如果str[j]为小写，则与str[i]对换位置,并将i++,j--
func sortStr(input string) (reply string) {
	str := []rune(input)
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] <= 'z' && str[i] >= 'a' {
			i++
			continue
		}

		if str[j] >= 'A' && str[j] <= 'Z' {
			j--
			continue
		}
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}

	return string(str)
}
