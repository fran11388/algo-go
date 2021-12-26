package main

import "fmt"

func main(){
	candidates := []int{2,3,6,7}
	target := 7
	fmt.Println(combinationSum(candidates,target))
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	dfs(candidates, target, 0, &[]int{}, 0, &result)
	return result
}

func dfs(candidates []int, target int, candidatesIdx int, curr *[]int, total int, result *[][]int) {
	if target == total {
		newResult := make([]int, len(*curr))
		copy(newResult, *curr)
		*result=append(*result,newResult)
		return
	}

	if candidatesIdx >= len(candidates) || total > target {
		return
	}

	//choose element
	e := candidates[candidatesIdx]
	*curr = append(*curr, e)
	total += e
	dfs(candidates, target, candidatesIdx, curr, total, result)

	//not choose (backtracking undo choice)
	*curr = (*curr)[:len(*curr)-1]
	total -= e
	dfs(candidates, target, candidatesIdx+1, curr, total, result) //這邊shift idx是為了後面不再考慮該元素
}
