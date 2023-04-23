//找出数组中所有的子数组最大和
package main

import "fmt"

func main() {
	arr := []int{1, -2, 4, 8, -4, 7, -1, -5}
	fmt.Println(findMaxSum(arr))

	fmt.Println("子数组为：", arr[1:1])
	fmt.Println("子数组为：", arr[1:len(arr)])
}

func findMaxSum(arr []int) (newArry []int, sum int) {

	for i := 0; i < len(arr); i++ {
		for k := i; k < len(arr); k++ {
			subArr := arr[i:k]
			var subSum int
			for _, num := range subArr {
				subSum += num
				if subSum > sum {
					newArry = subArr
					sum = subSum
				}
			}
		}
	}
	return
}
