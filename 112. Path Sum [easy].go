package main

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

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumHelp(root, 0, targetSum)
}

func hasPathSumHelp(root *TreeNode, currsum int, targetSum int) bool {
	if root==nil{
		return false
	}
	isCurrsumEqualTargetsum := currsum+root.Val == targetSum


	if root.Left == nil && root.Right == nil && isCurrsumEqualTargetsum {
		return true
	}

	l := hasPathSumHelp(root.Left, currsum+root.Val, targetSum)
	if l {
		return true
	}
	r := hasPathSumHelp(root.Right, currsum+root.Val, targetSum)
	if r {
		return true
	}
	return false

}
