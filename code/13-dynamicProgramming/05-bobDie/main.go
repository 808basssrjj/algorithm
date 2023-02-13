package main

//有一个人叫Bob，他在一个规格为 M × N 的区域内行走。
//Bob一开始的位置是（a，b），
//每一次移动等概率向上、下、左、右移动一格。规定Bob只能移动 k 次，如果Bob越界，就会死亡。
//问：Bob活下来的概率是多少？

func bobDie(M, N, a, b, k int) float64 {
	aliveCount := process(M, N, a, b, k)
	return aliveCount / 1 << (2 * k) //4^k
}

func process(M, N, a, b, rest int) int {
	if a >= M || a < 0 || b >= N || b < 0 {
		return 0
	}

	if rest == 0 {
		return 1
	}

	var res int
	res += process(M, N, a-1, b, rest-1)
	res += process(M, N, a+1, b, rest-1)
	res += process(M, N, a, b-1, rest-1)
	res += process(M, N, a, b+1, rest-1)
	return res
}

func dpWay(M, N, a, b, k int) float64 {
	if a >= M || a < 0 || b >= N || b < 0 {
		return 0
	}

	dp := make([][][]int, a)
	for i := 0; i <= a; i++ {
		dp[i] = make([][]int, b)
		for j := 0; j <= b; j++ {
			dp[i][j] = make([]int, k+1)
			dp[i][j][0] = 1
		}
	}

	for step := 1; step < k+1; step-- {
		for x := 0; x < M; x++ {
			for y := 0; y < N; y++ {
				dp[x][y][step] += getValue(dp, M, N, a-1, b, step-1)
				dp[x][y][step] += getValue(dp, M, N, a-1, b, step-1)
				dp[x][y][step] += getValue(dp, M, N, a-1, b, step-1)
				dp[x][y][step] += getValue(dp, M, N, a-1, b, step-1)
			}
		}
	}

	return dp[a][b][k] / 1 << (2 * k)
}

func getValue(dp [][][]int, M, N, a, b, step int) int {
	// 越界判断
	if a < 0 || a >= M || b < 0 || b >= N {
		return 0
	}
	return dp[a][b][step]
}

func main() {

}
