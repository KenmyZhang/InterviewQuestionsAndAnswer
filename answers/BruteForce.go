package main

import (
	"fmt"
)

func main() {
	s := "anmincd"
	p := "mincd"
	fmt.Println(ViolentMatch(s,p))
	s = "anmincd"
	p = "abcde"
	fmt.Println(ViolentMatch(s,p))	
}

func ViolentMatch(s, p string) int {
	sLen := len(s)
	pLen := len(p)

	i := 0
	j := 0

	for ; i < sLen && j < pLen; {
		if s[i] == p[j] {
			i++
			j++
		} else {
			i = i -j + 1
			j = 0
		}
	}

	if j == pLen {
		return i - j
	} else {
		return -1
	}
}
