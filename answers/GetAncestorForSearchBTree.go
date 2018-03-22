package main 

import (
	"fmt"
)
//搜索二叉树，根据搜索二叉树的性质，左子树的所有节点比根节点小，右子树的所有节点比跟节点大。
type BTree struct {
	Data int
	LeftChild *BTree
	RightChild *BTree
}

func GetAncestor(bTree, a, b *BTree) *BTree{
	if a.Data < bTree.Data && b.Data < bTree.Data {
		return GetAncestor(bTree.LeftChild, a, b)
	} else if a.Data > bTree.Data && b.Data > bTree.Data {
		return GetAncestor(bTree.RightChild, a, b)
	} else {
		return bTree
	}
}

func main() {
	

}