package main

import (
	"algorithm/12-MonotonousStack/stack"
	"fmt"
)

// 单调栈 解决的问题:
// 我想知道每一个元素左边比该元素大且离得最近的元素和右边比该元素大且离得最近的元素都是什么?
// 单调栈要保证从栈底到栈顶存储的下标对应的元素 单调递减

type record struct {
	left  int
	right int
}

func GetRecord(arr []int) []*record {
	sk := stack.NewStack()
	res := make([]*record, len(arr))

	sk.Push([]int{0})
	for i := 1; i < len(arr); i++ {
		lastList := sk.Peek().([]int)
		lastVal := lastList[len(lastList)-1]
		if sk.IsEmpty() || arr[i] < arr[lastVal] {
			//如果栈空或者arr[i]小于栈顶下标链表对应的元素的值，直接压栈
			sk.Push([]int{i})
		} else if arr[i] == arr[lastVal] {
			//如果arr[i]等于栈顶下标链表对应的元素的值，连接到栈顶下标链表的末尾
			curArr := sk.Pop().([]int)
			curArr = append(curArr, i)
			sk.Push(curArr)
		} else {
			//如果arr[i]大于栈顶下标链表对应的元素的值，栈顶下标链表弹栈
			//直到arr[i]小于等于栈顶下标链表对应的元素值为止
			//for arr[i] > arr[lastVal] && !sk.IsEmpty() {
			for arr[i] > arr[lastVal] {
				curArr := sk.Pop().([]int)

				var left int
				if sk.IsEmpty() { //是否是栈底下标链表
					left = -1
					lastVal = i //使循环break掉
				} else {
					lastList = sk.Peek().([]int)
					lastVal = lastList[len(lastList)-1]
					left = lastVal
				}
				for _, item := range curArr {
					res[item] = &record{left: left, right: i}
				}
			}

			//加入当前遍历的值
			if sk.IsEmpty() || arr[i] < arr[lastVal] {
				sk.Push([]int{i})
			} else if arr[i] == arr[lastVal] {
				curArr := sk.Pop().([]int)
				curArr = append(curArr, i)
				sk.Push(curArr)
			}
		}
	}

	//清算阶段
	for !sk.IsEmpty() {
		curArr := sk.Pop().([]int)
		left, right := -1, -1
		if !sk.IsEmpty() { //是否是栈底下标链表
			left = len(sk.Peek().([]int)) - 1
		}
		for _, item := range curArr {
			res[item] = &record{left: left, right: right}
		}
	}

	return res
}

func GetRecord1(arr []int) []*record {
	stack := make([][]int, 0, len(arr))
	res := make([]*record, len(arr))

	stack = append(stack, []int{0})
	for i := 1; i < len(arr); i++ {
		l := len(stack)
		last := len(stack[l-1])
		lastVal := stack[l-1][last-1]
		if l == 0 || arr[i] < arr[lastVal] {
			//1.如果栈空或者arr[i]小于栈顶下标链表对应的元素的值，直接压栈
			stack = append(stack, []int{i})
		} else if arr[i] == arr[lastVal] {
			//2.如果arr[i]等于栈顶下标链表对应的元素的值，连接到栈顶下标链表的末尾
			stack[l-1] = append(stack[l-1], i)
		} else {
			//3.如果arr[i]大于栈顶下标链表对应的元素的值，栈顶下标链表弹栈
			//直到arr[i]小于等于栈顶下标链表对应的元素值为止
			//for arr[i] > arr[lastVal] && l != 0 {
			for arr[i] > arr[lastVal] {
				curArr := stack[l-1]
				stack = stack[:l-1]
				l = len(stack)

				var left int
				if l == 0 { //是否是栈底下标链表
					left = -1
					lastVal = i //使循环break掉
				} else {
					last = len(stack[l-1])
					lastVal = stack[l-1][last-1]
					left = lastVal
				}

				for _, item := range curArr {
					res[item] = &record{left: left, right: i}
				}
			}

			//加入当前遍历的值
			if len(stack) == 0 || arr[i] < arr[lastVal] {
				stack = append(stack, []int{i})
			} else if arr[i] == arr[lastVal] {
				stack[l-1] = append(stack[l-1], i)
			}
		}
	}

	//清算阶段
	for l := len(stack); l != 0; {
		curArr := stack[l-1]
		stack = stack[:l-1]
		l = len(stack)

		right := -1
		var left int
		if l == 0 { //是否是栈底下标链表
			left = -1
		} else {
			last := len(stack[l-1])
			left = stack[l-1][last-1]
		}
		for _, item := range curArr {
			res[item] = &record{left: left, right: right}
		}
	}

	return res
}

func main() {
	arr := []int{5, 4, 3, 4, 5, 3, 5, 6}
	res := GetRecord(arr)
	for _, re := range res {
		fmt.Printf("left:%d right:%d\n", re.left, re.right)
	}

	fmt.Println("aaa")
	res1 := GetRecord1(arr)
	for _, re := range res1 {
		fmt.Printf("left:%d right:%d\n", re.left, re.right)
	}
}
