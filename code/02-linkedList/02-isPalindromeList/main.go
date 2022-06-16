package main

import (
	//"Project/02-linkedList/02-isPalindromeList/stack"
	"fmt"
)

// 判断一个单链表是否是回文
type listNode struct {
	val  int
	next *listNode
}

// 1.栈结构实现
// need n extra space
//func isPalindrome1(head *listNode) bool {
//	if head == nil || head.next == nil {
//		return false
//	}
//
//	cur := head
//	// 入栈
//	stk := stack.NewStack()
//	for cur != nil {
//		stk.Push(cur.val)
//		cur = cur.next
//	}
//	// 比较
//	for head != nil {
//		if head.val != stk.Pop() {
//			return false
//		}
//		head = head.next
//	}
//
//	return true
//}

// 2.栈+快慢指针: 只把右边部分压入栈中
// need n/2 extra space
//func isPalindrome2(head *listNode) bool {
//	if head == nil || head.next == nil {
//		return false
//	}
//
//	cur := head
//	right := head.next
//	// 找到中间位置
//	for cur.next != nil && cur.next.next != nil {
//		right = right.next  // 慢指针一次走一步 ->mid
//		cur = cur.next.next // 快指针一次走二步 ->end
//	}
//
//	stk := stack.NewStack()
//	// 右边部分入栈
//	for right != nil {
//		stk.Push(right.val)
//		right = right.next
//	}
//
//	for !stk.Empty() {
//		if head.val != stk.Pop() {
//			return false
//		}
//		head = head.next
//	}
//	return true
//}

// 3.只用快慢指针: 右边部分逆序  两边同时向中间走,并比较
// 额外空间复杂度O(1)
func isPalindrome3(head *listNode) bool {
	if head == nil || head.next == nil {
		return false
	}
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next      // -> mid
		fast = fast.next.next // -> end
	}

	// 切断成两个链表
	rStart := slow.next // 右边部分的开头
	slow.next = nil     // 中间部分的下一个指向空
	// 右边链表反转
	rEnd := reverse(rStart)

	n1 := head //左边第一个节点
	n2 := rEnd //右边最后一个节点
	res := true
	for n1 != nil && n2 != nil {
		if n1.val != n2.val {
			res = false
			break
		}
		n1, n2 = n1.next, n2.next
	}

	// 右链表变回去
	reverse(rEnd)
	// 左右链表合并
	slow.next = rStart

	return res
}

func reverse(head *listNode) *listNode {
	cur := head
	var pre *listNode
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	return pre
}
func printlnList(head *listNode) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, "\t")
	}
	fmt.Println()
}

func main() {
	head := &listNode{val: 1}
	head.next = &listNode{val: 2}
	head.next.next = &listNode{val: 3}
	head.next.next.next = &listNode{val: 3}
	head.next.next.next.next = &listNode{val: 2}
	head.next.next.next.next.next = &listNode{val: 1}

	//fmt.Println(isPalindrome1(head))
	//fmt.Println(isPalindrome2(head))
	printlnList(head)
	fmt.Println(isPalindrome3(head))
	printlnList(head)
}
