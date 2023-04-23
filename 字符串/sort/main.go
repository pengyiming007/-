//对字符串进行全排列

package main

import "fmt"

func main() {
	str := "abc"
	// arr := []rune(str)
	permutation([]rune(str), 0)
	// permutation2(arr, 0, 2)
	// CalcALLPermutation(str, 0, 1)
}

// 书本上的解法
func permutation(arr []rune, start int) {
	if arr == nil {
		return
	}
	if start == len(arr)-1 {
		fmt.Println(string(arr))
	} else {
		for i := start; i < len(arr); i++ {
			// fmt.Println("交换前：", string(arr))
			arr[start], arr[i] = arr[i], arr[start]
			fmt.Println("递归前：", string(arr))
			permutation(arr, start+1)
			fmt.Println("递归后：", string(arr))
			arr[start], arr[i] = arr[i], arr[start]

		}
	}
}

func permutation1(arr []rune, start int) {
	if arr == nil {
		return
	}
	if start == len(arr)-1 {
		fmt.Println(string(arr))
	} else {
		for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start]
			permutation(arr, start+1)
			arr[start], arr[i] = arr[i], arr[start]
		}
	}
}

func permutation2(arr []rune, from, to int) {
	if to <= 1 {
		return
	}
	if from == to {
		fmt.Println(string(arr))
	} else {
		for j := from; j <= to; j++ {
			arr[j], arr[from] = arr[from], arr[j]
			permutation2(arr, from+1, to)
		}
	}
}

func CalcALLPermutation(str string, from int, to int) {
	strr := []rune(str)
	if to <= 1 {
		return
	}
	if from == to {
		fmt.Println("最终结果：", string(strr))
	} else {
		for j := from; j <= to; j++ {
			strr[j], strr[from] = strr[from], strr[j]
			str1 := string(strr)
			CalcALLPermutation(str1, from+1, to)
		}
	}
}
