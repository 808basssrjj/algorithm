package main

import "fmt"

// 冒泡排序 O(N^2)
func bubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{9, 7, 8, 6, 2, 1, 3}
	bubbleSort(arr)
}
