package main

// 题目：
//给定两个长度都为 N 的数组 weights 和 values，
//weights[ i ] 和 values[ i ]分别代表 i 号物品的重量和价值。
//给定一个正数bag，表示一个载重为 bag 的袋子，你装的物品不能超过这个重量。
//你能装下物品的总价值最大是多少?
//分析：
//总左往右尝试，0号货要或不要，1号货要或不要 ...
//

func knapsackProblem(weights, values []int, bag int) int {
	return process(weights, values, bag, 0, 0)
}

// curWeight 之前所做的决策袋子的重量
// i 当前第i号物品
func process(weights, values []int, bag, curWeight, i int) int {
	// 如果所有物品尝试完
	if i == len(weights) {
		return 0
	}

	// 如果当前袋子超重
	if curWeight > bag {
		return 0
	}

	choose := values[i] + process(weights, values, bag, curWeight+weights[i], i+1)
	notChoose := process(weights, values, bag, curWeight, i+1)
	// 放入第i号物品和不放入产生的价值大的返回
	return max(choose, notChoose)
}

func process2(weights, values []int, bag, curWeight, curValue, i int) int {
	// 如果所有物品尝试完
	if i == len(weights) {
		return curValue
	}

	// 如果当前袋子超重
	if curWeight > bag {
		return 0
	}

	choose := process2(weights, values, bag, curWeight+weights[i], curValue+values[i], i+1)
	notChoose := process2(weights, values, bag, curWeight, curValue, i+1)
	// 放入第i号物品和不放入产生的价值大的返回
	return max(choose, notChoose)
}

func main() {

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
