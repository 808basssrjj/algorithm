package main

import "fmt"

// 插入排序 O(N^2)
func insertionSort(arr []int) {
	if  len(arr) < 2 {
		return
	}

	// 0~0 有序的
	// 0~i 想有序
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{7, 6, 2, 3, 4, 1}
	insertionSort(arr)
}
