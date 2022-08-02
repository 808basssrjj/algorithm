package main

import (
	"fmt"
)

// MaxWindow 获得窗口最大值(队列的头元素)的结构
// L和R移动时都要保持队列头到尾递减
type MaxWindow struct {
	R   int
	L   int
	arr []int
	ls  []int
	//双端队列存的是 arr的下标
	//双端队列维持的是如果此时不让 R 向右移动而选择让 L 依次向右移动，谁会依次成为最大值这个信息。
	//ls  list.List
}

func NewMaxWindow(arr []int) *MaxWindow {
	return &MaxWindow{
		R:   0,
		L:   -1,
		arr: arr,
		ls:  make([]int, 0, len(arr)),
		//ls:  list.List{},
	}
}

// AddFromRight R移动时
// 从尾部加入,需从尾部弹出<=加入值的数
func (w *MaxWindow) AddFromRight() {
	if w.R == len(w.arr) {
		return
	}

	cur := len(w.ls) - 1
	for len(w.ls) != 0 && w.arr[w.ls[cur]] <= w.arr[w.R] { //从尾部弹出所有 <=加入值的数
		w.ls = w.ls[:cur]
		cur--
	}
	w.ls = append(w.ls, w.R)
	w.R++
}

// RemoveFromLeft L移动时
// 如果删除的值是最大值,从队列中删除最大值
func (w *MaxWindow) RemoveFromLeft() {
	if w.L >= w.R-1 {
		return
	}

	w.L++
	if w.ls[0] == w.L {
		w.ls = w.ls[1:]
	}
}

// GetMax 获得窗口最大值
func (w *MaxWindow) GetMax() int {
	if len(w.ls) != 0 {
		return w.arr[w.ls[0]]
	}
	return 0
}

func main() {
	ar := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(solution(ar, 3))
}

// 给定一个数组 nums 和滑动窗口的大小k，请找出所有滑动窗口里的最大值。
// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
func solution(arr []int, k int) []int {
	if len(arr) < 0 {
		return nil
	}
	res := make([]int, len(arr)-k+1)

	maxW := NewMaxWindow(arr)
	for ; k > 0; k-- {
		maxW.AddFromRight()
	}

	for i := 0; i < len(res); i++ {
		res[i] = maxW.GetMax()
		maxW.AddFromRight()
		maxW.RemoveFromLeft()
	}
	return res
}

//// AddFromRight R移动时
//// 从尾部加入,需从尾部弹出<=加入值的数
//func (w *MaxWindow) AddFromRight() {
//	if w.R == len(w.arr) {
//		return
//	}
//
//	for w.ls.Len() != 0 && w.arr[w.ls.Back().Value.(int)] <= w.arr[w.R] { // 从尾部弹出<=加入值的数
//		w.ls.Remove(w.ls.Back())
//	}
//	w.ls.PushBack(w.R)
//	w.R++
//}
//
//// RemoveFromLeft L移动时
//// 如果删除的值是最大值,从队列中删除最大值
//func (w *MaxWindow) RemoveFromLeft() {
//	if w.L >= w.R-1 {
//		return
//	}
//	w.L++
//	if w.ls.Front().Value == w.L {
//		w.ls.Remove(w.ls.Front())
//	}
//}
//
//// GetMax 获得窗口最大值
//func (w *MaxWindow) GetMax() int {
//	if w.ls.Len() != 0 {
//		return w.arr[w.ls.Front().Value.(int)]
//	}
//	return 0
//}
