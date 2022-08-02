package main

import (
	"fmt"
)

//位运算实现加减乘除

// Add a+b
// 两个数^, 得到两个数相加的无进位信息
// 两个数&后<<1,得到两个数相加的进位信息
// 所以两个数相加, 就等于这两个结果相加
// 这两个结果相加, 同上
// 直到进位信息为0, 最终结果就是异或的值
func Add(a, b int) int {
	//13:01101  7:00111
	//   01010    01010
	//   00000    10100
	//   10100    00000
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}

// Subtraction a-b == a+(-b)
func Subtraction(a, b int) int {
	return Add(a, negNum(b))
}

// Multiplication a*b
func Multiplication(a, b int) int {
	res := 0
	c := uint(b)
	for c != 0 {
		if c&1 != 0 {
			res = Add(res, a)
		}
		a = a << 1
		c = c >> 1 //go没有>>>
	}
	return res
}

//Division a/b
func Division(a, b int) int {
	intSize := 32 << (^uint(0) >> 63) //64位还是32位
	// ^uint(0)在32是0xFFFFFFFF 64是0xFFFFFFFFFFFFFFFF
	// ^uint(0)>>63  32是0 64是1


	res := 0
	for i := intSize; i > -1; i = Subtraction(i, 1) {
		if (a >> i) >= b { //a右移或b左移都行, 但b左移有可能溢出
			res |= 1 << i
			a = Subtraction(a, b<<i)//减掉b<<i 进if肯定能减
		}
	}
	return res
}

//negNum 相反数: 取反后+1
func negNum(n int) int {
	return Add(^n, 1)
}

func main() {
	fmt.Println(Add(-1, 10))
	fmt.Println(Add(-1, -10))
	fmt.Println(Subtraction(-10, 5))
	fmt.Println(Subtraction(-10, -5))
	fmt.Println(Multiplication(-10, 5))
	fmt.Println(Multiplication(-10, -5))
}
