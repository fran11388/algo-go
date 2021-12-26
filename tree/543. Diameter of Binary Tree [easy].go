package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	info:=getInfo(root)
	return info.Diameter
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return -1 //這邊回傳-1是為了方便計算，這樣算diameter時可以直接+1，若是null回傳-1可直接抵銷
	}
	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)
	return 1 + getMAx(leftH, rightH)
}

func getMAx(a ...int) int {
	max:=a[0]
	for _,v:=range a{
		if v>max{
			max=v
		}
	}
	return max
}

type NodeInfo struct{
	Height int
	Diameter int
}

func getInfo(root *TreeNode)*NodeInfo{
	info:=&NodeInfo{Height: -1}
	if root==nil{
		return info
	}

	leftInfo:=getInfo(root.Left)
	rightInfo:=getInfo(root.Right)
	info.Height=1+getMAx(leftInfo.Height,rightInfo.Height)
	diameter:=2+leftInfo.Height+rightInfo.Height

	info.Diameter=getMAx(diameter,leftInfo.Diameter,rightInfo.Diameter)
	return info
}
