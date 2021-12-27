package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	//return permuteByBacktracking(nums)
	return permuteByForLoop(nums)
}

func permuteByBacktracking(nums []int) [][]int {
	perms := []int{}
	used := make([]bool, len(nums))
	result := [][]int{}
	backtrack(nums, &perms, used, &result)
	return result
}

func backtrack(nums []int, perms *[]int, used []bool, result *[][]int) {
	if len(*perms) == len(nums) {
		res := make([]int, len(*perms))
		copy(res, *perms)
		*result = append(*result, res)
		return
	}

	for i, n := range nums {
		if used[i] {
			continue
		}

		//make choice
		used[i] = true
		*perms = append(*perms, n)
		backtrack(nums, perms, used, result)

		//undo choice
		used[i] = false
		*perms = (*perms)[:len(*perms)-1] //remove last element
	}
}

func permuteByForLoop(nums []int) [][]int {
	result := [][]int{{}}
	for _, n := range nums {
		result = placeNumber(result, n)
	}
	return result
}

func placeNumber(arr [][]int, n int) [][]int {
	result := [][]int{}

	for _, list := range arr {
		tmpResult := make([]int, 0, len(list)+1)
		tmpResult = append(tmpResult, n)
		tmpResult = append(tmpResult, list...)

		//add to result
		res := make([]int, len(tmpResult))
		copy(res, tmpResult)
		result = append(result, res)
		for i := 0; i < len(tmpResult)-1; i++ { //移動n-1次(第一個元素依序移動到尾)
			tmpResult[i], tmpResult[i+1] = tmpResult[i+1], tmpResult[i]
			res = make([]int, len(tmpResult))
			copy(res, tmpResult)
			result = append(result, res)
		}

	}
	return result
}
