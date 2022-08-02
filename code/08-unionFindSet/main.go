package main

import "fmt"

type (
	node struct {
		val interface{}
	}
	// unionFindSet 并查集
	unionFindSet struct {
		nodeMap   map[interface{}]*node
		fatherMap map[*node]*node // key某个元素  value该元素的父
		sizeMap   map[*node]int   // key某个集合的代表元素  value该集合的大小
	}
)

// NewUnionFindSet 构造函数
func NewUnionFindSet(values []interface{}) *unionFindSet {
	l := len(values)
	nodeMap := make(map[interface{}]*node, l)
	fatherMap := make(map[*node]*node, l)
	sizeMap := make(map[*node]int, l)

	//将每一个样本构建成一个Node
	for _, value := range values {
		node := &node{val: value}
		nodeMap[value] = node  //将样本与Node一一对应
		fatherMap[node] = node //初始化时先将每个节点的指针指向自身
		sizeMap[node] = 1      //初始化时每个Node都构成一个集合，每个Node都是代表Node
	}

	return &unionFindSet{
		nodeMap:   nodeMap,
		fatherMap: fatherMap,
		sizeMap:   sizeMap,
	}
}

// findHead 给定一个元素，找到集合的代表元素（扁平优化）
// 当findHead的调用次数逼近O(N)或大于时，单次时间复杂度就很快
func (s *unionFindSet) findHead(n *node) *node {
	path := make([]*node, 0, len(s.fatherMap))

	// 直到遍历到指针指向自身的代表节点位置
	for n != s.fatherMap[n] {
		path = append(path, n) //沿途所有Node压栈
		n = s.fatherMap[n]     //赋值为其父节点
	}

	// 压缩路径:防止链太长，影响性能
	for i := 0; i < len(path); i++ {
		s.fatherMap[path[i]] = n //沿途所有Node的指针直接指向代表节点
	}

	return n
}

// IsSameSet 判断两个值，是否在同一个集合
func (s *unionFindSet) IsSameSet(a, b interface{}) bool {
	nodeA, okA := s.nodeMap[a]
	nodeB, okB := s.nodeMap[b]
	if okA && okB {
		return s.findHead(nodeA) == s.findHead(nodeB)
	}
	return false
}

// Union 合并两个集合
func (s *unionFindSet) Union(a, b interface{}) {
	nodeA, okA := s.nodeMap[a]
	nodeB, okB := s.nodeMap[b]
	if okA && okB {
		fatherA := s.fatherMap[nodeA]
		fatherB := s.fatherMap[nodeB]
		if fatherA != fatherB { //代表元素不是一同个才合并
			var big, small *node
			sizeA, sizeB := s.sizeMap[fatherA], s.sizeMap[fatherB]
			if sizeA >= sizeB {
				big, small = fatherA, fatherB
			} else {
				big, small = fatherB, fatherA
			}
			// 合并:把元素少的集合顶部，挂到元素多的集合底部
			s.fatherMap[small] = big
			s.sizeMap[big] = sizeA + sizeB
			delete(s.sizeMap, small) //元素少的不再是代表元素
		}
	}
}

func main() {
	nodes := []interface{}{"a", "b", "c", "d", 1}
	ufs := NewUnionFindSet(nodes)
	fmt.Println(ufs.IsSameSet("a", 1))
	ufs.Union("a", 1)
	fmt.Println(ufs.IsSameSet("a", 1))

	ufs.Union("b", "c")
	ufs.Union("a", "b")
	fmt.Println(ufs.IsSameSet("a", "b"))
}
