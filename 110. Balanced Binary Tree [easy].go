package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getHeight(root *TreeNode)int{
	if root==nil{
		return -1
	}
	return 1+max(getHeight(root.Left),getHeight(root.Right))
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func abs(i int)int{
	if i<0{
		return -i
	}
	return i
}


func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}


	//左右子樹高度相差不超過1
	leftSubHeight:=getHeight(root.Left)
	rightSubHeight:=getHeight(root.Right)
	if abs(leftSubHeight-rightSubHeight)>1{
		return false
	}

	//左右都是平衡
	if !isBalanced(root.Left) || !isBalanced(root.Right){
		return false
	}


	return true
}
