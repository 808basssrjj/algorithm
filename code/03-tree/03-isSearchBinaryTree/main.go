package main

import (
	"algorithm/03-tree/03-isSearchBinaryTree/stack"
	"fmt"
	"math"
)

// 判断是否是搜索二叉树
// 搜索二叉树: 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

var preValue = -1

// 一.中序遍历(左头右)时 看是否是升序
func isSBT(head *treeNode) bool {
	if head == nil {
		return true
	}
	leftIsSbt := isSBT(head.left)
	if !leftIsSbt {
		return false
	}
	if head.val <= preValue {
		return false
	} else {
		preValue = head.val
	}
	return isSBT(head.right)
}

// 一
func isSBT2(head *treeNode) bool {
	if head == nil {
		return true
	}
	preValue := -1
	stk := stack.NewStack()
	for !stk.Empty() || head != nil {
		if head != nil { // 1.左边界入栈
			stk.Push(head)
			head = head.left
		} else { // 2.弹出,处理
			head = stk.Pop().(*treeNode)

			if head.val <= preValue { // 2.1如果当前值小于或等于上一个值(即左>头), 不是搜索二叉树
				return false
			} else { // 2.2如果不是上一个值更新为当前值,继续循环
				preValue = head.val
			}

			head = head.right // 3.来到右边(如果有)
		}
	}
	return true
}

// 二. 递归套路解决
// 需要条件: 左树是搜索二叉树 && 右树是搜索二叉树  && 左树最大值<自身 && 右数最小值>自身
// 左树需返回: 是否搜索二叉树 + max   右树需返回: 是否搜索二叉树 + min
func isSBT3(head *treeNode) bool {
	_, _, res, _ := process(head)
	return res
}
func process(x *treeNode) (int, int, bool, bool) {
	if x == nil { // base case
		// 最后一个参数用来判断是否返回信息
		// 因为要判断最大最小,设置为啥都会影响
		return 0, 0, true, false
	}
	leftMin, leftMax, leftIsSBT, leftHasInfo := process(x.left)
	rightMin, rightMax, rightIsSBT, rightHasInfo := process(x.right)

	min, max := x.val, x.val
	isSearch := true
	if leftHasInfo { //如果有信息
		min = int(math.Max(float64(leftMin), float64(min)))
		max = int(math.Max(float64(leftMax), float64(max)))
		// 如果左树不是搜索二叉树则false, 如果左树最大值大于等于x则false
		if !leftIsSBT || leftMax >= x.val {
			isSearch = false
		}
	}
	if rightHasInfo { //如果有信息
		min = int(math.Max(float64(rightMin), float64(min)))
		max = int(math.Max(float64(rightMax), float64(max)))
		// 果右树不是搜索二叉树则false, 如果右树最小值小于等于x则false
		if !rightIsSBT || rightMin <= x.val {
			isSearch = false
		}
	}

	//isSearch = leftIsSBT && rightIsSBT && leftMax < x.val && rightMin > x.val
	return min, max, isSearch, true
}
func main() {
	t1 := &treeNode{val: 5}
	t2 := &treeNode{val: 3}
	t3 := &treeNode{val: 7}
	t4 := &treeNode{val: 0}
	t5 := &treeNode{val: 4}
	t6 := &treeNode{val: 6}
	t7 := &treeNode{val: 8}
	t1.left, t1.right = t2, t3
	t2.left, t2.right = t4, t5
	t3.left, t3.right = t6, t7
	fmt.Println(isSBT(t1))
	fmt.Println(isSBT2(t1))
	fmt.Println(isSBT3(t1))
}
