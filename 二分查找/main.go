// 在一个有序排列的数组中，使用二分查找，每次以中间数作为左边或者右边的边界来查找目标
package main

import "fmt"

var i int

func main() {
	s := []int{0, 1}
	target := 2
	println(binarySearch(target, s, 0, len(s)-1))
}

// 使用递归进行二分查找
func binarySearch(target int, nums []int, left, right int) int {
	i++
	if left > right {
		return -1
	}
	if target < nums[left] {
		return -1
	}
	if target > nums[right] {
		return -1
	}
	mid := (right + left + 1) / 2
	fmt.Printf("第%d次执行,mid:%d\n", i, mid)
	if target == nums[mid] {
		return mid
	}
	if target > nums[mid] {
		return binarySearch(target, nums, mid+1, right)
	}
	if target < nums[mid] {
		return binarySearch(target, nums, left, mid-1)
	}
	return -1
}
