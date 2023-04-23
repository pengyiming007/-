// 给定一棵二叉树，它的每个结点是正整数或负整数，找到一棵子树，使它的所有结点的和最大
package main

import "fmt"

type BNode struct {
	Data  int
	Left  *BNode
	Right *BNode
}

var m = make(map[int]*BNode, 0)

func main() {
	arr := []int{-11, 3, 6, 9, -7}
	root := Create(arr)

	var sum int
	sum = Sum(root)
	fmt.Println(sum)

	max := FindMaxTree(root)
	fmt.Printf("最大子树和：%d, 节点值为:%v\n", max, m[max])
	PrintTree(m[max])
}

// 获取最大子树
func FindMaxTree(node *BNode) (max int) {
	if node == nil {
		return
	}
	sum := Sum(node)
	if sum > max {
		max = sum
		m[max] = node
	}
	if node.Left != nil {
		sum := Sum(node.Left)
		if sum > max {
			max = sum
			m[max] = node.Left
		}
	}
	if node.Right != nil {
		sum := Sum(node.Right)
		if sum > max {
			max = sum
			m[max] = node.Right
		}
	}
	return
}

// 当前节点下的子树求和
func Sum(root *BNode) (sum int) {
	sum += root.Data
	if root.Left != nil {
		sum += Sum(root.Left)
	}
	if root.Right != nil {
		sum += Sum(root.Right)
	}
	return
}

// 创建二叉树
func Create(arr []int) (node *BNode) {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		node = new(BNode)
		node.Data = arr[0]
		return
	}

	c := len(arr) / 2
	node = new(BNode)
	node.Data = arr[c]

	left := Create(arr[:c])
	right := Create(arr[c+1:])
	node.Left = left
	node.Right = right
	return
}

func PrintTree(root *BNode) {
	if root == nil {
		return
	}
	fmt.Printf("Node:%d\n", root.Data)
	if root.Left != nil {
		fmt.Printf("Node:%d, has left tree: ", root.Data)
		PrintTree(root.Left)
	}

	if root.Right != nil {
		fmt.Printf("Node:%d, has right tree: ", root.Data)
		PrintTree(root.Right)
	}
}
