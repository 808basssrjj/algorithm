package main

import "fmt"

// 在一个数组中,每一个左右比当前数小的数累加起来,叫做这个数组的小和.求一个数组的小和
// 求左侧有多少元素比当前元素小，我们可以转换为求右侧有多少元素比当前元素大 ，这样就说明对当前元素在计算小和的时候出现的次数。
func smallSum(arr []int) {
	if len(arr) < 2 {
		return
	}
	res := process(arr, 0, len(arr)-1)
	fmt.Println(res)
}

// arr[l...r]既要排好序又要求小和
func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}

	middle := l + (r-l)/2
	leftSum := process(arr, l, middle)
	rightSum := process(arr, middle+1, r)
	sum := merge(arr, l, middle, r)
	return leftSum + rightSum + sum
}

func merge(arr []int, l, m, r int) int {
	help := make([]int, 0, r-l+1)
	res := 0
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] < arr[p2] {
			//严格的是左侧小于右侧
			res += arr[p1] * (r - p2 + 1)
			help = append(help, arr[p1])
			p1++
		} else {
			help = append(help, arr[p2])
			p2++
		}
	}
	if p1 <= m {
		help = append(help, arr[p1:m+1]...)
	}
	if p1 <= r {
		help = append(help, arr[p2:r+1]...)
	}

	for i, v := range help {
		arr[l+i] = v
	}

	return res
}

func main() {
	arr := []int{1, 3, 4, 2, 5}
	smallSum(arr)

	arr2 := []int{3, 2, 4, 5, 0}
	nixv(arr2)
}

// 求逆序
func nixv(arr []int) {
	if len(arr) < 2 {
		return
	}
	res := process1(arr, 0, len(arr)-1)
	fmt.Println(res)
}

func process1(arr []int, l, r int) []string {
	if l == r {
		return nil
	}

	middle := l + (r-l)/2
	r1 := process1(arr, l, middle)
	r2 := process1(arr, middle+1, r)
	r3:=merge1(arr, l, middle, r)
	r1 = append(r1, r2[:]...)
	r1 = append(r1, r3[:]...)
	return r1
}

func merge1(arr []int, l, m, r int) []string {
	help := make([]int, 0, r-l+1)
	var res []string
	p1 := l
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] > arr[p2] {
			//严格的是左侧小于右侧
			s := fmt.Sprintf("%d-%d",arr[p1],arr[p2])
			res = append(res, s)
			help = append(help, arr[p1])
			p1++
		} else {
			help = append(help, arr[p2])
			p2++
		}
	}
	if p1 <= m {
		help = append(help, arr[p1:m+1]...)
	}
	if p1 <= r {
		help = append(help, arr[p2:r+1]...)
	}

	for i, v := range help {
		arr[l+i] = v
	}

	return res
}