package main

import "fmt"

type Stack struct {
	arr  []interface{}
	size int
}

func main() {
	var s []interface{}
	s = append(s, "a")
	stack := &Stack{
		arr:  s,
		size: len(s),
	}
	stack.Push("b")   
	fmt.Println(stack.Pop())

}

func (s *Stack) Push(value interface{}) {
	s.size++
	s.arr = append(s.arr, value)
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		temp := s.arr[s.size-1]
		s.size--
		s.arr = s.arr[:s.size]
		return temp
	}
	return nil
}
