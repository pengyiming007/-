package main

import "fmt"

//链表定义
type Node struct {
	data interface{}
	next *Node
}

func main() {
	head := &Node{}
	createNode(head, 5)
	printNodeList(head)
	remove(head)
	printNodeList(head)
}

//打印链表
func printNodeList(head *Node) {
	cur := head
	for {
		if cur.data != 0 {
			fmt.Printf("%d ", cur.data)
		}
		cur = cur.next
		if cur == nil {
			break
		}

	}
	fmt.Printf("\n")
}

//创建链表
func createNode(head *Node, len int) {
	cur := head
	for i := 1; i < len; i++ {
		cur.next = &Node{}
		cur.data = i
		cur = cur.next
		//重复链表
		if i == 3 {
			cur.next = &Node{}
			cur.data = i
			cur = cur.next
		}
	}
	cur.data = len
	dupNode := &Node{data: 2} //增加重复链表
	cur.next = dupNode
}

// 顺序法
func remove(head *Node) {
	if head == nil || head.next == nil {
		return
	}
	var innerCur, innerPre *Node
	for outerCur := head; outerCur != nil; outerCur = outerCur.next {
		for innerCur, innerPre = outerCur.next, outerCur; innerCur != nil; {
			if outerCur.data == innerCur.data {
				innerPre.next = innerCur.next
				innerCur = innerCur.next
			} else {
				innerPre = innerCur
				innerCur = innerCur.next
			}
		}
	}
}

// 递归法
func remove2(head *Node) {
	if head.next == nil {
		return
	}

	remove2(head.next)
	var cur *Node
	pre := head
	cur = head.next
	for cur != nil {
		if head.data == cur.data {
			pre.next = cur.next
			cur = pre.next
		} else {
			cur = cur.next
			pre = pre.next
		}
	}
	return
}
