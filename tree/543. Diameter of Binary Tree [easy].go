package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//宣告一變數追蹤最長的diameter
//對每個節點考慮左右的longestPath，並將其連起來，看看是否大於diameter
//因為在計算longestPath時，就會需要考慮左右，所以這時就可以同時更新diameter
func diameterOfBinaryTree(root *TreeNode) int {
	diameter:=0
	longestPath(root,&diameter)
	return diameter
}

func longestPath(root *TreeNode,diameter *int ) int {
	if root==nil{
		return -1  //空節點設為-1，當一節點左右子樹都空 +1後剛好會是0
	}

	left:=longestPath(root.Left,diameter)
	right:=longestPath(root.Right,diameter)
	m:=getMax(left,right)

	//因為左右子樹的longestPath都有，可以同時把他們連起來看看有沒有大於diameter
	d:=2+left+right
	if d>*diameter{
		*diameter=d
	}

	return m+1//增加高度1
}

func getMax(a,b int)int{
	if a>b{
		return a
	}
	return b
}
