package main

import "fmt"

func is2Power(n int) bool {
	return (n & (n-1)) == 0
}

func is4Power(n int) bool {
	//:  010101..01010101
	return (n & (n-1)) == 0 && (n & 0x55555555) != 0
}

func main() {
	fmt.Println(is4Power(4))
	fmt.Println(is4Power(16))
	fmt.Println(is4Power(15))
	fmt.Printf("%b\n", 0x55555555)
}
