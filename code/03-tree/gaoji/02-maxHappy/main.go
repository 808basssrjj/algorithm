package main

type employee struct {
	happy  int
	subEmp []*employee
}

// 整个公司的人员结构可以看作一颗标准的,没有环的多叉树.
// 树的头节点时唯一的老板.出老板外的每个每个员工都有唯一的直接上级.
// 叶节点是没有任何下属的基层员工
// 现在要办part你可以决定哪些员工来,哪些不来
// 规则: 如果某个来了,那么这个员工的直接下级不能来
// 返回最大的快乐值

//1.思路:
// 例:x->[a,b,c]
// x员工来时   最大快乐值为:x乐 + a子树,a不来时的最大快乐值 + b子树,b不来时的最大快乐值 + ...
// x员工不来时 最大快乐值为:0 + max(a来,a不来) + max(b来,b不来) + ...
//2.实现:
// 需向左右子树,要两个信息: 子树头节点来时的最大值,和不来时的最大值

func maxHappy(head *employee) int {
	return max(process(head))
}

// process 返回x来和x不来的最大快乐值
func process(x *employee) (come, notCome int) {
	if len(x.subEmp) == 0 {
		return x.happy, 0
	}
	come, notCome = x.happy, 0
	for i := 0; i < len(x.subEmp); i++ {
		subCome, subNotCome := process(x.subEmp[i])
		come += subNotCome                 //x员工来时  最大快乐值为:x乐 + a子树,a不来时的最大快乐值 + b子树,b不来时的最大快乐值 + ...
		notCome += max(subCome, subNotCome) //x员工不来时 最大快乐值为:0 + max(a来,a不来) + max(b来,b不来) + ...
	}
	return
}

func main() {

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
