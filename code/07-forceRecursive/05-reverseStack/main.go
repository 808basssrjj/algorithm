package main

import "fmt"

//给你一个栈，请你逆序这个栈，不能申请额外的数据结构，只能使用递归函数。

func reverseStack(stack *[]int) {
	if len(*stack) == 0 {
		return
	}
	res := f(stack)
	reverseStack(stack)
	*stack = append(*stack, res)
}

// f 删除栈底元素,并返回
func f(stack *[]int) int {
	l := len(*stack) - 1

	// 弹栈, 并保存栈顶元素
	res := (*stack)[l]
	*stack = (*stack)[:l]

	// 当栈空的时候，最后一个栈顶元素不压栈，直接返回
	if len(*stack) == 0 {
		return res
	}

	// 获取最后一个栈顶元素
	last := f(stack)
	// 将其他栈顶元素再压回栈中
	*stack = append(*stack, res)
	// 返回最后一个栈顶元素
	return last
}

func main() {
	arr := []int{1, 2, 3}
	reverseStack(&arr)
	fmt.Println(arr)
}
