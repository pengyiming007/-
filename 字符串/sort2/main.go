//对字母进行大小写排序
package main

func main() {

}

func sort(arr []rune) (newStr string) {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			last := len(arr) - 1
			for j := last; j > i; j-- {
				if arr[j] <= 'z' && arr[j] >= 'a' {

				}
			}
		}
	}
	return
}
