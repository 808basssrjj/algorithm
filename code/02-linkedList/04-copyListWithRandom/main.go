package main

import "fmt"

// rand指针可能只想链表中任意一个节点, 也可能指向nil
// 给定一个head节点, 完成链表的复制,返回新链表的头节点
type node struct {
	val  int
	next *node
	rand *node
}

// 1.利用哈希表 extra space O(n)
func copyRandomList1(head *node) *node {
	if head == nil {
		return nil
	}
	mapList := make(map[*node]*node)
	// 1)存入map
	cur := head
	for cur != nil { // map中: key存老节点 value存新节点
		newNode := &node{val: cur.val}
		mapList[cur] = newNode
		cur = cur.next
	}
	// 2)设置克隆节点的后继和随机节点
	cur = head
	for cur != nil {
		// cur为老节点
		// mapList[cur] 为新节点
		mapList[cur].next, mapList[cur].rand = mapList[cur.next], mapList[cur.rand]
		cur = cur.next
	}
	return mapList[head]
}

// 2.extra space O(1)
func copyRandomList2(head *node) *node {
	if head == nil {
		return nil
	}
	// 1)克隆节点 变成这样 1--1'--2--2'--3--3'
	cur := head
	next := new(node)
	for cur != nil {
		next = cur.next // 保存下一个要遍历的节点
		newNode := &node{val: cur.val}
		cur.next = newNode
		newNode.next = next
		cur = next // 遍历下一个
	}
	// 2)设置克隆链表的随机节点 (一对一对取)
	cur = head
	curCopy := new(node)
	for cur != nil {
		next = cur.next.next // 保存下一个要遍历的节点
		if cur.rand != nil {
			cur.next.rand = cur.rand.next //克隆节点的随机节点 == 当前节点的随机节点 的下个节点
		}
		cur = next
	}

	// 3)分离两个链表
	res := head.next
	cur = head
	for cur != nil {
		next = cur.next.next //当前节点的下一个节点
		curCopy = cur.next   //当前节点的克隆节点
		cur.next = next      // 老链表回去
		if next != nil {     // 新链表回去
			curCopy.next = next.next //克隆节点的下个节点 == 当前节点的下个节点 的下个节点
		}
		cur = next // 遍历下一个
	}
	return res
}

func printList(head *node) {
	fmt.Printf("%p\n", head)
	for cur := head; cur != nil; cur = cur.next {
		if cur.rand != nil {
			fmt.Printf("%d(%d)\t", cur.val, cur.rand.val)
		} else {
			fmt.Printf("%d(nil)\t", cur.val)
		}
	}
	fmt.Println()
}
func main() {
	head := &node{val: 1}
	node1 := &node{val: 2}
	node2 := &node{val: 3}
	node3 := &node{val: 4}
	node4 := &node{val: 5}
	head.next = node1
	head.rand = node2
	node1.next = node2
	node1.rand = node4
	node2.next = node3
	node2.rand = nil
	node3.next = node4
	node3.rand = head

	printList(head)
	copyList := copyRandomList1(head)
	printList(copyList)
	copyList2 := copyRandomList2(head)
	printList(copyList2)
}
