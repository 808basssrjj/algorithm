package main

import "fmt"

// TrieNode 前缀树
// 前缀树是一种用于快速检索的多叉树结构，利用字符串的公共前缀来降低查询时间，
// 核心思想是空间换时间，经常被搜索引擎用于文本词频统计。
// 优点：最大限度地减少无谓的字符串比较，查询效率高；
// 缺点：内存消耗较大；
type TrieNode struct {
	pass int           //此节点通过了多少次
	end  int           //如果是一个单词的结束 end++
	next [26]*TrieNode //a-z的路  next[0]!=nil代表有路
	//next map[int32]*TrieNode //字符多用map
}

// NewTrie 构造方法
func NewTrie() *TrieNode {
	return &TrieNode{}
}

// Insert 插入一个单词
func (t *TrieNode) Insert(word string) {
	if word == "" {
		return
	}
	t.pass++

	cur := t
	var index int32
	for _, c := range word { // 遍历单词
		index = c - 'a' // 获得字符对应那条路
		if cur.next[index] == nil {
			cur.next[index] = &TrieNode{}
		}
		cur = cur.next[index] // 跳到下一个
		cur.pass++
	}
	cur.end++ // 单词结束
}

// Search 查找单词添加过几次
func (t *TrieNode) Search(word string) int {
	if word == "" {
		return 0
	}
	cur := t
	var index int32
	for _, c := range word {
		index = c - 'a'
		if cur.next[index] == nil {
			return 0
		}
		cur = cur.next[index]
	}
	return cur.end
}

// PrefixNum 查找加入的单词有几个是word为前缀的
func (t *TrieNode) PrefixNum(word string) int {
	if word == "" {
		return 0
	}
	cur := t
	var index int32
	for _, c := range word {
		index = c - 'a'
		if cur.next[index] == nil {
			return 0
		}
		cur = cur.next[index]
	}
	return cur.pass
}

func (t *TrieNode) Delete(word string) {
	if t.Search(word) <= 0 { //加入才删除
		return
	}
	t.pass--
	cur := t
	var index int32
	for _, c := range word {
		index = c - 'a'
		cur.next[index].pass--
		if cur.next[index].pass == 0 { //释放内存
			cur.next[index] = nil
			return
		}
		cur = cur.next[index]
	}
	cur.end--
}

func main() {
	root := NewTrie()
	root.Insert("abc")
	root.Insert("abd")
	root.Insert("ab")
	root.Insert("abc")

	abc := root.Search("abc")
	fmt.Println(abc)

	num := root.PrefixNum("ab")
	fmt.Println(num)

	root.Delete("abc")
	num = root.PrefixNum("ab")
	fmt.Println(num)
}
