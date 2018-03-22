package main

import (
	"fmt"
)

type BTree struct {
	Data int    //节点数据
	LeftTag, RightTag bool //增加的线索标记
	LeftChild  *BTree //当LeftTag=0时LeftChild代表左节点，LeftTag=1时代表没有左节点LeftChild代表的是前驱节点
	RightChild *BTree //当RightTag=0时RightChild代表右节点，RightTag=1时代表没有右节点RightChild代表的是后继节点
}

var pre BTree

func CreateBTree(bTree *BTree) {
	var str string
        fmt.Scanf("%s", &str)
	if str == "#" {
		bTree = nil
	} else {
		bTree = &BTree{}
		bTree.Data = str
		CreateBTree(bTree.LeftChild)
		CreateBTree(bTree.RightChild)
	}
}

//为了算法的设计方便我们在线索化二叉树上再增加一个头节点。头节点的Data域为空，LeftChild指向没有线索时的根节点，将LeftTag设为0，而RightChild指向某种遍历方式中的最后一个节点，RightTag=1
func InThreading(bTree *BTree) {
	if nil != bTree {
		InThreading(bTree.LeftChild)  //左子树递归线索化
		if nil == bTree.LeftChild {
			bTree.LeftTag = 1
			bTree.LeftChild = pre
		} else {
			bTree.LeftTag = 0
		}

		if nil == pre.RightChild {
			pre.RightTag = 1
			pre.RightChild = p
		} else {
			pre.RightTag = 0
		}
		pre =  bTree
		InThreading(bTree.RightChild)
	}
}

//中序线索化
func InOrderThreading(head *BTree, bTree *BTree) {
	head = &BTree{}			//建立头节点
	head.LeftTag = 0		//头节点有左孩子，若树非空，则其左孩子是树根
	head.RightTag = 1		//头节点的右孩子指针为右线索
	head.RightChild = head		//初始化时，右指针指向自己
	if  nil == bTree {
		head.LeftChild = head	//若树为空，则左指针也指向自己
	} else {
		head.LeftChild = bTree   //头节点的左孩子指向根
		pre = head		 //pre初值指向头节点
		InThreading(pTree)       //中序线索化
		pre.RightChild = head	 //中序线索化算法结束后，pre为最右节点，pre的右线索指向头节点
		pre.RightTag = 1 
		head.RightChild = pre   //头节点的右线索指向pre
	}
}

//遍历中序线索二叉树
func InOrderTraverse_Thr(bTree *BTree) {
	p := bTree.LeftChild
	for {
		if p == bTree {
			break
		}
		//沿左孩子向下
		for{
			if p.LeftTag != 0 {
				break
			}
			p = p.LeftChild
		}
		fmt.Println("data", p.Data)  //访问其左子树为空的节点
		//沿右线索
		for {
			if p.RightTag != 1 || p.RightChild != bTree { 
				break
			}
			p = p.RightChild
			fmt.Println("data2", p.Data)
		}
		p = p.RightChild

	}
	fmt.Println("data3", p.Data)
}


