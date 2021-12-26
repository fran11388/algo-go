package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root==p || root==q{
		return root
	}

	leftSearchResult:=lowestCommonAncestor(root.Left,p,q)
	rightSearchResult:=lowestCommonAncestor(root.Right,p,q)

	if leftSearchResult==nil{
		return rightSearchResult
	}
	if rightSearchResult==nil{
		return leftSearchResult
	}
	return root
}
