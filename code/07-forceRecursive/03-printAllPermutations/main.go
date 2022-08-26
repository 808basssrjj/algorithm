package main

import "fmt"

//打印一个字符串的全部排列, 要求不出现重复的排列

func printAll(str string) []string {
	if len(str) == 0 {
		return nil
	}
	chars := []byte(str)
	res := make([]string, 0, 100)
	process(chars, 0, &res)
	return res
}

// str[i:]范围上所有的字符都可以在i位置
// str[:i]范围上是之前做的选择
// res是所有形成字符串的集合
func process(str []byte, i int, res *[]string) {
	if i == len(str) {
		*res = append(*res, string(str))
	}
	visit := make([]bool, 26)

	// 遍历当前字符后的每一个字符
	for j := i; j < len(str); j++ {
		// 分支限界  之前没选过才选(去重)
		if !visit[str[j]-'a'] {
			visit[str[j]-'a'] = true

			str[i], str[j] = str[j], str[i] // 字符和当前字符交换
			process(str, i+1, res)          // 递归到下一个字符
			str[i], str[j] = str[j], str[i] // 还原字符串
		}
	}
}

func main() {
	r := printAll("abca")
	for _, item := range r {
		fmt.Println(item)
	}
}
