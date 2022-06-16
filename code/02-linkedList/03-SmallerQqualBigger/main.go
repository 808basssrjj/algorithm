package main

import "fmt"

// 将单链表按某值划分左边小,中间相等,右边大的形式
type listNode struct {
	val  int
	next *listNode
}


func listPartition(head *listNode, pivot int) *listNode {
	var sHead, sTail *listNode // <区头尾
	var eHead, eTail *listNode // ==区头尾
	var bHead, bTail *listNode // >区头尾

	for head != nil {
		next := head.next
		head.next = nil
		if head.val < pivot {
			if sHead == nil {
				sHead = head
				sTail = head
			} else {
				sTail.next = head
				sTail = head
			}
		} else if head.val == pivot {
			if eHead == nil {
				eHead = head
				eTail = head
			} else {
				eTail.next = head
				eTail = head
			}
		} else {
			if bHead == nil {
				bHead = head
				bTail = head
			} else {
				bTail.next = head
				bTail = head
			}
		}
		head = next
	}

	// 小于和等于合并
	if sTail != nil { //如果有小于区
		sTail.next = eHead
		if eTail == nil { //如果没有等于区,等于区尾巴变为小于区的尾巴 (谁去连大于区的头)
			eTail = sTail
		}
	}
	// 和大于区合并
	if eTail != nil { //如果小于和等于区,不是都没有
		eTail.next = bHead
	}
	var res *listNode
	if sHead != nil {
		res = sHead
	} else {
		if eHead != nil {
			res = eHead
		} else {
			res = bHead
		}
	}
	return res
}
func printList(head *listNode) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, "\t")
	}
	fmt.Println()
}

func main() {
	head := &listNode{val: 4}
	head.next = &listNode{val: 6}
	head.next.next = &listNode{val: 8}
	head.next.next.next = &listNode{val: 5}
	head.next.next.next.next = &listNode{val: 8}
	head.next.next.next.next.next = &listNode{val: 9}
	head.next.next.next.next.next.next = &listNode{val: 5}
	head.next.next.next.next.next.next.next = &listNode{val: 6}

	printList(head)
	printList(listPartition(head, 5))
}
