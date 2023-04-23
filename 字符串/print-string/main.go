package main

import "fmt"

func main() {
	str := "abcde"
	for _, c := range str {
		fmt.Printf("%s", string(c))
	}
}
