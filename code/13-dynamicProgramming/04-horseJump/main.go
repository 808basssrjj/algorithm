package main

import "fmt"

//一个象棋棋盘，有一个棋子 "马"，停在（0，0）位置。
//给定一个位置（a，b），"马" 要通过跳 k 次 "日" 去那里。
//求马从（0，0）通过跳 k 次 "日" 到达（a，b）方法数是多少。

func chess(a, b, k int) int {
	if a < 0 || b < 0 || k < 0 {
		return -1
	}
	return process(a, b, 0, 0, k)
}

func process(a, b, x, y, step int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}

	if step == 0 {
		if x == a && y == b {
			return 1
		}
		return 0
	}

	p1 := process(a, b, x-1, y-2, step-1)
	p2 := process(a, b, x-2, y-1, step-1)
	p3 := process(a, b, x-1, y+2, step-1)
	p4 := process(a, b, x-2, y+1, step-1)
	p5 := process(a, b, x+1, y+2, step-1)
	p6 := process(a, b, x+2, y+1, step-1)
	p7 := process(a, b, x+1, y-2, step-1)
	p8 := process(a, b, x+2, y-1, step-1)
	return p1 + p2 + p3 + p4 + p5 + p6 + p7 + p8
}

func main() {
	fmt.Println(chess(7, 7, 10))
	fmt.Println(dpWay(7, 7, 10))
}

func dpWay(a, b, k int) int {
	if a < 0 || b < 0 || k < 0 {
		return -1
	}

	var dp [10][9][]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			dp[i][j] = make([]int, k+1)
			//dp[i][j][0] = 0
		}
	}
	dp[a][b][0] = 1

	// 从下往上逐层计算
	for step := 1; step <= k; step++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 9; y++ {
				dp[x][y][step] += getValue(x-1, y-2, step-1, dp)
				dp[x][y][step] += getValue(x-2, y-1, step-1, dp)
				dp[x][y][step] += getValue(x-1, y+2, step-1, dp)
				dp[x][y][step] += getValue(x-2, y+1, step-1, dp)
				dp[x][y][step] += getValue(x+1, y+2, step-1, dp)
				dp[x][y][step] += getValue(x+2, y+1, step-1, dp)
				dp[x][y][step] += getValue(x+1, y-2, step-1, dp)
				dp[x][y][step] += getValue(x+2, y-1, step-1, dp)
			}
		}
	}
	return dp[0][0][k]
}
func getValue(x, y, step int, dp [10][9][]int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	return dp[x][y][step]
}
