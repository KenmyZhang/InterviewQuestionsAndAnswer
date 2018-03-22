package main 

import (
	"fmt"
)

type BTree struct {
	Data int
	Right *BTree
	Left *BTree
}

func JudgeNNextBTree(next, child *BTree) bool {
	if child == nil {
		return true
	}
	
	if next == nil {
		return false
	}

	if next.Data == child.Data {
		return JudgeNNextBTree(next.Left, child.Left) && JudgeNNextBTree(next.Right, child.Right)
	} else {
		return false
	}
}

func JudgeBTree(parent, child *BTree) {
	if child == nil {
		return true
	}

	if parent == nil {
		return false
	}

	if parent.Data == child.Data {
		return JudgeNNextBTree(parent, child)
	} else if JudgeBTree(parent.Left, child.Left) == true {
		return true
	} else if JudgeBTree(parent.Right, child.Right) == true {
		return true
	}
}

func main() {
		

}