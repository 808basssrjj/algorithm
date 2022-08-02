package main

import (
	"algorithm/03-tree/01-traversingBinaryTree/stack"
	"fmt"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 递归打印1 先序遍历:头左右
func preOrderRecur(head *treeNode) {
	if head == nil {
		return
	}
	fmt.Print(head.val, "\t")
	preOrderRecur(head.left)
	preOrderRecur(head.right)
}

// 递归打印2 中序遍历:左头右
func inOrderRecur(head *treeNode) {
	if head == nil {
		return
	}
	inOrderRecur(head.left)
	fmt.Print(head.val, "\t")
	inOrderRecur(head.right)
}

// 递归打印3 后序遍历:左右头
func posOrderRecur(head *treeNode) {
	if head == nil {
		return
	}
	posOrderRecur(head.left)
	posOrderRecur(head.right)
	fmt.Print(head.val, "\t")
}

// 非递归打印1 先序遍历:头左右
func preOrderUnrecur(head *treeNode) {
	if head == nil {
		return
	}
	sk := make([]*treeNode, 0)
	sk = append(sk, head)
	for len(sk) != 0 {
		cur := sk[len(sk)-1]
		fmt.Print(cur.val, "\t")
		sk = sk[:len(sk)-1]
		if cur.right != nil {
			sk = append(sk, cur.right)
		}
		if cur.left != nil {
			sk = append(sk, cur.left)
		}
	}
	fmt.Println()
}

//非递归打印3 后序遍历:左右头
func posOrderUnrecur(head *treeNode) {
	// 1.弹出cur
	// 2.cur放入收集栈
	// 3.先左再右入栈
	// 4.重复123
	// 5.处理收集栈内容
	// 即头右左顺序 放入收集栈
	if head != nil {
		s1 := stack.NewStack()
		s2 := stack.NewStack()
		s1.Push(head)
		for !s1.Empty() {
			head = s1.Pop().(*treeNode) // 类型断言
			s2.Push(head)
			if head.left != nil {
				s1.Push(head.left)
			}
			if head.right != nil {
				s1.Push(head.right)
			}
		}
		for !s2.Empty() {
			fmt.Print(s2.Pop().(*treeNode).val, "\t")
		}
		fmt.Println()
	}
}

//非递归打印2 中序遍历:左头右
func inOrderUnrecur(head *treeNode) {
	// 1.左边界入栈
	// 2.弹出,处理
	// 3.来到右边(如果有)
	// 4.重复123
	if head != nil {
		s1 := stack.NewStack()
		for !s1.Empty() || head != nil {
			if head != nil {
				s1.Push(head)
				head = head.left
			} else {
				head = s1.Pop().(*treeNode)
				fmt.Print(head.val, "\t")
				head = head.right
			}
		}
		fmt.Println()
	}
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

	// 递归序 124442555213666377731
	// 利用递归序 根据打印时机不同,实现不同顺序打印
	preOrderRecur(t1)
	fmt.Println()
	inOrderRecur(t1)
	fmt.Println()
	posOrderRecur(t1)
	fmt.Println()
	fmt.Println("--------------------------")
	// 非递归遍历
	preOrderUnrecur(t1)
	inOrderUnrecur(t1)
	posOrderUnrecur(t1)
}
