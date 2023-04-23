package main

import "fmt"

type BNode struct {
	Data  int
	Left  *BNode
	Right *BNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	root := Create(arr)
	fmt.Printf("%v\n", *root)
	PrintTree(root)

	var arr1 []int
	PrintBTree2(&arr1, root)
	fmt.Println(arr1)
}

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

func PrintBTree2(arr *[]int, root *BNode) {
	if root.Data != 0 {
		*arr = append(*arr, root.Data)
	}
	if root.Left != nil {
		PrintBTree2(arr, root.Left)
	}
	if root.Right != nil {
		PrintBTree2(arr, root.Right)
	}
}
