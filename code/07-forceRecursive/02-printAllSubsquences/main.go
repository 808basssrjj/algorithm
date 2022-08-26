package main

import "fmt"

// 打印一个字符串的全部子序列, 包括空字符串
// 例:abc  打印abc,ab,ac,a,bc,b,c,''

func printAllSubsequences(str string) {
	chars := []byte(str)
	//res := make([]byte,0)
	//process(chars, 0, res)
	process2(chars, 0)
}

// 当前来到i位置,选择要还是不要
// res是之前的选择,所形成的列表
func process(chars []byte, i int, res []byte) {
	if i == len(chars) {
		fmt.Println(string(res))
		return
	}
	//要当前字符
	resKeep := make([]byte, len(res))
	copy(resKeep, res)
	resKeep = append(resKeep, chars[i])
	process(chars, i+1, resKeep)

	//不要当前字符
	resNoKeep := make([]byte, len(res))
	copy(resNoKeep, res)
	resNoKeep = append(resNoKeep, 32)
	process(chars, i+1, resNoKeep)
}

// 当前来到i位置,选择要还是不要
// 之前的选择,所形成的结果 是chars
func process2(chars []byte, i int) {
	if i == len(chars) {
		fmt.Println(string(chars))
		return
	}
	//要当前字符
	process2(chars, i+1)

	//不要当前字符
	tmp := chars[i]
	chars[i] = 32 //当前位置改为空  32代表空格
	process2(chars, i+1)
	chars[i] = tmp //当前位置改回去, 所以str不变
}

func main() {
	printAllSubsequences("abc")
}
