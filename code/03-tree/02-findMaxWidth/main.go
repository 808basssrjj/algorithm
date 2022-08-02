package main

import (
	"fmt"
)

// 返回二叉树的最大宽度
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 1.宽度优先遍历+使用哈希表
func findMaxWidth(head *treeNode) int {
	if head == nil {
		return 0
	}
	queue := make([]*treeNode, 0)
	queue = append(queue, head)

	levelMap := make(map[*treeNode]int) // 记录每个节点的层数
	levelMap[head] = 1

	curLevel, curLevelNum := 1, 0
	max := -1
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		curNodeLevel := levelMap[cur]
		if curNodeLevel == curLevel { // 当前节点在当前统计的层
			curLevelNum++
		} else { // 当前节点已经是下一层的节点
			if curLevelNum > max {
				max = curLevelNum
			}
			curLevel++
			curLevelNum = 1
		}

		if cur.left != nil {
			levelMap[cur.left] = curLevel + 1
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			levelMap[cur.right] = curLevel + 1
			queue = append(queue, cur.right)
		}
	}

	if curLevelNum > max {
		max = curLevelNum
	}
	return max
}

// 2.不用哈希表
func findMaxWidth2(head *treeNode) int {
	if head == nil {
		return 0
	}

	queue := make([]*treeNode, 0)
	queue = append(queue, head)

	curEnd := head
	nextEnd := &treeNode{}
	curNum := 0
	max := -1
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.left != nil {
			nextEnd = cur.left
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			nextEnd = cur.right
			queue = append(queue, cur.right)
		}

		curNum++
		if cur == curEnd { // 如果是最后一个
			if curNum > max { //1.结算max
				max = curNum
			}
			curEnd = nextEnd         //2.当前层最后一个变为下一层最后一个
			nextEnd, curNum = nil, 0 //3.变量重置
		}
	}

	return max
}

func main() {
	t1 := &treeNode{val: 1}
	t2 := &treeNode{val: 2}
	t3 := &treeNode{val: 3}
	t4 := &treeNode{val: 4}
	t5 := &treeNode{val: 5}
	t6 := &treeNode{val: 6}
	t7 := &treeNode{val: 7}
	t8 := &treeNode{val: 8}
	t1.left = t2
	t1.right = t3
	t2.left = t4
	t2.right = t5
	t3.left = t6
	t5.left = t7
	t6.right = t8

	fmt.Println(findMaxWidth(t1))
	fmt.Println(findMaxWidth2(t1))
}
