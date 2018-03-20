package main

import(
	"fmt"
)

type BTreeNode struct {
	value    int
	p_left  *BTreeNode
	p_right *BTreeNode
}

var pHead *BTreeNode
var pListIndex *BTreeNode

// 创建二叉树
func addBTreeNode(pCurrent *BTreeNode, value int) {
	if nil == pCurrent {
		pBTree := &BTreeNode{}
		pBTree.p_left = nil
		pBTree.p_right = nil
		pBTree.value = value
		pCurrent = pBTree
	} else {
		if pCurrent.value > value {
			addBTreeNode(pCurrent.p_left, value)
		} else if pCurrent.value < value {
			addBTreeNode(pCurrent.p_right, value)
		} else {
			fmt.Println("重复加入节点:", value)
		}
	}
}

// 中序遍历二叉树
func  MiddleOrderErgodicBTree(pCurrent *BTreeNode) {
	if nil == pCurrent {
		return
	}

	if nil != pCurrent.p_left {
		MiddleOrderErgodicBTree(pCurrent.p_left)
	}

	convertToDoubleList(pCurrent)

	if nil != pCurrent.p_right {
		MiddleOrderErgodicBTree(pCurrent.p_right)
	}
}

// 二叉树转换成List

func convertToDoubleList(pCurrent *BTreeNode) {
	pCurrent.p_left = pListIndex
	if nil != pListIndex {
		pListIndex.p_right = pCurrent
	} else {
		pHead = pCurrent
	}

	pListIndex = pCurrent

	fmt.Println(pCurrent.value)
}

func main() {
	var pRoot *BTreeNode
	addBTreeNode(pRoot, 10)
	addBTreeNode(pRoot, 4)
	addBTreeNode(pRoot, 6)
	addBTreeNode(pRoot, 8)
	addBTreeNode(pRoot, 12)
	addBTreeNode(pRoot, 14)
	addBTreeNode(pRoot, 15)
	addBTreeNode(pRoot, 16)
	MiddleOrderErgodicBTree(pRoot)
	return
}






























