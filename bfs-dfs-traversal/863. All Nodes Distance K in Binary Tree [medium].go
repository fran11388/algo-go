package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	nodeParentMap := getNodeParentMap(root)

	level := 0
	queue := []*TreeNode{target}
	result := []int{}
	seen := map[int]bool{
		target.Val: true,
	}
	for len(queue) > 0 {
		size := len(queue)
		if level == k {
			for i := 0; i < size; i++ {
				poped := queue[0]
				queue = queue[1:]
				result = append(result, poped.Val)

			}
			break
		} else {
			for i := 0; i < size; i++ {
				poped := queue[0]
				queue = queue[1:]
				nodeVal := poped.Val
				//traverse neighbor
				if parent, ok := nodeParentMap[nodeVal]; ok {
					if _, ok := seen[parent.Val]; !ok { //bfs parent
						seen[parent.Val] = true
						queue = append(queue, parent)

					}
				}
				if poped.Left != nil {
					leftNode := poped.Left
					if _, ok := seen[leftNode.Val]; !ok { //bfs left node
						seen[leftNode.Val] = true
						queue = append(queue, leftNode)

					}
				}

				if poped.Right != nil {
					rightNode := poped.Right
					if _, ok := seen[rightNode.Val]; !ok { //bfs right node
						seen[rightNode.Val] = true
						queue = append(queue, rightNode)

					}
				}
			}
			level++
		}

	}
	return result
}

func getNodeParentMap(root *TreeNode) map[int]*TreeNode {
	result := map[int]*TreeNode{}
	dfs(root, result)
	return result
}

func dfs(root *TreeNode, m map[int]*TreeNode) {
	if root != nil {
		if root.Left != nil {
			m[root.Left.Val] = root
		}
		if root.Right != nil {
			m[root.Right.Val] = root
		}

		dfs(root.Left, m)
		dfs(root.Right, m)
	}

}
