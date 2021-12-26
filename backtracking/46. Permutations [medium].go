package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	return permuteByBacktrack(nums)
}

func permuteByBacktrack(nums []int) [][]int {
	perms := []int{}
	result := [][]int{}
	used := make([]bool, len(nums))
	backtrack(nums, &perms, &result, &used)
	return result
}

func backtrack(nums []int, perms *[]int, result *[][]int, used *[]bool) {
	//goal
	if len(*perms) == len(nums) {
		*result = append(*result, append([]int{}, *perms...))
	}

	u := *used
	for i, _ := range nums {
		if u[i] == true {
			continue
		}
		//make choices
		num := nums[i]
		*perms = append(*perms, num)
		u[i] = true
		backtrack(nums, perms, result, used)
		//undo choice
		p := *perms
		*perms = p[:len(p)-1]
		u[i] = false
	}
}
