package main

import "fmt"

// 汉诺塔问题
// 把圆盘按大小顺序重新摆放在另一根柱子上。并且规定，任何时候，在小圆盘上都不能放大圆盘，且在三根柱子之间一次只能移动一个圆盘

func hanno(n int) {
	if n > 0 {
		process(n, "左", "右", "中")
	}
}

// process   把1~i从from移到to
func process(i int, start, end, other string) {
	// 1.把1到i-1 从from移到other上
	// 2.把i      从from移到to上
	// 3.把1到i-1 从other移到to上
	// 当i==1 说明只剩最上面的一个圆盘
	if i == 1 { //base case
		fmt.Printf("move 1 from %s to %s\n", start, end)
	} else {
		process(i-1, start, other, end)
		fmt.Printf("move %d from %s to %s\n", i, start, end)
		process(i-1, other, end, start)
	}
}

func main() {
	n := 3
	hanno(n)
}
