package treemap

import (
	"golang.org/x/exp/constraints"
)

type node[K constraints.Ordered, V any] struct {
	k K
	v V
	h int // 节点高度
	l *node[K, V]
	r *node[K, V]
}

// dealHeight 计算节点高度
func (cur *node[K, V]) dealHeight() {
	var leftHeight, rightHeight int
	if cur.l != nil {
		leftHeight = cur.l.h
	}
	if cur.r != nil {
		rightHeight = cur.r.h
	}
	cur.h = max(leftHeight, rightHeight) + 1
}

// rightRotate 右旋
func rightRotate[K constraints.Ordered, V any](cur *node[K, V]) *node[K, V] {
	//A节点的左孩子B节点代替A节点的原位置，
	//B节点的右子树作为A节点的左子树（没有则不用管），
	//A节点的右子树和B节点的左子树保持不变。
	left := cur.l
	cur.l = left.r
	left.r = cur

	cur.dealHeight()
	left.dealHeight()
	return left
}

// leftRotate 左旋
func leftRotate[K constraints.Ordered, V any](cur *node[K, V]) *node[K, V] {
	//A节点的右孩子C节点代替A节点的原位置
	//C节点的左子树作为A节点的右子树（没有则不用管）
	//A节点的左子树和C节点的右子树保持不变。
	right := cur.r
	cur.r = right.l
	right.l = cur

	cur.dealHeight()
	right.dealHeight()
	return right
}

// maintain 自平衡操作
func maintain[K constraints.Ordered, V any](cur *node[K, V]) *node[K, V] {
	if cur == nil {
		return nil
	}

	var leftHeight, rightHeight int
	if cur.l != nil {
		leftHeight = cur.l.h
	}
	if cur.r != nil {
		rightHeight = cur.r.h
	}
	//计算当前节点左右子树深度差（左子树深度 - 右子树深度）。
	//如果深度差为2，判断是LL型还是LR型：
	//如果当前节点的左孩子的左孩子不是空，那么是LL型，进行LL型调整。
	//如果当前节点的左孩子的左孩子是空，那么是LR型，进行LR型调整。
	//如果深度差为-2，判断是RR型还是RL型：
	//如果当前节点的右孩子的右孩子不是空，那么是RR型，进行RR型调整。
	//如果当前节点的右孩子的右孩子是空，那么是RL型，进行RL型调整。
	if abs(leftHeight-rightHeight) > 1 {
		if leftHeight > rightHeight { //LL或LR
			var leftLeftHeight, leftRightHeight int
			if cur.l != nil && cur.l.l != nil {
				leftLeftHeight = cur.l.l.h
			}
			if cur.l != nil && cur.l.r != nil {
				leftRightHeight = cur.l.r.h
			}
			if leftLeftHeight >= leftRightHeight {
				//LL 对当前节点右旋
				cur = rightRotate(cur)
			} else {
				//LR 调整的大方向是：将当前节点的左孩子的右孩子调整到当前节点的位置
				//先对当前节点的左孩子做一次左旋，然后对当前节点做一次右旋即可。
				cur.l = leftRotate(cur.l)
				cur = rightRotate(cur)
			}
		} else { //RR或RL
			var rightLeftHeight, rightRightHeight int
			if cur.r != nil && cur.r.l != nil {
				rightLeftHeight = cur.r.l.h
			}
			if cur.r != nil && cur.r.r != nil {
				rightRightHeight = cur.r.r.h
			}
			if rightRightHeight >= rightLeftHeight {
				//RR 对当前节点左旋
				cur = leftRotate(cur)
			} else {
				//RL 调整的大方向是：将当前节点的左孩子的右孩子调整到当前节点的位置
				//先对当前节点的左孩子做一次左旋，然后对当前节点做一次右旋即可。
				cur.r = rightRotate(cur.r)
				cur = leftRotate(cur)
			}
		}
	}
	return cur
}

// addNode 在当前节点新增节点
func addNode[K constraints.Ordered, V any](cur *node[K, V], key K, value V) *node[K, V] {
	if cur == nil {
		return &node[K, V]{
			k: key,
			v: value,
			h: 1,
		}
	} else {
		if key < cur.k {
			cur.l = addNode(cur.l, key, value)
		} else {
			cur.r = addNode(cur.r, key, value)
		}
		cur.dealHeight()
		return maintain(cur)
	}
}

// deleteNode 在cur这棵树上，删掉key所代表的节点
// 返回cur这棵树的新头部
func deleteNode[K constraints.Ordered, V any](cur *node[K, V], key K) *node[K, V] {
	if key > cur.k {
		cur.r = deleteNode(cur.r, key)
	} else if key < cur.k {
		cur.l = deleteNode(cur.l, key)
	} else {
		if cur.l == nil && cur.r == nil {
			cur = nil
		} else if cur.l == nil && cur.r != nil {
			cur = cur.r
		} else if cur.l != nil && cur.r == nil {
			cur = cur.l
		} else {
			des := cur.r
			for des.l != nil {
				des = des.l
			}
			cur.r = deleteNode(cur.r, key)
			des.l = cur.l
			des.r = cur.r
			cur = des
		}
	}
	if cur != nil {
		cur.dealHeight()
	}
	return maintain(cur)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
