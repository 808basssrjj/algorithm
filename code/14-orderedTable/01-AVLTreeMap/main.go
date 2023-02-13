package main

import (
	"avltree/treemap"
	"fmt"
)

func main() {
	tm := treemap.NewTreeMap[int, string]()
	fmt.Println(tm.Size())
	tm.Put(1, "a")
	tm.Put(2, "b")
	tm.Put(9, "c")
	fmt.Println(tm.Size())

	fmt.Println(tm.Get(1))
	tm.Put(1, "hello")
	fmt.Println(tm.Get(1))

	fmt.Println(tm.Has(1))
	fmt.Println(tm.Has(10))

	tm.Remove(1)
	fmt.Println(tm.Has(1))
	fmt.Println(tm.Size())
}
