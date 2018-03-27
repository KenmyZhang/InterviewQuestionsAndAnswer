package main

import (
	"fmt"
)

type BTree struct {
	Data string    //节点数据
	LeftTag int  //增加的线索标记
	RightTag int //增加的线索标记
	LeftChild  *BTree //当LeftTag=0时LeftChild代表左节点，LeftTag=1时代表没有左节点LeftChild代表的是前驱节点
	RightChild *BTree //当RightTag=0时RightChild代表右节点，RightTag=1时代表没有右节点RightChild代表的是后继节点
}

var pre *BTree

func main() {
	t := CreateBinaryTree()
	h := InOrderThreading(t)
	InOrderTraverse_Thr(h)
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
 		bTree.LeftChild = CreateBinaryTree()
 		bTree.RightChild = CreateBinaryTree()
 	}
 	return bTree
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
			pre.RightChild = bTree
		} else {
			pre.RightTag = 0
		}
		pre =  bTree
		InThreading(bTree.RightChild)
	}
}

//中序线索化
func InOrderThreading(bTree *BTree) *BTree{
	head := &BTree{}			//建立头节点
	head.LeftTag = 0		//头节点有左孩子，若树非空，则其左孩子是树根
	head.RightTag = 1		//头节点的右孩子指针为右线索
	head.RightChild = head		//初始化时，右指针指向自己
	if  nil == bTree {
		head.LeftChild = head	//若树为空，则左指针也指向自己
	} else {
		head.LeftChild = bTree   //头节点的左孩子指向根
		pre = head		 //pre初值指向头节点
		InThreading(bTree)       //中序线索化
		pre.RightChild = head	 //中序线索化算法结束后，pre为最右节点，pre的右线索指向头节点
		pre.RightTag = 1 
		head.RightChild = pre   //头节点的右线索指向pre
	}
	return head
}

/*遍历中序线索二叉树
	1)左子树线索化
	2)对空指针线索化：
		①如果p的左孩子为空，则给p加上左线索，将其LTag置为1，让p的左孩子指针指向pre(前驱);
		②如果pre的右孩子为空，则给pre加上右线索，将其RTag置为1，让pre的右孩子指针指向p(后继);
	3)将pre指向刚访问过的结点p，即pre = p;
	4)右子树线索化。

	等于是一个链表的扫描，所以时间复杂度为O(n)
*/
func InOrderTraverse_Thr(bTree *BTree) {
	p := bTree.LeftChild
	for ; p != bTree; {
		//沿左孩子向下
		for ; p.LeftTag == 0; {
			p = p.LeftChild
		}
		fmt.Print(" ", p.Data)  //访问其左子树为空的节点
		//沿右线索
		for ; p.RightTag == 1 && p.RightChild != bTree; { 
			p = p.RightChild
			fmt.Print(" ", p.Data)
		}
		p = p.RightChild

	}
	fmt.Print(" ", p.Data)
}

//求后继的算法
func InOrderNext(bTree *BTree) *BTree {
	if bTree.RightTag == 1 {
		return bTree.RightChild
	} else {
		q := bTree.RightChild
		for ; q.LeftTag == 0; {
			q = q.LeftChild
		}
		return q
	}
}

//求前驱的算法
func InOrderPre(bTree *BTree) *BTree {
	if bTree.LeftTag == 1 {
		return bTree.LeftChild
	} else {
		q := bTree.LeftChild
		for ; q.RightTag == 0 ; {
			q = q.RightChild
		}
		return q
	}
}

