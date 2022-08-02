package main

import (
	"container/list"
	"fmt"
)

// 是否是完全二叉树
// 完全二叉树:一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，
// 如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，则这棵二叉树称为完全二叉树。
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func isCompleteBt(head *treeNode) bool {
	if head == nil {
		return true
	}

	queue := list.New()
	queue.PushBack(head)
	leaf := false // 是否出现过左右孩子不全的节点
	var l, r *treeNode
	for queue.Len() != 0 {
		t := queue.Front()
		queue.Remove(t)
		cur := t.Value.(*treeNode)
		l = cur.left
		r = cur.right

		// 1.如果有仍一节点: 有右无左, 则不是
		if r != nil && l == nil {
			return false
		}
		// 2.在不违反1的情况下,如果出现不双全的节点后,  且有孩子(即不为叶节点),则不是
		if leaf && (l != nil || r != nil) {
			return false
		}

		if l != nil {
			queue.PushBack(l)
		}
		if r != nil {
			queue.PushBack(r)
		}
		// 如果出现左右不全, 则left变为true,代表出现过左右孩子不全的节点
		if l == nil || r == nil {
			leaf = true
		}
	}
	return true
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

	fmt.Println(isCompleteBt(t1))
}
