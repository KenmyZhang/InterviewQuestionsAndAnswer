package main 

type BTree struct {
	Data int
	Right *BTree
	Left *BTree
}


//时间复杂度为O（N^2）
func IsBalanceBTree(bTree *BTree) bool {
	if bTree == nil {
		return false
	}

	leftHeight = Depth(bTree.Left)
	rightHeight = Depth(bTree.Right)
	return math.Abs(Right - Left) < 2 && IsBalanceBTree(root.Left) && IsBalanceBTree(root.Right)
}

//时间复杂度为O（N）的
func IsBalanceBTree2(bTree *BTree, depth *int) bool {
	if bTree == nil {
		*depth = 0
		return true
	}
	leftHeight := 0
	if IsBalanceBTree2(bTree.Left, &leftHeight) == false {
		return false
	}

	rightHeight := 0 
	if IsBalanceBTree2(bTree.Right, &rightHeight) == false {
		return false
	}

	*depth = rightHeight > leftHeight ? rightHeight + 1: leftHeight +1
	return math.Abs(leftHeight - rightHeight) < 2

}

func main() {
	

}