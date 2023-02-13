package main

import "fmt"

//int型整数N，表示有N个位置（1~N）。
//int型整数s，表示机器人一开始停在哪个位置。
//int型整数e，表示机器人结束运动的位置。
//int型整数k，表示机器人必须走k步。
//整个模型就是一个机器人从 s 走 k 步到达 e，在走的过程中，可以向左走也可以向右走，但是不能越界和原地踏步。
//比如说机器人来到 1 位置下一步只能向右走到 2 位置；机器人来到 N 位置下一步只能向左走到 N-1 位置。
//一个机器人在只能走 k 步的情况下，从 s 到达 e ，一共有多少种走法？

// walkWays1 暴力递归版
// 整个递归调用过程可以整体看作是一棵深度为 k 的二叉树，因此时间复杂度为O(2^k)。
func walkWays1(n, s, e, k int) int {
	if s < 1 || s > n || e < 1 || e > n || n < 1 {
		return 0
	}

	return process1(n, s, e, k)
}

// N 表示有N个位置（1~N）
// e 表示机器人结束运动的位置
// rest 还有多少步数
// cur 当前位置
func process1(n, e, rest, cur int) int {
	if rest == 0 {
		if cur == e { //步数位0, 当前来到e说明有一种可行
			return 1
		}
		return 0
	}
	//来到1只能走到2
	if cur == 1 {
		return process1(n, e, rest-1, 2)
	}
	//来到N只能走到N-1
	if cur == n {
		return process1(n, e, rest-1, n-1)
	}
	//来带中间位置 既能向前又能向后
	return process1(n, e, rest-1, cur+1) + process1(n, e, rest-1, cur-1)
}

// walkWays2 记忆化搜索版本
// 记忆化搜索相当于在原来暴力递归的基础上添加缓存结构，这个缓存结构就是后面所说的dp表。
// 无后效性尝试，也就是说之前的决定不会影响后面决定的结果。这种尝试类型非常适合改成动态规划。
// dp表的规模是 k×N，dp表中的 k×N 个位置在计算时只计算一次，同一个位置在重复访问时直接返回，代价是O(1)，
// 因此记忆化搜索解决该题的时间复杂度是O(k×N)。
func walkWays2(n, s, e, k int) int {
	if s < 1 || s > n || e < 1 || e > n || n < 1 {
		return 0
	}
	//1.有几个可变参数就创建一个几维的dp表。
	//2.确定每个可变参数的变化范围，从而确定dp表的规模，也就是dp表各个维度的初始容量。
	//	一般情况下，假如一个可变参数的变化范围是 1~N，那么我们会给该参数在dp数组中的对应维度开 N+1 的初始空间，该维度的第 0 位我们不会去使用，只使用第 1 位到第 N 位的空间。
	//3.将dp表的每一个位置初始化。 一般情况下，都初始化为 -1，表示该位置所代表的递归状态没有被计算过。
	//4. 将dp表加入原递归方法的形参列表，作为参数之一。
	//5. 使用缓存结构，具体使用方法为：如果当前递归状态所对应的dp表位置的值不是 -1，缓存命中，那么直接返回dp表对应位置的值。
	//6. 建立缓存结构，具体建立方法为：如果当前递归状态所对应的dp表位置的值是 -1，缓存没有命中，将计算结果存入dp表对应位置，再继续递归展开。
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	return process2(n, s, e, k, dp)
}
func process2(n, e, rest, cur int, dp [][]int) int {
	//命中缓存
	if dp[rest][cur] != -1 {
		return dp[rest][cur]
	}

	if rest == 0 {
		if cur == e { //步数位0, 当前来到e说明有一种可行
			dp[rest][cur] = 1
		} else {
			dp[rest][cur] = 0
		}
		return dp[rest][cur]
	}

	if cur == 1 { //来到1只能走到2
		dp[rest][cur] = process2(n, e, rest-1, 2, dp)
	} else if cur == n { //来到N只能走到N-1
		dp[rest][cur] = process2(n, e, rest-1, n-1, dp)
	} else { //来带中间位置 既能向前又能向后
		dp[rest][cur] = process2(n, e, rest-1, cur+1, dp) + process2(n, e, rest-1, cur-1, dp)
	}
	return dp[rest][cur]
}

func main() {
	fmt.Println(walkWays1(5, 2, 4, 4)) //4
	fmt.Println(walkWays2(5, 2, 4, 4)) //4
	fmt.Println(walkWays3(5, 2, 4, 4)) //4
}

// walkWays3 严格表结构的动态规划版本
// 严格表结构的动态规划就不需要使用递归了，也不需要将dp表中每个位置都初始化为 -1 了，
// 我们通过找到位置与位置的依赖关系，然后规定好依赖的顺序，
// 最后使用迭代的方式将暴力递归版本中调用递归函数的代码改成dp表的替代即可。
func walkWays3(n, s, e, k int) int {
	if s < 1 || s > n || e < 1 || e > n || n < 1 {
		return 0
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	// 当cur==0时, rest全部不符合题意,为-1
	for rest := 0; rest <= k; rest++ {
		dp[rest][0] = -1
	}

	// 当rest==0时, 无论cur是多少都可以直接得到是否到达终点的判断
	// base case 的转化
	//for cur := 0; cur <= n; cur++ {
	//	dp[0][cur] = 0
	//}
	dp[0][e] = 1
	fmt.Println(dp)

	// 从上往下，从左往右根据位置依赖关系填表
	// cur==0和rest==0的情况已被规避
	for rest := 1; rest <= k; rest++ {
		for cur := 1; cur <= n; cur++ {
			if cur == 1 { //1位置时, 答案为右上角的答案
				dp[rest][cur] = dp[rest-1][2]
			} else if cur == n { //n位置时, 答案为左上角的答案
				dp[rest][cur] = dp[rest-1][n-1]
			} else { //中间位置,  答案为右上角+左上角
				dp[rest][cur] = dp[rest-1][cur-1] + dp[rest-1][cur+1]
			}
		}
	}
	fmt.Println(dp)
	return dp[k][s]
}
