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

/*
下标i           0  1  2  3  4  5  6  7  8  9
p(i)           a  b  c  d  a  a  b  c  a  b
next[i]       -1  0  0  0  0  1  1  2  3  1
优化的next[i]  -1  0  0  0  -1 1  0  0  3  0

    next[0]= -1 意义：任何串的第一个字符的模式值规定为-1。
    
    next[j]= -1 意义：模式串T中下标为j的字符，如果与首字符相同，且j的前面的1—k个字符与开头的1—k个字符不等（或者相等但T[k]==T[j]）（1≤k<j），如：T=”abCabCad” 则 next[6]=-1，因T[3]=T[6].
    
    next[j]=k  意义：模式串T中下标为j的字符，如果j的前面k个字符与开头的k个字符相等，且T[j] != T[k] （1≤k<j）即T[0]T[1]T[2]......T[k-1]==T[j-k]T[j-k+1]T[j-k+2]…T[j-1]且T[j] != T[k].（1≤k<j）;
    
    next[j]=0 意义：除（1）（2）（3）的其他情况。
*/
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
