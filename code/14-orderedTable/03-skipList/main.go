package main

import "fmt"

//type SkipListNode struct {
	//K
//}

func s(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}

	return process(nums, S, 0, 0)
}

func process(nums []int, s, total, cur int) int {
	if cur == len(nums) {
		if total == s {
			return 1
		}
		return 0
	}

	return process(nums, s, total+nums[cur], cur+1) + process(nums, s, total-nums[cur], cur+1)
}

func dpWay(nums []int, S int) int {
	n := len(nums)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, S+1)
	}
	dp[n][S] = 1
	fmt.Println(dp)

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= S; j++ {
			//dp[i][j] += dp[i+1][j+nums[i]]
			//dp[i][j] += dp[i+1][j-nums[i]]
			dp[i][j] += getValue(dp, S, n, i+1, j+nums[i])
			dp[i][j] += getValue(dp, S, n, i+1, j-nums[i])
		}
	}

	fmt.Println(dp)
	return dp[0][0]
}

func getValue(dp [][]int, s, n, cur, total int) int {
	if total < 0 || total > s || cur < 0 || cur > n {
		return 0
	}
	return dp[cur][total]
}

func main() {
	fmt.Println(s([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(dpWay([]int{1, 1, 1, 1, 1}, 3))
}
