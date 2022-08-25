package main

import "fmt"

// KMP
// 先求出next数组
// 如果str1 从i到x-1位置都和 str2 0到y-1位置相对，x!=y
// 那么y向前移动到y的next值位置(最大前缀的下一个)
// 1：i......x
// 1：0......y
// 2：i...j...x
// 2：    0...y...
// 相当于把str2整体向右移动，来比较str1的x和str2的y
// 本质就是从j位置重新开始匹配
// 为什么j到x-1可以不用匹配？  因为前后缀相同
// 为什么i到j-1位置一定匹配不出来？ 假设能匹配那么x位置的next值就和事实不符
func KMP(str1, str2 string) int {
	if len(str2) < 1 || len(str2) < 1 || len(str2) > len(str1) {
		return -1
	}
	next := getNext(str2) //O(M)
	var i1, i2 int
	//O(N)
	for i1 < len(str1) && i2 < len(str2) {
		if str1[i1] == str2[i2] { //对应位置一样,继续往后
			i1++
			i2++
			//} else if i2 == 0 { //对应位置不一样,且str2来到字符串开头，不能再向前跳了
		} else if next[i2] == -1 { //对应位置不一样,且str2来到字符串开头，不能再向前跳了
			i1++
		} else { //对应位置不一样,str2不在开头, i2跳到最大前后缀的下一个位置
			i2 = next[i2]
		}
	}

	// i1越界代表没匹配到 i2越界代表匹配到了
	if i2 == len(str2) {
		return i1 - i2
	} else {
		return -1
	}
}

// getNext 获取每个字符的next值:前面字符串的(不包括整体) 前缀==后缀 的最大长度
// 求i位置的next，主要看i-1位置
// 比较 i-1字符 和 i-1 next值的字符(cn) 是否相等
// 1.如果相等 i的next等于i-1的next+1
// 2.如果不相等 cn往前跳
// 3.如果cn不能向前跳了 那i位置的next==0
// abbstabbecabbstabbex   x:9 跳一次
// abbstabbecabbstabbsx   x:4 跳两次
// abbstabbecabbstabbyx   x:0 来到0位置
func getNext(str string) []int {
	if len(str) == 1 {
		return []int{-1}
	}
	next := make([]int, len(str))
	next[0], next[1] = -1, 0 //0位置规定-1, 1位置规定0

	i := 2
	prefix := 0 // prefix既代表哪个字符和i-1字符比，也代表i-1的next值
	for i < len(str) {
		if str[i-1] == str[prefix] {
			// str2[i]的前一位和str2[i]当前最有可能的最长前缀的后一位的下标相同，说明最长前缀还能延长，需要包含str2[prefix]
			// 同时当前第i位匹配成功
			prefix += 1
			next[i] = prefix
			i++
		} else if prefix > 0 {
			// 如果str2[i]的前一位和str2[i]当前最有可能的最长前缀的后一位的下标不相同，说明最长前缀必须缩小，prefix需要向前跳
			// prefix需要跳到c[prefix]最长前缀的后一位
			// 当前第i位匹配失败，下一轮继续匹配第i位
			prefix = next[prefix]
		} else {
			// 当prefix跳到第0位时，还和第i位匹配不上，说明str2[i]没有最长前缀，置为0
			// 同时当前第i位匹配成功
			next[i] = 0
			i++
		}
	}
	return next
}

func main() {
	//fmt.Println(getNext("abbstabbecabbstabbyx"))
	fmt.Println(KMP("xxyyabbstabbecabbstabbyx", "abbstabbecabbstabbyx"))
}
