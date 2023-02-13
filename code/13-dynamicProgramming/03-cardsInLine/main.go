package main

import "fmt"

//给定一个正整型数组arr，代表数值不同的纸牌排成一条线。
//玩家A和玩家B依次拿走每张纸牌，规定玩家A先拿，玩家B后拿，
//但是每个玩家每次只能拿走最左或最右的纸牌。玩家A和玩家B都绝顶聪明，请返回最后获胜者的分数。
//例如： arr = [1, 2, 100, 4]
//开始时，玩家A只能拿走1或4。如果开始时玩家A拿走1，则排列变为[2, 100, 4]，接下来玩家B可以拿走2或4，然后继续轮到玩家A...
//如果开始时玩家A拿走4，则排列变为[1, 2, 100]，接下来玩家B可以拿走1或100，然后继续轮到玩家A...
//玩家A作为绝顶聪明的人不会先拿4，因为拿4之后，玩家B将拿走100。
//所以玩家A会先拿1，让排列变为[2, 100, 4]，接下来玩家B不管怎么选，100都会被玩家A拿走。玩家A会获胜，分数为101。所以返回101。

func pokerGame(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// 0~arr.length-1上先拿牌和后拿牌哪个分高哪个赢
	return max(choose(arr, 0, n-1), latter(arr, 0, n-1))
}
func choose(arr []int, L, R int) int {
	// base case 如果只有一张牌，直接拿走
	if L == R {
		return arr[R]
	}

	// 选择拿左边的牌
	p1 := latter(arr, L+1, R) + arr[L]
	// 选择拿右边的牌
	p2 := latter(arr, L, R-1) + arr[R]

	// 拿牌的人一定会给自己一个更好的选择
	return max(p1, p2)
}
func latter(arr []int, L, R int) int {
	// base case 如果只有一张牌，已经被拿走了
	if L == R {
		return 0
	}

	// 对手先拿了左边的牌，轮到你拿牌
	p1 := choose(arr, L+1, R)
	// 对手先拿了右边的牌，轮到你拿牌
	p2 := choose(arr, L, R-1)

	// 先拿牌的人一定会给你一个更差的选择
	return min(p1, p2)
}

func main() {
	fmt.Println(pokerGame([]int{1, 2, 100, 4}))
	fmt.Println(pokerGame2([]int{1, 2, 100, 4}))
	fmt.Println(pokerGame3([]int{1, 2, 100, 4}))
}

// pokerGame2 记忆化搜索版
func pokerGame2(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	// 0~arr.length-1上先拿牌和后拿牌哪个分高哪个赢
	return max(choose2(arr, 0, n-1, dp), latter2(arr, 0, n-1, dp))
}
func choose2(arr []int, L, R int, dp [][]int) int {
	if L >= len(arr) || R <= 0 {
		return 0
	}

	if dp[L][R] != -1 {
		return dp[L][R]
	}

	// base case 如果只有一张牌，直接拿走
	if L == R {
		dp[L][R] = arr[L]
	} else {
		// 选择拿左边的牌
		p1 := latter2(arr, L+1, R, dp) + arr[L]
		// 选择拿右边的牌
		p2 := latter2(arr, L, R-1, dp) + arr[R]
		// 拿牌的人一定会给自己一个更好的选择
		dp[L][R] = max(p1, p2)
	}
	return dp[L][R]
}
func latter2(arr []int, L, R int, dp [][]int) int {
	if L >= len(arr) || R <= 0 {
		return 0
	}

	if dp[L][R] != -1 {
		return dp[L][R]
	}

	// base case 如果只有一张牌，已经被拿走了
	if L == R {
		dp[L][R] = 0
	} else {
		// 对手先拿了左边的牌，轮到你拿牌
		p1 := choose2(arr, L+1, R, dp)
		// 对手先拿了右边的牌，轮到你拿牌
		p2 := choose2(arr, L, R-1, dp)
		// 先拿牌的人一定会给你一个更差的选择
		dp[L][R] = min(p1, p2)
	}
	return dp[L][R]
}

// pokerGame3 动态规划版
func pokerGame3(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	dp1 := make([][]int, n)
	dp2 := make([][]int, n)
	// 初始化
	for i := 0; i < n; i++ {
		dp1[i] = make([]int, n)
		dp2[i] = make([]int, n)
		dp1[i][i] = arr[i]
		dp2[i][i] = 0
	}
	//fmt.Println(dp1)
	//fmt.Println(dp2)

	// 从下往上，从左往右遍历，在遍历计算位置依赖时就规避无效位置
	//for L := n - 2; L >= 0; L-- {
	//	for R := L + 1; R < n; R++ {
	//		dp1[L][R] = max(dp2[L+1][R]+arr[L], dp2[L][R-1]+arr[R])
	//		dp2[L][R] = min(dp1[L+1][R], dp1[L][R-1])
	//	}
	//}

	// 对角线开始位置row行 col列
	col := 1
	for col < n {
		L, R := 0, col
		for L < n && R < n {
			dp1[L][R] = max(dp2[L+1][R]+arr[L], dp2[L][R-1]+arr[R])
			dp2[L][R] = min(dp1[L+1][R], dp1[L][R-1])
			L++
			R++
		}
		col++
	}

	fmt.Println(dp1)
	fmt.Println(dp2)
	return max(dp1[0][n-1], dp2[0][n-1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
