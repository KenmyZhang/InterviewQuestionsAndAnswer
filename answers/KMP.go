package main

import(
        "fmt"
)

func main() {
	next := make([]int, 30)
	str := "anmincd"
	pattern := "mincd"
	CalNext(pattern, next)
	fmt.Println(KMPSearch(str, pattern, next))
	str = "anmincd"
	pattern = "abcde"
	fmt.Println(KMPSearch(str, pattern, next))		
}

func KMPSearch(str, pattern string, next []int) int {
	i, j := 0, 0
	sLen, pLen := len(str), len(pattern)
	for ; i<sLen && j<pLen; {
		if j == -1 || str[i] == pattern[j] {
			i++
			j++
		} else {
			//如果j != -1，且当前字符匹配失败（即S[i] != P[j]），则令 i 不变，j = next[j];相当于模式串向右移动的位数为j - next[j]
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}

func CalNext(str string, next []int) {
	strLen := len(str)
	next[0] = -1
	k, j := -1, 0
	for ; j < strLen - 1; {
		if k == -1 || str[j] == str[k] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
}
