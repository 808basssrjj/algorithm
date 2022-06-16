package main

import "fmt"

type listNode struct {
	val  int
	next *listNode
}

// 单链表反转
func reverseList(head *listNode) *listNode {
	//cur := head
	//var pre *listNode //上一个节点
	//var next *listNode //下一个节点
	//
	//for cur != nil {
	//	next = cur.next //保存当前节点的下节点
	//	cur.next = pre //当前节点指向上一个节点
	//	//准备下一次循环  下一循环的上节点为当前节点   下一循环的节点为当前节点的下节点
	//	pre, cur = cur, next //
	//}
	//return pre

	cur := head       // 当前节点
	var pre *listNode //上一个节点
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	return pre
}
func listLink(head *listNode) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Print( cur.val,"\t")
	}
	fmt.Println()
}

func main() {
	head := new(listNode)
	head.val = 1
	l1 := new(listNode)
	l1.val = 2
	l2 := &listNode{val: 3}
	l3 := &listNode{val: 4}
	head.next = l1
	l1.next = l2
	l2.next = l3

	listLink(head)
	now := reverseList(head)
	listLink(now)
}
