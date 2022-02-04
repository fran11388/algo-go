package main

func numIslands(grid [][]byte) int {
	islandCount:=0
	visit:=initVisit(grid)

	for r:=0;r<len(grid);r++{
		for c:=0;c<len(grid[r]);c++{

			if !visit[r][c] && grid[r][c]=='1'{
				islandCount++
				dfs(grid,visit,r,c)
			}
		}
	}
	return islandCount
}

func initVisit(grid [][]byte)[][]bool{
	r:=make([][]bool,len(grid))
	for i,_:= range grid{
		r[i]=make([]bool,len(grid[i]))
	}
	return r

}

func dfs(grid [][]byte ,visit [][]bool,r,c int){
	if !isValid(grid, r,c){
		return
	}
	if visit[r][c]{
		return
	}
	visit[r][c]=true
	dfs(grid,visit,r-1,c)// top
	dfs(grid,visit,r+1,c)// down
	dfs(grid,visit,r,c-1)// left
	dfs(grid,visit,r,c+1)// right
}

func isValid(grid [][]byte,r,c int) bool{
	if !( r<len(grid)&& r>=0){
		return false
	}

	if !(c<len(grid[r]) && c>=0){
		return false
	}

	if !(grid[r][c]=='1'){
		return false
	}

	return true

}



