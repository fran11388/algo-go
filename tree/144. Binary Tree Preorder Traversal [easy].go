package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorderTraversalHelper(root, &result)
	return result
}

func preorderTraversalHelper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	v := root.Val
	*result = append(*result, v)

	if root.Left != nil {
		preorderTraversalHelper(root.Left, result)
	}

	if root.Right != nil {
		preorderTraversalHelper(root.Right, result)
	}
}
