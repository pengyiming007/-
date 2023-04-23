package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(revert(s))
}
func revert(in []int) []int {
	head := 0
	rear := len(in) - 1
	if len(in) == 0 {
		return nil
	}

	for {
		in[head], in[rear] = in[rear], in[head]
		head++
		rear--
		if len(in)%2 == 0 {
			if rear == head+1 {
				break
			}
		}
		if len(in)%2 == 1 {
			if rear == head {
				break
			}
		}
	}
	return in
}
