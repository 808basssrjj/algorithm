package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 二叉树序列化 先序
// _ 作为结束符
// # 为空节点
func serialization(head *treeNode) string {
	if head == nil {
		return "#_"
	}
	res := strconv.Itoa(head.val) + "_"
	res += serialization(head.left)
	res += serialization(head.right)
	return res
}

// 二叉树反序列化 先序
// 根据_切割为[]string 放入队列
func deserialization(s string) *treeNode {
	values := strings.Split(s, "_")
	queue := list.New()
	for i := 0; i < len(values); i++ {
		queue.PushBack(values[i])
	}
	return process(queue)
}
func process(queue *list.List) *treeNode {
	temp := queue.Front()
	queue.Remove(temp)
	curValue := temp.Value.(string)
	if curValue == "#" {
		return nil
	}
	intVal, _ := strconv.Atoi(curValue)
	head := &treeNode{val: intVal}
	head.left = process(queue)
	head.right = process(queue)
	return head
}
func preOrderRecur(head *treeNode) {
	if head == nil {
		return
	}
	fmt.Print(head.val, "\t")
	preOrderRecur(head.left)
	preOrderRecur(head.right)
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
	s := serialization(t1)
	fmt.Println(s)
	head := deserialization(s)
	preOrderRecur(head)
}
