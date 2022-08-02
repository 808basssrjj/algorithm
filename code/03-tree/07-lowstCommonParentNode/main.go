package main

import "fmt"

// 给定两个节点,返回两节点的最低公共祖先
type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// 用map保存孩子节点对应的父节点。
// 然后，从给定的两个节点中的任意一个开始向上回溯父节点，将路径上的所有节点都放入set中，
// 然后，再回溯另一个，每遍历到一个父节点就在set中查找是否存在，返回找到的第一个节点，就是最小公共父节点。
func lcpNode1(head, n1, n2 *treeNode) *treeNode {
	fatherMap := make(map[*treeNode]*treeNode)
	fatherMap[head] = head
	putMap(fatherMap, head)
	set := make(map[*treeNode]struct{}) // 利用空结构体实现set
	for n1 != head {
		set[n1] = struct{}{}
		n1 = fatherMap[n1]
	}
	set[head] = struct{}{}
	for {
		_, isSet := set[n2]
		if isSet {
			return n2
		} else {
			n2 = fatherMap[n2]
		}
	}
}
func putMap(fatherMap map[*treeNode]*treeNode, node *treeNode) {
	if node == nil {
		return
	}
	if node.left != nil {
		fatherMap[node.left] = node
	}
	if node.right != nil {
		fatherMap[node.right] = node
	}
	putMap(fatherMap, node.left)
	putMap(fatherMap, node.right)
}

// 利用递归向上返回的值来返回最小公共父节点
// 情况一：n1或者n2本身是另一个的父类节点，那么，最终结果应该是这个父节点本身。
// 因此，向上返回的数据，返回它本身就可以了，如果它下面有另一个，也返回它自己本身就行，这样就会把下面的覆盖掉。
// 然后，其余节点，如果不是n1和n2，就返回null，如果接收的返回值包含n1或者n2，就继续向上传递。
// 情况二：n1和n2本身没有父子关系，他们有一个公共的父节点，那么，向上传递的信息，n1或者n2就传递他们本身，
// 但是，和上面有一点不同的地方在于，其余节点，有可能接收的返回值包含n1和n2两个，那么，其实这个节点就是最小的公共父类节点，
// 因此，如果接收的左右孩子分别为n1和n2，那么，就返回他自己本身。其余节点还是接收什么就原样向上返回。
func lcpNode2(head, n1, n2 *treeNode) *treeNode {
	// 总结:
	// 本身为n1或者n2，就向上返回自己
	// 接收两个的话，就返回自己
	// 如果接收一个，就原样照抄返回
	// 自己不是n1和n2，且没有接收，就返回null
	if head == nil || head == n1 || head == n2 { //base case
		return head
	}
	left := lcpNode2(head.left, n1, n2)
	right := lcpNode2(head.right, n1, n2)
	if left != nil && right != nil { // 出答案了
		return head
	}
	if left != nil {
		return left
	} else {
		return right
	}
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

	res := lcpNode1(t1, t4, t1)
	fmt.Println(res.val)
	res = lcpNode2(t1, t4, t2)
	fmt.Println(res.val)
}
