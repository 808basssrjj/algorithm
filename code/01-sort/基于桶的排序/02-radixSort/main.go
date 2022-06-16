package main

import (
	"fmt"
	"math"
)

// 基数统计
func radixSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	rSort(arr, 0, len(arr)-1, maxBatis(arr))
}

// 返回最大值有几位
func maxBatis(arr []int) (res int) {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	for max != 0 {
		res++
		max /= 10
	}
	return
}

// arr[L...R]排序
func rSort(arr []int, L, R int, digit int) {
	const radix = 10
	i, j := 0, 0
	bucket := make([]int, R-L+1, R-L+1) // 有多少就准备多少辅助空间
	for d := 1; d <= digit; d++ {       // 有多少位就进出几次
		// count: 10个空间
		// count[0] 当前位(d位)是0的数有几个
		// count[1] 当前位(d位)是0和1的数有几个
		// count[i] 当前位(d位)是0~i的数有几个
		//count := make([]int, 0, radix)
		count := make([]int, radix, radix)

		// 1.遍历全部,依次取出d位上的数字,放入count中 (入桶)
		// count[i]为 有几个数==i
		for i := L; i <= R; i++ {
			j = getDigit(arr[i], d)
			count[j]++
		}

		// 2.把count[i]变成 有几个数<=i
		for i = 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		// 3.从右往左遍历, 取出d位上的值, 根据count[j-1]的填,填入bucket
		for i = R; i >= L; i-- {
			j = getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]-- //填完一次 当前词频减一
		}
		// 4.把bucket的值,放回arr
		//j = 0
		for i = L; i <= R; i++ {
			arr[i] = bucket[i]
			//j++
		}
	}
}

// 获得d位上的数
func getDigit(x int, d int) int {
	power := int(math.Pow(10, float64(d-1))) //10的d-1次方
	return x / power % 10
}
func main() {
	a := []int{112, 28, 26, 36, 18}
	//fmt.Println(maxBatis(a))
	radixSort(a)
	fmt.Println(a)
}
