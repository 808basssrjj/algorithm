package main

import (
	"fmt"
)

//Heap 数据结构：堆
//堆就是用数组实现的二叉树，所以它没有使用父指针或者子指针。堆根据“堆属性”来排序，“堆属性”决定了树中节点的位置。
//堆分为两种：最大堆和最小堆，两者的差别在于节点的排序方式。
//在最大堆中，父节点的值比每一个子节点的值都要大。
//在最小堆中，父节点的值比每一个子节点的值都要小。这就是所谓的“堆属性”，并且这个属性对堆中的每一个节点都成立。
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	// 如果实现的是大顶堆，修改这里的判断即可
	return h[i] < h[j]
}

// heapInsert 向堆中插入数据时，首先会把元素放到末尾
// 某个数出现在index位置上, 能否向上移动
func (h Heap) heapInsert(index int) {
	for h.less(index, (index-1)/2) {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// heapify 删除元素
// 把最末端的节点和要删除节点的位置进行交换。删除末端节点
// 某个数出现在index位置上, 能否向下移动
func (h Heap) heapify(index int) {
	left := index*2 + 1
	heapSize := len(h)

	var mini int
	for left < heapSize { // 当下方还有孩子
		// 1.选出较小的孩子
		if left+1 < heapSize && h.less(left+1, left) {
			mini = left + 1
		} else {
			mini = left
		}

		// 2. 父元素比 较小的孩子 小
		if h[index] < h[mini] {
			break
		}

		// 子元素小, 交换, 并接着向下
		h.swap(mini, index)
		index = mini
		left = index*2 + 1
	}
}

func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.heapInsert(len(*h) - 1)
}

// Remove 删除下标为X的元素, 返回改元素
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h) -1 {
		return 0, false
	}

	n := len(*h) - 1
	h.swap(i, n)
	res := (*h)[n]
	*h = (*h)[0:n]

	if n < 2 {
		return res, true
	}

	if h.less(i, (i-1)/2) {
		h.heapInsert(i)
	} else {
		h.heapify(i)
	}

	return res, true
}

func (h *Heap) Pop() int {
	x, _ := h.Remove(0)
	return x
}

func NewHeap(arr []int) Heap {
	n := len(arr)
	h := Heap(make([]int, 0, n))

	for i := 0; i < n; i++ {
		h.Push(arr[i])
	}

	return h
}

func HeapSort(arr []int) {
	heap := NewHeap(arr)

	sortedArr := make([]int, 0, len(heap))
	for len(heap) > 0 {
		sortedArr = append(sortedArr, heap.Pop())
	}

	fmt.Println(sortedArr)
}



//某个数出现在index位置上, 能否向上移动
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] { //大于父元素
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

//某个数出现在index位置上, 能否向下移动
func heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1 //左孩子下标
	var largest int
	for left < heapSize { //当下方还有孩子
		// 1.选出最大的孩子
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}

		// 2.比较父和较大孩子谁大
		if arr[largest] < arr[index] {
			largest = index
		}
		//3.父元素大,结束
		if largest == index {
			break
		}
		// 4.子元素大,接着向下
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = index*2 + 1
	}
}

// 堆排序  O(N*logN) 额外空间复杂度O(1)
func heapSort(arr []int) {
	heapSize := len(arr)
	if heapSize < 2 {
		return
	}
	// 1.先成为一个堆
	for i := 0; i < heapSize; i++ { //O(N)
		heapInsert(arr, i) //O(logN)
	}
	//for i := heapSize-1; i >= 0; i-- { //O(N)
	//	heapify(arr, i ,heapSize) //O(logN)
	//}
	// 2.不断的把第一个元素弹出,并放到最后,实现排序
	arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
	for heapSize > 0 { //O(N)
		heapify(arr, 0, heapSize)                         //O(logN)
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0] //O(1)
		heapSize--
	}
}

func main() {
	a := []int{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19}
	heapSort(a)
	fmt.Println(a)


	HeapSort([]int{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19})
}
