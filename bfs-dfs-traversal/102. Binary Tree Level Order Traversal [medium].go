package main

//思路:使用bfs level by level traverse, currLevelCount紀錄當前level要被pop up的數量,
//nextLevelCount紀錄訪問當前level時push 進queue的數量, 最後做為下一level的數量
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	currLevelCount := 1
	nextLevelCount := 0

	for len(queue) > 0 {
		row := []int{}

		for i := 0; i < currLevelCount; i++ {
			pop := queue[0]
			queue = queue[1:]
			row = append(row, pop.Val)
			if pop.Left != nil {
				nextLevelCount++
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				nextLevelCount++
				queue = append(queue, pop.Right)
			}
		}
		result = append(result, row)
		currLevelCount = nextLevelCount
		nextLevelCount = 0
	}
	return result
}
