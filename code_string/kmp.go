package code_string

import "fmt"

func GetNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	i, j := 0, -1
	for i < len(s)-1 {
		if j != -1 {
			fmt.Println("原始对比下标：j =:", j, "i=:", i, "对比值：", string(s[i]), string(s[j]))
		}
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j // 后驱等于前驱下标记
			fmt.Println("步进：j =:", j, "i=:", i)
			fmt.Println("next:数组", next)
		} else {
			fmt.Println("回朔之前：j =:", j, "i=:", i)
			j = next[j]
			fmt.Println("回朔：j =:", j, "i=:", i)
			fmt.Println("next:数组", next)
		}
	}
	return next
}

func KMPSearch(target string, text string) int {
	i, j := 0, 0
	next := GetNext(target)
	for j < len(text) {
		if i == len(target)-1 && target[i] == text[j] {
			return j - i
		}
		if i == -1 || target[i] == text[j] {
			i++
			j++
		} else {
			i = next[i]
		}
	}
	return -1
}
