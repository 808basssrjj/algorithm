package main

import (
	"fmt"
)

// 给定两个有符号32位整数a,b, 返回a和b中较大的.  不能做任何比较判断

// flip 0->1 1->0
func flip(n int32) int32 {
	return n ^ 1
}

func getSign(n int32) int32 {
	return flip((n >> 31) & 1)
}

func maxNoCompare1(a, b int32) int32 {
	c := a - b        //有可能溢出
	scA := getSign(c) //a-b为非负数1 a-b负数0
	scB := flip(scA)  //scA为0,scB为1  scA为1,scB为0

	return a*scA + b*scB //if else  可以转化为互斥条件的相加
}

func maxNoCompare(a, b int32) int32 {
	c := a - b
	sa, sb, sc := getSign(a), getSign(b), getSign(c)

	diffSab := sa ^ sb       //a,b符号相同为0, 不同为1
	sameSab := flip(diffSab) //与diffSab互斥

	returnA := sameSab*sc + diffSab*sa //a,b符号相同且c为非负数,返回a  或者a,b符号不同且a为非负数
	returnB := flip(returnA)           //与returnA互斥

	return returnA*a + returnB*b
}

func main() {
	fmt.Println(maxNoCompare1(1, 2))
	fmt.Println(maxNoCompare1(-1, 2))
	fmt.Println(maxNoCompare1(-1, -2))
	fmt.Println(maxNoCompare1(1 << 31 -1, -10))

	fmt.Println(maxNoCompare(1, 2))
	fmt.Println(maxNoCompare(-1, 2))
	fmt.Println(maxNoCompare(-1, -2))
	fmt.Println(maxNoCompare(1 << 31 -1, -10))
}
