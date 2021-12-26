package main

import (
	"fmt"
)

func main() {
	data := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(data))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//init a 2d array represent ways
	var ways [][]int
	//note :must init 2d array
	ways = make([][]int, len(obstacleGrid))
	for i, v := range obstacleGrid {
		ways[i] = make([]int, len(v))
	}
	ways[0][0] = 1 //在起點只有一種方法   注意這邊要判斷是不是有障礙
	if obstacleGrid[0][0] == 1 {
		ways[0][0] = 0
	}

	//dynamic program implement
	for rowIdx, row := range ways {
		for colIdx, _ := range row {
			currways := ways[rowIdx][colIdx]
			//向右
			if colIdx+1 < len(row) {
				targetHasObstacle := HasObstacle(obstacleGrid, rowIdx, colIdx+1)
				if !targetHasObstacle {
					ways[rowIdx][colIdx+1] = ways[rowIdx][colIdx+1] + currways
				}
			}

			//向下
			if rowIdx+1 < len(ways) {
				targetHasObstacle := HasObstacle(obstacleGrid, rowIdx+1, colIdx)
				if !targetHasObstacle {
					ways[rowIdx+1][colIdx] = ways[rowIdx+1][colIdx] + currways
				}
			}

		}
	}

	fmt.Println(ways)
	rowIdx := len(obstacleGrid) - 1
	colIdx := len(obstacleGrid[0]) - 1
	return ways[rowIdx][colIdx]
}

func HasObstacle(arr [][]int, rowidx, colidx int) bool {
	if arr[rowidx][colidx] == 1 {
		return true
	}
	return false
}
