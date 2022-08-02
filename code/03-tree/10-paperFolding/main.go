package main

import "fmt"

func printAllFolds(n int) {
	process(1, n, true)
}

// 递归过程,来到了某一结点
// i是节点的层数, N一共几层, down==true 凹  down==false 凸
func process(i int, n int, down bool) {
	if i > n {
		return
	}
	process(i+1, n, true)
	if down {
		fmt.Printf("%d层 凹", i)
	} else {
		fmt.Printf("%d层 凸", i)
	}
	fmt.Println()
	process(i+1, n, false)
}
func main() {
	printAllFolds(3)
}
