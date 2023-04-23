package main

import "fmt"

//链表定义
type LNode struct {
	data interface{}
	next *LNode
}

func reverse(node *LNode) {
	if node == nil || node.next == nil {
		return
	}
	var pre *LNode    //定义前驱结点
	var cur *LNode    //定义当前结点
	next := node.next //把后继结点存起来，防止丢失
	for next != nil { //node.next为空说明遍历到最后一个了，反之说明没有到最后
		cur = next.next
		next.next = pre //next指向pre，实现逆序
		pre = next      //后移前驱结点
		next = cur      //后移后驱结点
	}
	node.next = pre //让链表的第一个结点指向头结点
}

func createNode(node *LNode, max int) {
	//创建链表
	cur := node
	for i := 1; i <= max; i++ {
		cur.next = &LNode{} //分配内存
		cur.next.data = i
		cur = cur.next //向后移动指针
	}
}

func createNode2(head *LNode, len int) {
	cur := head
	for i := 1; i < len; i++ {
		cur.next = &LNode{}
		cur.data = i
		cur = cur.next
	}
	cur.data = len
	cur.next = nil
}

func printNode2(head *LNode) {
	cur := head
	for {
		if cur.data != nil {
			fmt.Printf("%d ", cur.data)
		}
		cur = cur.next
		if cur == nil {
			break
		}
	}
	fmt.Printf("\n")
}

//打印链表
func printNode(info string, node *LNode) {
	fmt.Print(info)                                    //打印“逆序前或逆序后”
	for cur := node.next; cur != nil; cur = cur.next { //遍历链表打印
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}

func revertPrint(head *LNode) {
	if head == nil {
		// fmt.Printf("%d ", head.data)
		return
	}
	revertPrint(head.next)
	fmt.Printf("%d ", head.data)
}

func main() {
	head := &LNode{}
	// fmt.Println("就地逆序")
	createNode(head, 5)
	// printNode("逆序前", head)
	// reverse2(head)
	// printNode("逆序后", head)

	// printNode2(head)
	// newHead := reverse3(head)
	// printNode2(newHead)

	printNode2(head)
	reverse2(head)
	printNode2(head)
	// revertPrint(head)
}

func reverseChild(node *LNode) *LNode {
	if node == nil || node.next == nil {
		return node
	}
	newHead := reverseChild(node.next)
	node.next.next = node
	node.next = nil

	return newHead
}

//递归法逆序
func reverse2(node *LNode) {
	firstNode := node.next
	newHead := reverseChild(firstNode)
	node.next = newHead
}

// 递归
func reverse3(head *LNode) (newHead *LNode) {
	// 当前节点的下一个节点为空，或者当前节点为空时，递归结束
	if head.next == nil || head == nil {
		return head
	}

	// 先逆序除了当前节点，后面的所有节点
	n := head.next
	newHead = reverse3(n)
	// 再将当前节点逆序放在最后端，当前节点的下一个节点为空
	n.next = head
	head.next = nil
	return
}
