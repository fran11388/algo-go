package main

import "fmt"

func main() {
	g := [][]int{
		{1, 0, 1, 0},
		{1, 0, 0, 0},
		{1, 0, 1, 1},
		{0, 0, 1, 1}}
	fmt.Println(minimumIsland(g))
}
func minimumIsland(grid [][]int) int {
	visit := make(map[string]bool)
	min := len(grid) * len(grid[0])
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			s := exploreSize(grid, r, c, visit)
			if s < min&&s>0 { //注意這邊要判斷s需大於0  因為這邊是每個element都會去explore  不滿足的會回0
				min = s
			}
		}
	}
	return min
}

func exploreSize(grid [][]int, r, c int, visit map[string]bool) int {
	rowInBounds := r >= 0 && r < len(grid)
	colInBounds := c >= 0 && c < len(grid[0])
	if (!rowInBounds) || (!colInBounds) {
		return 0
	}
	if grid[r][c] == 0 {
		return 0
	}
	pos := fmt.Sprintf("%d,%d", r, c)
	if _, ok := visit[pos]; ok {
		return 0
	}

	visit[pos] = true
	count := 1
	count += exploreSize(grid, r+1, c, visit)
	count += exploreSize(grid, r-1, c, visit)
	count += exploreSize(grid, r, c+1, visit)
	count += exploreSize(grid, r, c-1, visit)
	return count

}
