package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	result := [][]string{}
	solveNQueensHelper(grid, 0, &result)
	return result
}

func solveNQueensHelper(grid [][]int, row int, result *[][]string) {
	if row == len(grid) {
		//代表放滿皇后了
		//把當前的放置方法家道result
		addToResult(grid, result)
		return
	}

	//否則，嘗試每一col
	for col := 0; col < len(grid); col++ {
		//嘗試置放 col i
		grid[row][col] = 1

		//判斷是否有效
		if isValid(grid, row) {
			//有效的話繼續下一row
			solveNQueensHelper(grid, row+1, result)

		}

		//因為要給出所有解 所以不管怎樣都會回逤
		grid[row][col] = 0
	}

}

func addToResult(grid [][]int, result *[][]string) {
	//fmt.Println(grid)

	solution:=[]string{}
	for _,row:=range grid{
		placementStr :=""
		for _,col:=range row{
			if col==0{
				placementStr +="."
			}else{
				placementStr +="Q"
			}
		}
		solution=append(solution, placementStr)
	}
	*result=append(*result,solution)
}

func isValid(grid [][]int, row int) bool {
	//將當前row與先前row依序比對是否在同一col或是對角  若是的話 =false
	currRowColIdx:=getPlacement(grid[row])
	for r := 0; r < row; r++ {
		prevColIdx:=getPlacement(grid[r])
		diff:=abs(prevColIdx-currRowColIdx)
		if prevColIdx==currRowColIdx||diff==row-r{
			return false
		}
	}

	return true
}

func abs(i int)int{
	if i<0{
		return -i
	}
	return i
}

func getPlacement(row []int) int {
	//todo 這邊可以設計資料結構來改善查詢特定row的皇后位置
	for idx, v := range row {
		if v == 1 {
			return idx
		}
	}
	return -1
}
