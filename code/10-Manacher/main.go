package main

import (
	"fmt"
)

// Manacher1 分支判断版
func Manacher1(str string) int {
	if len(str) < 0 {
		return 0
	}
	chars := manacherString(str)

	res := -1 << 31
	radius := make([]int, len(chars)) // 回文半径数组
	R, C := -1, -1                    // 最右边界和中心
	for i := 0; i < len(chars); i++ {
		if i > R { // 一.i在R外部:从i开始暴力扩, R,C更新
			//半径至少为1
			radius[i] = 1
			left, right := i-1, i+1
			for left > -1 && right < len(chars) && chars[left] == chars[right] {
				radius[i]++ //半径+1
				left--
				right++
			}
			//R,C更新
			if right-1 > R {
				R = right - 1
				C = i
			}
		} else { // 二.i在R内部
			i2 := 2*C - i //i2: i关于C对称的位置
			// 左边界 = mid-radius+1
			iL := i2 - radius[i2] + 1 //i'的左边界
			L := C - radius[C] + 1    //C的左边界
			if iL > L {               // 2.1 i'回文范围在L..R内 : i的答案就是i'的答案
				//radius[i] = radius[2*C-i]
				radius[i] = radius[i2]
			} else if iL < L { // 2.2 i'回文范围有一部分在L..R外  : i的答案就是i到R的距离
				radius[i] = R - i + 1
			} else { // 2.3 i'的回文左边界压线
				//范围至少为[R',R],半径为i到R的距离
				radius[i] = R - i + 1
				left, right := i-radius[i], R+1
				//left, right := 2*i-R-1, R+1 // 也可以按中心对称求
				for ; left > -1 && right < len(chars) && chars[left] == chars[right]; left, right = left-1, right+1 {
					radius[i]++ //半径+1
				}
				//R,C更新
				if right-1 > R {
					R = right - 1
					C = i
				}
			}
		}
		if radius[i] > res {
			res = radius[i]
		}
	}
	fmt.Println(radius)
	return res - 1 //回文长度=回文半径-1(因为一开始加了字符)
}

// Manacher 改进版
func Manacher(str string) int {
	if len(str) < 0 {
		return 0
	}
	chars := manacherString(str)

	res := -1 << 31
	radius := make([]int, len(chars)) // 回文半径数组
	R, C := -1, -1                    // R:最右边界再往右一个位置,最右有效范围是R-1位置  C:R边界的中心位置
	for i := 0; i < len(chars); i++ {
		// i至少的回文区域,先给radius
		if i >= R {
			radius[i] = 1
		} else {
			radius[i] = min(radius[2*C-i], R-i)
		}
		// 然后向两边扩
		left, right := i-radius[i], i+radius[i]
		for left > -1 && right < len(chars) {
			if chars[left] == chars[right] {
				radius[i]++ //半径+1
				left, right = left-1, right+1
			} else {
				break
			}
		}
		// 更新C,R
		if i+radius[i] > R {
			R = i + radius[i]
			C = i
		}

		res = max(radius[i], res)
	}
	fmt.Println(radius)
	return res - 1 //回文长度=回文半径-1(因为一开始加了字符)
}

// manacherString 获得处理串
func manacherString(str string) []rune {
	chars := []rune(str)
	res := make([]rune, len(chars)*2+1)
	for i := 0; i < len(res); i++ {
		if i&1 == 0 {
			res[i] = '#'
		} else {
			res[i] = chars[i>>1]
		}
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
func main() {
	fmt.Println(string(manacherString("acbbcbds")))
	fmt.Println(Manacher1("acbbcbds"))
	fmt.Println(Manacher("acbbcbds"))
}
