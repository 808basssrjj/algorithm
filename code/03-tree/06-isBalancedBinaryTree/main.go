package main

import (
	"fmt"
)

// 判断是否是平衡二叉树
// 平衡二叉树:任意节点的子树的高度差都小于等于1
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 递归套路解决: 需要条件+最终返回值+左右信息递归
// 需要条件: 左树平衡 && 右树平衡 && 左高-右高的绝对值<=1
// 左树需返回: 是否平衡 + 自身高度   右树需返回: 是否平衡 + 自身高度
func isBalanceBT(head *treeNode) bool {
	_, res := process(head)
	return res
}
func process(head *treeNode) (int, bool) {
	if head == nil { // base case
		return 0, true
	}
	leftHeight, leftIsB := process(head.left)
	rightHeight, rightIsB := process(head.right)

	var height, heightDiff int
	if leftHeight >= rightHeight {
		height = leftHeight + 1
		heightDiff = leftHeight - rightHeight
	} else {
		height = rightHeight + 1
		heightDiff = rightHeight - leftHeight
	}
	// 左右树都平衡,并且自身平衡(左右高度差不超过1)
	isB := leftIsB && rightIsB && heightDiff < 2
	return height, isB
}
func main() {
	t1 := &treeNode{val: 1}
	t2 := &treeNode{val: 2}
	t3 := &treeNode{val: 3}
	t4 := &treeNode{val: 4}
	t5 := &treeNode{val: 5}
	t6 := &treeNode{val: 6}
	t7 := &treeNode{val: 7}
	t1.left, t1.right = t2, t3
	t2.left, t2.right = t4, t5
	t3.left, t3.right = t6, t7

	fmt.Println(isBalanceBT(t1))
}
