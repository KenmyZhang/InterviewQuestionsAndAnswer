package main

//三叉链，二叉树节点有指向父节点的指针。
type BTree struct {
	Data int
	Right *BTree
	Left *BTree
	Parent *BTree
}

func GetCommonAncestor(bTree, node1, node2 *BTree) *BTree {
	var tmp *BTree
	for {
		if node1 == nil {
			break
		}
		node1 = node1.Parent
		tmp = node2
		for {
			if tmp == nil {
				break
			}
			if node1 == tmp.Parent {
				return node1
			}
			tmp = tmp.Parent
		}
	}
}