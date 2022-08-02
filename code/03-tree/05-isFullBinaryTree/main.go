package main

import (
	"fmt"
)

// 判断是否是满二叉树
// 满二叉树:除最后一层无任何子节点外，每一层上的所有结点都有两个子结点的二叉树。
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func isFullBT(head *treeNode) bool {
	if head == nil {
		return true
	}
	nodes, height := process(head)
	//return n == int(math.Pow(2, float64(h)))-1
	return nodes == 1<<height-1
}

// 总节点数 == 2**深度 - 1
func process(node *treeNode) (nodes, depth int) {
	if node == nil { // base case
		return
	}
	lNodes, lHeight := process(node.left)
	rNodes, rHeight := process(node.right)

	nodes = lNodes + rNodes + 1
	if lHeight >= rHeight {
		depth = lHeight + 1
	} else {
		depth = rHeight + 1
	}

	return
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
	t3.left = t6
	t3.right = t7

	fmt.Println(isFullBT(t1))
}
