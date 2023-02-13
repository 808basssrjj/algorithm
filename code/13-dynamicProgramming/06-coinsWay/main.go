package main

import "fmt"

//有一个正数无重复数组coins，该数组中每一个位置存储的是一种硬币的面值，
//每一种面值的硬币有无数个。
//现在给一个钱数total，让你使用coins中提供的硬币种类凑出来。
//请问方法数一共有多少种？

func payWays(coins []int, total int) int {
	return process(coins, 0, total)
}
func process(coins []int, index, rest int) int {
	// 所有种类的硬币都已经尝试完
	if index == len(coins) {
		if rest == 0 {
			return 1
		}
		return 0
	}

	// 如果还没有尝试玩所有种类的硬币就凑出，是一种方案
	if rest == 0 {
		return 1
	}

	res := 0
	// 第index的面值, 从一张开始往上试, 但总体总成的钱不要超过rest
	for num := 0; num*coins[index] <= rest; num++ {
		res += process(coins, index+1, rest-num*coins[index])
	}
	return res
}

//0(N * total^2)
func dpWay1(coins []int, total int) int {
	n := len(coins)
	if n <= 0 {
		return 0
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, total+1)
	}
	dp[n][0] = 1

	for x := n - 1; x >= 0; x-- {
		for y := 0; y <= total; y++ {
			way := 0
			// 枚举行为
			for num := 0; num*coins[x] <= y; num++ {
				way += dp[x+1][y-num*coins[x]]
			}
			dp[x][y] = way
		}
	}
	return dp[0][total]
}

//0(N * total)
func dpWay2(coins []int, total int) int {
	n := len(coins)
	if n <= 0 {
		return 0
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, total+1)
	}
	dp[n][0] = 1

	for x := n - 1; x >= 0; x-- {
		for y := 0; y <= total; y++ {
			//斜率优化就是：当填dp表的时候，如果有枚举行为，就观察邻近的位置能不能替代枚举行为。只和位置依赖关系有关，和原题意无关。
			//使用斜率优化版本的代码可能是完全没有逻辑的，这就是一种通过观察然后优化的一种统一的技巧。
			dp[x][y] = dp[x+1][y] //总是需要下一行的对应位置
			if y-coins[x] >= 0 {  //当前行前一个不越界 加上
				dp[x][y] += dp[x][y-coins[x]]
			}
		}
	}
	return dp[0][total]
}

func main() {
	fmt.Println(payWays([]int{3, 5, 10, 1}, 10)) //8
	fmt.Println(dpWay1([]int{3, 5, 10, 1}, 10))  //8
	fmt.Println(dpWay2([]int{3, 5, 10, 1}, 10))  //8
}
