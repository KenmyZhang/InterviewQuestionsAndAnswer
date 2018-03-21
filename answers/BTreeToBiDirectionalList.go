package main

import (
	"fmt"
)

type BTree struct {
	Value int
	Left  *BTree
	Right *BTree
}

var pHead *BTree
var pListIndex *BTree

// 创建二叉树
func NewBTree(arr []int) *BTree {
	var t *BTree
	for _, val := range arr {
		t = Insert(t, val)
	}
	return t
}

// 插入节点
func Insert(t *BTree, val int) *BTree {
	if t == nil {
		return &BTree{val, nil, nil}
	} else if val < t.Value {
		t.Left = Insert(t.Left, val)
	} else if val > t.Value {
		t.Right = Insert(t.Right, val)
	} else {
		fmt.Println("repeat insert:", val)
	}
	return t
}

// 中序遍历二叉树
func MiddleOrderErgodicBTree(t *BTree) {
	if nil == t {
		return
	}

	if nil != t.Left {
		MiddleOrderErgodicBTree(t.Left)
	}

	convertToDoubleList(t)

	if nil != t.Right {
		MiddleOrderErgodicBTree(t.Right)
	}
}

// 二叉树转换成List
func convertToDoubleList(t *BTree) {
	if nil != pListIndex {
		pListIndex.Right = t
	} else {
		pHead = t
	}
	pListIndex = t
	fmt.Println(t.Value)
}

//遍历链表
func RangeList(pHead *BTree) {
	if pHead != nil {
		fmt.Println( pHead.Value)
		RangeList(pHead.Right)
	}
}

func main() {
	arr := []int{3, 1, 5, 9, 3, 2, 5, 11, 19}
	t := NewBTree(arr)
        fmt.Println("tree")
	MiddleOrderErgodicBTree(t)
        fmt.Println("list")
	RangeList(pHead)
	return
}
