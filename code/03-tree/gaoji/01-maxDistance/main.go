package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func newNode(val int) *node {
	return &node{val: val}
}

// 从二叉树的节点A出发,可以向上或向下走,但沿途节点只经过一次,
// 到达节点B时路径上的节点个数,叫做A到B的距离,
// 那么二叉树任何两个节点都有距离,求整棵树的最大距离

//1.思路:
// 当X节点不参与计算时: 最大距离为 左子树最大距离或右子树最大距离
// 当X节点参与计算时 : 最大距离为  左子树高度+右子树高度+1
// 总最大距离为 三者较大那个
//2.实现:
// 需向左右子树,要两个信息: 子树的最大距离以及高度

func maxDistance(head *node) int {
	maxDis, _ := process(head)
	return maxDis
}

// process 返回以X开头的整棵树的,最大距离和高度
func process(x *node) (maxDis, height int) {
	if x == nil {
		return 0, 0
	}
	// 左右树信息
	lDis, lHeight := process(x.left)
	rDis, rHeight := process(x.right)

	maxDis = max(lHeight+rHeight+1, max(lDis, rDis))
	height = max(lHeight, rHeight) + 1
	return
}

func main() {
	head := newNode(1)
	head.left = newNode(2)
	head.right = newNode(3)
	head.left.left = newNode(4)
	head.left.right = newNode(5)
	head.left.left.left = newNode(6)
	fmt.Println(maxDistance(head))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
