/*
   冒泡排排序
   1. 将第一个元素与第二个元素比较，若为逆序，则交换位置
   2. 然后比较第二个元素与第三个元素，依次比较，直到第n-1和n比较完为止
   3. 该过程为第一轮冒泡，其结果是最大的记录被交换到第n个记录位置上
   4. 再进行第二轮冒泡，对前n-1个记录进行同样的操作
   其时间复杂度为O(n的平方)，空间复杂为O(1)
*/

/*
快速排序的思想：取一个枢轴数作为分界，并将记录划分为两部分，前一部分的元素均不大于后一部分的元素值
1. 设置枢轴记录(通常取第1个元素)，起点i、终点j；
2. 从终点分别与枢轴记录比较，如果大于，则j--,并继续比较，如果小于枢轴值，则将s[j]替换至枢轴值的位置，即s[i]=s[j]；
3. 再拿起点与枢轴逐个比较，如果小于则i++，并继续比较,如果大于，则将该值替换至步骤2中j的位置，即s[j]=s[i];
4. 重复步骤2、3，直到i>j，最后再将pivot放在s[i]，完成一轮比较
5. 递归调用1~4， 每次调用都将数组分割成2个，直到数据长度为1，不再能分割
其算法复杂度为O(nlogn)
*/

/*
简单选择排序：依次从数组中的第1个元素开始与之后每一个元素比较，如果后面的元素比当前元素值小，就交换位置
算法复杂度为O(n的平方)
*/
package main

import (
	"fmt"
)

func main() {
	numlist := []int{11, 1, 3, 2, 5, 19, 4, 25}
	// fmt.Println(numlist[1:])
	// sort3(numlist)
	// sort4(numlist)
	// QuickSort(numlist, 0, len(numlist)-1)
	// fmt.Println(numlist)
	// bubble(numlist)
	// fmt.Println(numlist)
	// fmt.Println(partition2(numlist, 0, len(numlist)-1))
	// quickSort2(numlist, 0, len(numlist)-1)
	// partition2(numlist, m+1, len(numlist)-1)
	simpleSort(numlist)
	fmt.Println(numlist)
}

// 冒泡算法，相邻两两比较
func bubble(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		k := 0
		j := 1
		for k < i && j < i {
			if s[k] > s[j] {
				s[k], s[j] = s[j], s[k]
			}
			k++
			j++
		}
	}
}

// 简单选择排序
func simpleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
func quickSort(s []int, start, end int) {
	if start >= end {
		return
	}
	i := partition2(s, start, end)
	quickSort(s, start, i)
	quickSort(s, i+1, end)
}

func partition2(s []int, start, end int) int {
	pivot := s[start]
	for start < end {
		// 先从尾端开始逐个比较，大的不动，小的更换与pivot的位置
		for start < end {
			if s[end] > pivot {
				end--
				continue
			}
			s[start], s[end] = s[end], s[start]
			break
		}
		for start < end {
			if s[start] < pivot {
				start++
				continue
			}
			s[start], s[end] = s[end], s[start]
			break
		}
	}
	return start
}

func sort1(numList []int) {
	for i := range numList {
		fmt.Println("外循环：", numList[:i])
		for k := range numList {
			if numList[i] < numList[k] {
				fmt.Println("内循环：", numList)
				numList[i], numList[k] = numList[k], numList[i]
			}
		}
	}
	return
}

//递归法
// func sort2(list []int) []int {
// 	if len(list) == 1 {
// 		return list
// 	}
// 	l := sort2(list[1:])
// 	for i, v := range l {
// 		if list[0] > l[i] {
// 			newList
// 		}
// 	}
// 	newList := append(l, list[0])
// 	return newList
// }

// 冒泡排
func sort3(numList []int) {
	for i := 0; i < len(numList); i++ {
		for k := len(numList) - 1; k > i; k-- {
			if numList[i] > numList[k] {
				numList[i], numList[k] = numList[k], numList[i]
			}
		}
	}
}

// 冒泡排
func sort4(numList []int) {
	for i := 0; i < len(numList); i++ {
		for k := i + 1; k < len(numList); k++ {
			if numList[i] > numList[k] {
				numList[i], numList[k] = numList[k], numList[i]
			}
		}
	}
}

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	pivotIndex := partition(arr, startIndex, endIndex)
	QuickSort(arr, startIndex, pivotIndex)
	QuickSort(arr, pivotIndex+1, endIndex)
}

// 快速排序
func partition(arr []int, start, end int) int {
	var (
		mark  = start
		point = start + 1
	)
	for point < end {
		if arr[point] < arr[start] {
			mark++
			// if mark != point {
			arr[mark], arr[point] = arr[point], arr[mark]
			// }
		}
		point++
	}
	if arr[start] != arr[mark] {
		arr[start], arr[mark] = arr[mark], arr[start]
	}
	return mark
}
