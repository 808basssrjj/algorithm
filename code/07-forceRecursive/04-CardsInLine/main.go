package main

import "fmt"

func winner(arr []int) int {
	return max(first(arr, 0, len(arr)-1), second(arr, 0, len(arr)-1))
}

func first(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}
	//返回 先选L+后续 和 先选R+后续 中最大的一个
	return max(arr[L]+second(arr, L+1, R), arr[R]+second(arr, L, R-1))
}

func second(arr []int, L, R int) int {
	if L == R {
		return 0
	}
	//返回 对方选L 和 对方选R 中最小的一个
	return min(first(arr, L+1, R), first(arr, L, R-1))
}

func main() {
	fmt.Println(winner([]int{1, 2, 100, 4})) // 101
	fmt.Println(winner([]int{1, 100, 2}))    // 100
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
