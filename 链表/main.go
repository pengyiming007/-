package main

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

func main() {
	head := new(Node)
	CreateNodeList(head, 5)
	PrintNodeList(head)
	newHead := reverse(head)
	PrintNodeList(newHead)
}

// 创建从1~size的链表
func CreateNodeList(head *Node, size int) {
	cur := head
	for i := 1; i < size; i++ {
		cur.Data = i
		next := new(Node)
		cur.Next = next
		cur = next
	}
	cur.Data = size
}

// 依次打印链接
func PrintNodeList(head *Node) {
	cur := head
	for {
		if cur.Data != nil {
			fmt.Printf("%d ", cur.Data)
		}
		if cur.Next != nil {
			cur = cur.Next
			continue
		}
		fmt.Printf("\n")
		break
	}
}

func printNode2(head *Node) {
	cur := head
	for {
		if cur.Data != nil {
			fmt.Printf("%d ", cur.Data)
		}
		cur = cur.Next
		if cur == nil {
			break
		}
	}
	fmt.Printf("\n")
}

// 反转链表
func reverse(head *Node) (newHead *Node) {
	if head.Next == nil || head == nil {
		return head
	}

	next := head.Next
	newHead = reverse(next)
	next.Next = head
	head.Next = nil
	return
}
