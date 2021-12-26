package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}

func permuteUnique(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	result := [][]int{}

	//蝶帶每個nums中的數字  但重複的不處理
	seen := make(map[int]bool)
	for i, _ := range nums {
		takenNum := nums[i]
		if _, ok := seen[takenNum]; ok {
			continue
		}
		seen[takenNum] = true
		leftPart := append([]int{}, nums[0:i]...)
		perms := permuteUnique(append(leftPart, nums[i+1:]...))
		for _, perm := range perms {
			newPerm := []int{takenNum}
			newPerm = append(newPerm, perm...)
			result = append(result, newPerm)
		}
	}
	return result

}
