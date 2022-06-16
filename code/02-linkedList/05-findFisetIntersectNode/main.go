package main

import (
	"fmt"
)

// 给定两个带环或不带环的链表, 返回他们的相交节点, 不相交返回nil
type node struct {
	val  int
	next *node
}

// 找到第一个入环节点,无则返回nil
// 快慢指针相遇后. 快指针回到起点,变为一步一步走
// 再次相遇,相遇位置即为入环节点
func isLoop(head *node) *node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head.next
	fast := head.next.next
	for slow != fast {
		if fast.next == nil || fast.next.next == nil {
			return nil
		}
		slow = slow.next
		fast = fast.next.next
	}
	fast = head //快指针回到起点
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

func findFirstIntersectNode(head1, head2 *node) *node {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1, loop2 := isLoop(head1), isLoop(head2)
	if loop1 == nil && loop2 == nil {
		return noLoop(head1, head2)
	}
	if loop1 != nil && loop2 != nil {
		return bothLoop(head1, head2, loop1, loop2)
	}
	return nil
}

// 一个带环一个不带环还相交的情况不存在
// 都不带环
func noLoop(head1, head2 *node) *node {
	if head1 == nil || head2 == nil {
		return nil
	}

	cur1 := head1
	cur2 := head2
	var n int
	for cur1.next != nil {
		n++
		cur1 = cur1.next
	}
	for cur2.next != nil {
		n--
		cur2 = cur2.next
	}
	if cur1 != cur2 { //不相交
		fmt.Println("无环不相交")
		return nil
	}

	// cur1为长链表, cur2为短链表
	if n > 0 {
		cur1, cur2 = head1, head2
	} else {
		cur1, cur2 = head2, head1
		n = -n
	}
	for ; n != 0; n-- { //长链表先走n步
		cur1 = cur1.next
	}

	for cur1 != cur2 { //相同找到
		cur1, cur2 = cur1.next, cur2.next
	}
	fmt.Println("无环相交")
	return cur1
}

// 都带环
func bothLoop(head1, head2, loop1, loop2 *node) *node {
	if head1 == nil || head2 == nil {
		return nil
	}
	// 第一种情况: 两个链表共用同一个入环节点  如何求相交节点?和不带环一样
	if loop1 == loop2 {
		fmt.Println("有环相交 Loop1 = loop2")
		cur1, cur2 := head1, head2
		var n int
		for cur1 != loop1 {
			n++
			cur1 = cur1.next
		}
		for cur2 != loop2 {
			n--
			cur2 = cur2.next
		}
		// cur1为长链表, cur2为短链表
		if n > 0 {
			cur1, cur2 = head1, head2
		} else {
			cur1, cur2 = head2, head1
			n = -n
		}
		for ; n != 0; n-- { //长链表先走n步
			cur1 = cur1.next
		}
		for cur1 != cur2 { //相同找到
			cur1, cur2 = cur1.next, cur2.next
		}
		return cur1
	} else {
		fmt.Println("有环相交 loop1 != loop2")
		// loop1继续往下走
		// 如果遇到loop2则是  第二种情况: 两个链表入环点不同
		// 如果没遇到loop2则是 第三种情况: 两个链表不相交
		cur := loop1.next
		for cur != loop1 {
			if cur == loop2 {
				return loop1
			}
			cur = cur.next
		}
		fmt.Println("有环不相交")
		return nil
	}
}

func main() {
	// head1->node1->node2->node3->node5->node6
	head1 := &node{val: 1}
	node1 := &node{val: 1}
	node2 := &node{val: 1}
	node3 := &node{val: 1}
	node4 := &node{val: 2}
	node5 := &node{val: 3}
	node6 := &node{val: 3}
	head1.next = node1
	node1.next = node2
	node2.next = node3
	node3.next = node5
	node5.next = node6
	head2 := &node{val: 2}
	head2.next = node4

	// 1. 无环不相交
	// head1->node1->node2->node3->node5->node6
	// head2->node4
	res := findFirstIntersectNode(head1, head2)
	fmt.Println(res)

	// 2. 无环相交
	// head1->node1->node2->node3->node5->node6
	// head2->node4->node3->node5->node6
	node4.next = node3
	res = findFirstIntersectNode(head1, head2)
	fmt.Println(res)

	// 3. 有环相交 同一入环点(node3)
	// head1->node1->node2->node3->node5->node6->node3
	// head2->node4->node3->node5-node6->node3
	node6.next = node3
	res = findFirstIntersectNode(head1, head2)
	fmt.Println(res)

	// 4. 有环相交 不一入环点(node3, node5)
	// head1->node1->node2->node3->node5->node6->node3
	// head2->node4->node5-node6->node3
	node4.next = node5
	res = findFirstIntersectNode(head1, head2)
	fmt.Println(res)
}
