package main

import "fmt"

//mater公式 :  T(N) = a * T(N/b) + O(N^d)
//log(b, a) < d  =>  O(N^d)
//log(b, a) > d  =>  O(N^log(b, a)
//log(b, a) = d  =>  O(N^d  * logN)
//归并排序  O(N*logN) 额外空间复杂度O(N)
func mergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)>>1
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) {
	//var help []int
	help := make([]int, 0, R-L+1)
	p1 := L     //左边下标
	p2 := M + 1 //右边下标
	for p1 <= M && p2 <= R {
		// 都不越界 比较大小 拷贝到help中
		if arr[p1] <= arr[p2] {
			help = append(help, arr[p1])
			p1++
		} else {
			help = append(help, arr[p2])
			p2++
		}
	}
	if p1 <= M {
		help = append(help, arr[p1:M+1]...)
	}
	if p1 <= R {
		help = append(help, arr[p2:R+1]...)
	}

	for i, v := range help {
		arr[L+i] = v
	}
}

func main() {
	arr := []int{3, 2, 1, 5, 6, 2}
	mergeSort(arr)
	fmt.Println(arr)
}
