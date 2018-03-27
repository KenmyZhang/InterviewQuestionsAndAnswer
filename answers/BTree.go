package main 

import (
	"fmt"
)

type BTree struct {
	Data string
	Left *BTree
	Right *BTree
}

func main() {
	var bTree *BTree 
	bTree = CreateBinaryTree()
	fmt.Println("Create binary tree success!")
	fmt.Println("Total Node of tree:", TotalNumOfNode(bTree))
	fmt.Println("Depth of tree:", DepthOfTree(bTree))
	fmt.Println("Leaf num::", LeafNum(bTree))
	fmt.Println("PreRange")
	PreRange(bTree)
	fmt.Println("")
	fmt.Println("MidRange")
	MidRange(bTree)
	fmt.Println("")
	fmt.Println("LastRange")
	LastRange(bTree)
	fmt.Println("")
}

func CreateBinaryTree() *BTree {
	var bTree *BTree
	var data string
 	fmt.Scan(&data)
 	if data == "#" {
 		bTree = nil
 	} else {
 		bTree = &BTree{}
 		bTree.Data = data
 		bTree.Left = CreateBinaryTree()
 		bTree.Right = CreateBinaryTree()
 	}
 	return bTree
 }

 func PreRange(bTree *BTree) {
 	if bTree != nil {
 		fmt.Print(bTree.Data)
 		PreRange(bTree.Left)
 		PreRange(bTree.Right)
 	}
 }

 func MidRange(bTree *BTree) {
 	if bTree != nil {
 		MidRange(bTree.Left)
 		fmt.Print(bTree.Data)
 		MidRange(bTree.Right)
 	}
 }

func LastRange(bTree *BTree) {
 	if bTree != nil {
 		LastRange(bTree.Left)
 		LastRange(bTree.Right)
 		fmt.Print(bTree.Data)
 	}
}

func TotalNumOfNode(bTree *BTree) int {
	if bTree == nil {
		return 0
	} else {
		return 1 + TotalNumOfNode(bTree.Left) + TotalNumOfNode(bTree.Right)
	}
}

func DepthOfTree(bTree *BTree) int {
	if bTree == nil {
		return 0
	} else {
		var dept int
		if DepthOfTree(bTree.Left) > DepthOfTree(bTree.Right) { 
		 dept = DepthOfTree(bTree.Left) + 1
		} else { 
		 dept = DepthOfTree(bTree.Right) + 1
		}
		return dept
	}
}

func LeafNum(bTree *BTree) int {
	if bTree == nil {
		return 0
	} else if bTree.Left == nil && bTree.Right == nil {
		return 1
	} else {
		return LeafNum(bTree.Left) + LeafNum(bTree.Right)
	}
}