package main

import "fmt"

// 规定1对应A, 2对应B, 3对应C....
// 那么一个数字字符串比如"111", 就可以转化为AAA, KA, AK
// 给定一个只有数字字符组成的字符串str, 返回有多少种转化结果

func process(arr []uint, i int) int {
	if i == len(arr) { //发现一种有效的
		return 1
	}

	cur := arr[i]
	//0无法转化, 所以无论前面选了什么,整体为0种
	if cur == 0 {
		return 0
	}
	//1可以单独转, 也可以和下一个整体转
	if cur == 1 {
		res := process(arr, i+1) //i单独转化,后续有多少种
		if i+1 < len(arr) {
			res += process(arr, i+2) //i和i+1整体转化,后续有多少种
		}
		return res
	}
	//2可以单独转, 和下一个整体<=26才可以整体转
	if cur == 2 {
		res := process(arr, i+1) //i单独转化,后续有多少种
		if i+1 < len(arr) && arr[i+1] <= 6 {
			res += process(arr, i+2) //i和i+1 整体转化,后续有多少种
		}
		return res
	}

	//3~9只能单独转
	return process(arr, i+1)
}

func main() {
	fmt.Println(process([]uint{1, 1, 1}, 0))
}
