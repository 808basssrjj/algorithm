package main

import "fmt"

//有一个正数数组coins存放当前拥有的所有硬币，数组中每一个位置代表一枚硬币，
//每一个位置的值是该硬币的面额，有可能有重复值。现在给你一个总额total，求出最少用多少硬币能够正好凑出total？
//假如当前 coins = { 2，7，3，5，3 }，total = 10；则求出最少使用 2 枚硬币（7 + 3）可以凑出total。

func coinsMin1(coins []int, total int) int {
	return process1(coins, 0, total)
}

//coinIndex: 当前硬币在数组中的下标
//restMoney: 离total还需要的钱数
//return: 当前决定下,需要的最少硬币数
func process1(coins []int, coinIndex int, restMoney int) int {
	// base case1
	// 如果该条尝试路径还有硬币没有选择，就已经凑不整total了，该条尝试路径失败
	if restMoney < 0 {
		return -1
	}

	// base case2
	// 如果该条路径正好凑整total，该条路径成功，当前不需要再选择硬币，直接返回
	if restMoney == 0 {
		return 0
	}

	// base case3
	// 如果该条尝试路径所有硬币都选择了，但还是没有凑整total，该条尝试路径失败
	//if restMoney > 0 && coinIndex == len(coins) {
	if coinIndex == len(coins) {
		return -1
	}

	// 如果该条尝试路径还有硬币没有选择，还没有凑整total，做决策
	p1 := process1(coins, coinIndex+1, restMoney)                  //选当前
	p2 := process1(coins, coinIndex+1, restMoney-coins[coinIndex]) //不选

	// 无论选不选当前硬币都不能成功凑出total
	if p1 == -1 && p2 == -1 {
		return -1
	} else {
		if p1 == -1 { //不选择当前硬币不能凑出total，选择当前硬币可以凑出total
			return p2 + 1
		} else if p2 == -1 { //不选择当前硬币可以凑出total，选择当前硬币不能凑出total
			return p1
		} else { // 不选择当前硬币和选择当前硬币都可以凑出total
			return min(p1, p2+1)
		}
	}
}

func coinsMin2(coins []int, total int) int {
	//在初始化dp表的时候，因为 -1 表示算过但是是无效解。，所以可以将dp表的所有位置初始化为 -2。
	l := len(coins)
	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, total+1)
		for j := 0; j <= total; j++ {
			dp[i][j] = -2
		}
	}

	return process2(coins, 0, total, dp)
}
func process2(coins []int, coinIndex int, restMoney int, dp [][]int) int {
	// 如果该条尝试路径还有硬币没有选择，就已经凑不整total了，该条尝试路径失败
	if restMoney < 0 {
		return -1
	}

	// 判断缓存是否命中
	if dp[coinIndex][restMoney] != -2 {
		return dp[coinIndex][restMoney]
	}

	// 如果该条路径正好凑整total，该条路径成功，当前不需要再选择硬币，直接返回
	if restMoney == 0 {
		dp[coinIndex][restMoney] = 0
	} else if restMoney > 0 && coinIndex == len(coins) {
		//如果该条尝试路径所有硬币都选择了，但还是没有凑整total，该条尝试路径失败
		dp[coinIndex][restMoney] = -1
	} else {
		// 如果该条尝试路径还有硬币没有选择，还没有凑整total，做决策
		p1 := process2(coins, coinIndex+1, restMoney, dp)                  //选当前
		p2 := process2(coins, coinIndex+1, restMoney-coins[coinIndex], dp) //不选

		// 无论选不选当前硬币都不能成功凑出total
		if p1 == -1 && p2 == -1 {
			dp[coinIndex][restMoney] = -1
		} else {
			if p1 == -1 { //不选择当前硬币不能凑出total，选择当前硬币可以凑出total
				dp[coinIndex][restMoney] = p2 + 1
			} else if p2 == -1 { //不选择当前硬币可以凑出total，选择当前硬币不能凑出total
				dp[coinIndex][restMoney] = p1
			} else { // 不选择当前硬币和选择当前硬币都可以凑出total
				dp[coinIndex][restMoney] = min(p1, p2+1)
			}
		}
	}
	return dp[coinIndex][restMoney]
}

func coinsMin3(coins []int, total int) int {
	l := len(coins)

	// row: 0 ~ coins.length
	// col: 0 ~ total
	dp := make([][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([]int, total+1)
	}

	// 将dp表已知位置的值初始化
	//for index := 0; index <= l; index++ {
	//	dp[index][0] = 0
	//}
	for rest := 1; rest <= total; rest++ {
		dp[l][rest] = -1
	}
	fmt.Println(dp)

	// 推dp表每一个位置的值，从下往上，从左往右
	for index := l - 1; index >= 0; index-- {
		for rest := 1; rest <= total; rest++ {
			// 在DP迭代的过程中已经规避了restMoney=0和coinIndex=coins.length-1的情况，这些情况的在dp表中对应的值已经被事先填上了
			p1 := dp[index+1][rest]
			p2 := -1
			if rest-coins[index] >= 0 {
				p2 = dp[index+1][rest-coins[index]]
			}
			if p1 == -1 && p2 == -1 {
				dp[index][rest] = -1
			} else {
				if p1 == -1 {
					dp[index][rest] = p2 + 1
				} else if p2 == -1 {
					dp[index][rest] = p1
				} else {
					dp[index][rest] = min(p1, p2+1)
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[0][total]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinsMin1([]int{2, 7, 3, 5, 3}, 10))
	fmt.Println(coinsMin2([]int{2, 7, 3, 5, 3}, 10))
	fmt.Println(coinsMin3([]int{2, 7, 3, 5, 3}, 10))
}
