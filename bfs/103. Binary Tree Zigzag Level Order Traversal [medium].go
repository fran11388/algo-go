package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		currentQueueSize := len(queue)
		var collect []int
		for i := 0; i < currentQueueSize; i++ {
			poped := queue[0]
			queue = queue[1:]
			collect = append(collect, poped.Val)
			if poped.Left != nil {
				queue = append(queue, poped.Left)
			}
			if poped.Right != nil {
				queue = append(queue, poped.Right)
			}

		}
		if level%2 != 0 {
			reverse(collect)
		}
		level++
		result = append(result, collect)
	}

	return result
}

func reverse(arr []int) {
	swapIdx := (len(arr) - 1) / 2
	j := len(arr) - 1
	for i := 0; i <= swapIdx; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
}
