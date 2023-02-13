package treemap

type node struct {
	k    int
	v    interface{}
	size int // 不同的key数量
	l    *node
	r    *node
}
