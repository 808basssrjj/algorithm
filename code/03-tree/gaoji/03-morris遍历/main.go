package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// morris 一种遍历二叉树的方式,时间复杂度O(N),额外空间复杂度O(1)
// 过程:假设当前节点为cur节点,一开始cur来到头节点位置
// 1.如果cur没有左孩子,cur向右移动 cur==cur.right
// 2.果然cur有左孩子,找到cur左子树上最右的节点mostRight
// 2.a如果mostRight右指针指向空,则让其指向cur,cur向左移动 cur==cur.left
// 2.b如果mostRight右指针指向cur自己,则让其指向空,cur向右移动 cur==cur.right
// 3.cur为空则停止循环
func morris(head *treeNode) {
	// 有左孩子的节点会访问两次
	if head == nil {
		return
	}

	cur, mostRight := head, &treeNode{}
	for cur != nil {
		if cur.left != nil {
			mostRight = cur.left // mostRight先来到cur左子树的头节点
			for mostRight.right != nil && mostRight.right != cur {
				// mostRight来到cur左子树最右节点
				mostRight = mostRight.right
			}

			if mostRight.right == nil { // 第一次来到cur
				mostRight.right = cur
				cur = cur.left
				continue
			} else { // 第二次回到cur  此时mostRight.right == cur
				mostRight.right = nil
			}
		}
		cur = cur.right
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

	//morris序: 1242513637
	morrisPre(t1) //1245367
	morrisIn(t1)  //4251637
	morrisPos(t1) //4526731
}

// morrisPre 先序遍历
// 访问一次的节点,直接打印
// 访问两次的节点,第一次时打印
func morrisPre(head *treeNode) {
	if head == nil {
		return
	}
	cur, mostRight := head, &treeNode{}
	for cur != nil {
		if cur.left != nil {
			mostRight = cur.left
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				fmt.Print(cur.val, "\t")
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		} else {
			fmt.Print(cur.val, "\t")
		}
		cur = cur.right
	}
	fmt.Println()
}

// morrisIn 中序遍历
// 访问一次的节点,直接打印
// 访问两次的节点,第二次时打印
func morrisIn(head *treeNode) {
	if head == nil {
		return
	}
	cur, mostRight := head, &treeNode{}
	for cur != nil {
		if cur.left != nil {
			mostRight = cur.left
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
			}
		}
		fmt.Print(cur.val, "\t")
		cur = cur.right
	}
	fmt.Println()
}

// morrisPos 后序遍历
// 访问两次的节点,第二次访问时:逆序打印左子树的右边界
// 最后再 :逆序打印整棵树左子树的右边界
// 逆序时如何保证空间复杂度O(1): 单链表逆序
func morrisPos(head *treeNode) {
	if head == nil {
		return
	}
	cur, mostRight := head, &treeNode{}
	for cur != nil {
		if cur.left != nil {
			mostRight = cur.left
			for mostRight.right != nil && mostRight.right != cur {
				mostRight = mostRight.right
			}
			if mostRight.right == nil {
				mostRight.right = cur
				cur = cur.left
				continue
			} else {
				mostRight.right = nil
				printEdge(cur.left) // 逆序打印左子树的右边界
			}
		}
		cur = cur.right
	}
	printEdge(head) // 最后打印整棵树左子树的右边界
	fmt.Println()
}

// printEdge 逆序打印node节点左子树的右边界
func printEdge(node *treeNode) {
	tail := reverseEdge(node)
	cur := tail
	for cur != nil {
		fmt.Print(cur.val, "\t")
		cur = cur.right
	}
	reverseEdge(tail)
}

// reverseEdge 翻转
func reverseEdge(node *treeNode) *treeNode {
	cur := node
	var pre, next *treeNode
	for cur != nil {
		next = cur.right
		cur.right = pre

		pre, cur = cur, next
	}
	return pre
}
