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
	CalNextVal(pattern, next)
	fmt.Println(KMPSearch(str, pattern, next))		
}
/*假设现在文本串S匹配到 i 位置，模式串P匹配到 j 位置
  如果j = -1，或者当前字符匹配成功（即S[i] == P[j]），都令i++，j++，继续匹配下一个字符；
  如果j != -1，且当前字符匹配失败（即S[i] != P[j]），则令 i 不变，j = next[j]。 此举意
  味着失配时，模式串P相对于文本串S向右移动了j - next [j] 位。换言之，当匹配失败时，模式串向
  右移动的位数为：失配字符所在位置 - 失配字符对应的next 值，即移动的实际位数为：j - next[j]，且此值大于等于1。*/
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
	j, k := 0, -1
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

//
func CalNextVal(str string, next []int) {
	strLen := len(str)
	next[0] = -1
	j, k := 0, -1
	for ; j < strLen - 1; {
		if k == -1 || str[j] == str[k] {
			k++
			j++
			if str[j] != str[k] {
				next[j] = k
			} else {
				next[j] = next[k]
			}
		} else {
			k = next[k]			
		}
	}
}
