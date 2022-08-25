package main

import (
	"fmt"
	"sort"
)

// 贪心算法:
// 在某一个标准下,优先考虑最满足标准的样本
// 最后考虑最不满足标准的样本,最终得到一个答案

// 一些项目要占用一个会议室,会议室不能同事容纳两个项目
// 给你每一个项目开始和结束时间.
// 返回能安排最多的宣讲场次
type meet struct {
	start int
	end   int
}

func bestArrange(ms []meet, timePoint int) int {
	// 按结束时间升序
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].end < ms[j].end
	})

	var res int
	for i := 0; i < len(ms); i++ {
		if timePoint <= ms[i].start {
			res++
			timePoint = ms[i].end
		}
	}
	return res
}

func main() {
	ms := []meet{
		{6, 8},
		{7, 11},
		{9, 10},
		{11, 13},
		{12, 16},
		{15, 18},
	}

	r := bestArrange(ms, 6)
	fmt.Println(r)
}
