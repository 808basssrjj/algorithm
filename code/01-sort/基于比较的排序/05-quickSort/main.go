package main

import (
	"fmt"
	"math/rand"

	// "sort"
	"time"
)

// 快速排序   O(N*logN) 最坏情况为0(N^2)   额外空间复杂度O(logN)
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

//arr[l...r]排好序
func process(arr []int, L, R int) {
	if L < R {
		// 随机取一个数和最后一个数交换
		rand.Seed(time.Now().UnixNano())
		rd := rand.Intn(R-L+1) + L
		arr[rd], arr[R] = arr[R], arr[rd]
		p1, p2 := partition(arr, L, R)

		process(arr, L, p1-1) // <区
		process(arr, p2+1, R) // >区
	}
}

// 处理arr[l...r]
// 默认以arr[r]做划分  arr[r]->p    <p  ==p  >p
// 返回  等于区域的左边界和右边界
func partition(arr []int, L, R int) (int, int) {
	less := L - 1  // <区右边界
	more := R      // >区左边界
	for L < more { // L表示当前数的位置
		if arr[L] < arr[R] { //当前数 < 划分值
			// 当前数和左边界后一个数交换, 左边界右移, 当前数到下一个(L++)
			arr[L], arr[less+1] = arr[less+1], arr[L]
			less++
			L++
		} else if arr[L] > arr[R] { //当前数 > 划分值
			// 当前数和右边界前一个数交换, 右边界左移, 当前数不变
			arr[L], arr[more-1] = arr[more-1], arr[L]
			more--
		} else { //当前数 = 划分值
			//当前数到下一个(L++)
			L++
		}
	}
	arr[more], arr[R] = arr[R], arr[more]
	return less + 1, more
}

func main() {
	a := []int{3, 5, 6, 7, 4, 3, 5, 8}
	fmt.Println(a)
	// sort.Ints(a)
	quickSort(a)
	fmt.Println(a)
}
