package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
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
