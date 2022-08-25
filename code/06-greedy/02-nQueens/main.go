package main

import (
	"fmt"
	"math"
	"time"
)

// n皇后问题  如何将n个皇后放置在n×n的棋盘上，并且使皇后彼此之间不能相互攻击。

func num1(n int) int {
	if n < 1 {
		return 0
	}
	// record[0] = 0 说明第0行放在了0列上
	record := make([]int, n)
	return process(0, n, record)
}

// 来到了row行
// record[0...row-1]表示之前的行,放了皇后的位置
// n代表一共几行
// 返回值: 摆完所有皇后,合理的摆法一共几种
func process(row int, n int, record []int) int {
	if row == n { //来到终止行说明有一种摆法可行  base case
		return 1
	}
	var res int
	// 当前在row行, 尝试row行所有的列
	for col := 0; col < n; col++ {
		// 当前row行的皇后,放在col上,会不会和之前(0...row-1)的皇后,共行共列或共斜线
		if isValid(record, row, col) {
			// 深度优先
			record[row] = col
			res += process(row+1, n, record)
		}
	}
	return res
}

// record[0..row-1]需要判断,后面的无需判断
func isValid(record []int, row, col int) bool {
	for k := 0; k < row; k++ { //k行的某个皇后
		if col == record[k] || math.Abs(float64(record[k]-col)) == math.Abs(float64(k-row)) {
			// 横坐标相减绝对值==纵坐标相减绝对值 则再一条斜线上
			return false
		}
	}
	return true
}
func main() {
	//fmt.Println(num1(1)) // 1
	//fmt.Println(num1(3)) // 0
	//fmt.Println(num1(8)) // 92

	start := time.Now()
	fmt.Println(num1(13))
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	start1 := time.Now()
	fmt.Println(num2(13))
	elapsed1 := time.Since(start1)
	fmt.Println(elapsed1)
}

// 位运算加速
func num2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}
	var limit int
	if n == 32 {
		limit = -1
	} else {
		limit = (1 << n) - 1 //后n位为1,其余为0  -1全是1
	}
	return process2(limit, 0, 0, 0)
}

//colLim 列的限制, 1的位置不能放皇后
//leftDiaLim 左斜线的限制, 1的位置不能放皇后
//rightDiaLim 右斜线的限制, 1的位置不能放皇后
//例:  c==00001000 则 l==c<<1==00010000 r==c>>1==00000100
func process2(limit int, colLim, leftDiaLim, rightDiaLim int) (res int) {
	if colLim == limit { //base case  每一位都为1了说明每个位置都放好了
		return 1
	}
	// colLim|leftDiaLim|rightDiaLim:总限制  此时0可以放,1不可以放
	// 取反后: 变为1可以放,0不可以放   (为了后面更好的取出可以放的位置)
	// &limit 是为了去掉取反后 高位的1
	pos := limit & (^(colLim | leftDiaLim | rightDiaLim))
	var mostRightOne int // 最右边的1(即可以放皇后的位置)
	for pos != 0 {
		mostRightOne = pos & (^pos + 1) //提取最右边的1
		//pos = pos - mostRightOne //此位置删除, 不能放
		pos &= pos - 1 // 删除最低位的1
		c := colLim | mostRightOne
		l := (leftDiaLim | mostRightOne) << 1
		r := (rightDiaLim | mostRightOne) >> 1
		res += process2(limit, c, l, r)
	}
	return
}
