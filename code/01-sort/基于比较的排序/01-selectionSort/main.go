package main

import "fmt"

// 选择排序  O(N^2)
// 每次从余下的数中找最小的，并排到余下的数的最开头。
func selectionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		min := i //初始的最小值位置从0开始，依次向右

		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{9, 7, 8, 6, 2, 1, 3}
	selectionSort(arr)
}
