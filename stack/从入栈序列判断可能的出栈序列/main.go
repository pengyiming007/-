package main

import "fmt"

type stackSlice struct {
	s []int
}

var pushMap = make(map[int]bool)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{3, 2, 5, 4, 1}
	var s stackSlice
	for _, num2 := range s2 {
		if pushMap[num2] {
			if s.pop() == num2 {
				fmt.Println("pop ", num2)
				continue
			} else {
				fmt.Printf("%d error\n", num2)
				break
			}
		}
		for _, num1 := range s1 {
			s.push(num1)
			s1 = s1[1:]
			fmt.Println("push ", num1)
			fmt.Println(s)
			if num1 == num2 {
				fmt.Println("pop ", s.pop())
				break
			}
		}

	}
}

func (stack *stackSlice) pop() (v int) {
	size := len(stack.s)
	if size > 0 {
		v = stack.s[size-1]

		stack.s = stack.s[:size-1]
		delete(pushMap, v)
		return
	}
	return 0
}

func (stack *stackSlice) push(v int) {
	stack.s = append(stack.s, v)
	pushMap[v] = true
}
