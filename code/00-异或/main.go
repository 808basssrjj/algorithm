package main

import (
	"fmt"
)

func solution(arr []int) {
	eor := 0
	// 全部异或  得到 a^b
	for _, v := range arr {
		eor ^= v
	}

	// a!=b 所以 eor!=0 所以eor二进制至少有一位为 1 即a,b二进制至少有一位不同
	// 提取出最右边的1  (&上补码)
	rightOne := eor & (^eor + 1)

	var firstOne int
	for _, v := range arr {
		if v&rightOne == rightOne {
			firstOne ^= v
		}
	}
	otherOne := firstOne ^ eor
	fmt.Println(firstOne, otherOne)
}

func main() {
	//找出 出现次数为基数的两个数
	arr := []int{1, 2, 3, 1, 1, 2, 2, 2, 3, 3, 3, 4}
	solution(arr)
}
