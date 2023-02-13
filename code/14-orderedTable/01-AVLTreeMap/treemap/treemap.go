package treemap

import "golang.org/x/exp/constraints"

type TreeMap[K constraints.Ordered, V any] struct {
	size int
	root *node[K, V]
}

func NewTreeMap[K constraints.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{}
}

// findLastIndex 找到最后添加的位置
func (m *TreeMap[K, V]) findLastIndex(key K) *node[K, V] {
	pre, cur := m.root, m.root
	for cur != nil {
		pre = cur
		if key == cur.k {
			break
		} else if key < cur.k {
			cur = cur.l
		} else {
			cur = cur.r
		}
	}
	return pre
}

// findLastNoSmallIndex
func (m *TreeMap[K, V]) findLastNoSmallIndex(key K) *node[K, V] {
	var ans *node[K, V]
	cur := m.root
	for cur != nil {
		if key == cur.k {
			ans = cur
			break
		} else if key < cur.k {
			ans = cur
			cur = cur.l
		} else {
			cur = cur.r
		}
	}
	return ans
}

// findLastNoBigIndex
func (m *TreeMap[K, V]) findLastNoBigIndex(key K) *node[K, V] {
	var ans *node[K, V]
	cur := m.root
	for cur != nil {
		if key == cur.k {
			ans = cur
			break
		} else if key < cur.k {
			cur = cur.l
		} else {
			ans = cur
			cur = cur.r
		}
	}
	return ans
}

func (m *TreeMap[K, V]) Size() int {
	return m.size
}

func (m *TreeMap[K, V]) Put(key K, value V) {
	lastNode := m.findLastIndex(key)
	if lastNode != nil && key == lastNode.k {
		lastNode.v = value
	} else {
		m.size++
		m.root = addNode(m.root, key, value)
	}
}

func (m *TreeMap[K, V]) Has(key K) bool {
	lastNode := m.findLastIndex(key)
	if lastNode != nil && key == lastNode.k {
		return true
	}
	return false
}

func (m *TreeMap[K, V]) Remove(key K) {
	if m.Has(key) {
		m.size--
		m.root = deleteNode(m.root, key)
	}
}

func (m *TreeMap[K, V]) Get(key K) V {
	var res V
	lastNode := m.findLastIndex(key)
	if lastNode != nil && key == lastNode.k {
		return lastNode.v
	}
	return res
}

func (m *TreeMap[K, V]) FirstKey() K {
	cur := m.root
	for cur.l != nil {
		cur = cur.l
	}
	return cur.k
}

func (m *TreeMap[K, V]) LastKey() K {
	cur := m.root
	for cur.r != nil {
		cur = cur.r
	}
	return cur.k
}

func (m *TreeMap[K, V]) FloorKey(key K) K {
	return m.findLastNoBigIndex(key).k
}

func (m *TreeMap[K, V]) ceilKey(key K) K {
	return m.findLastNoSmallIndex(key).k
}
