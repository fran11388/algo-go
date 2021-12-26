package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}

func findTargetSumWays(nums []int, target int) int {
	return findTargetSumWaysHelper(nums, target, 0, 0)
}

func findTargetSumWaysHelper(nums []int, target int, currIdx int, total int) int {
	if currIdx == len(nums) && total == target { //注意這邊的index判斷  所有數都做過加減以後，最後的IDX會超出陣列
		//所以這邊的INDEX才會是等於陣列的長度
		return 1
	}
	if currIdx == len(nums) { //注意這邊的index判斷
		return 0
	}

	currNum := nums[currIdx]
	return findTargetSumWaysHelper(nums, target, currIdx+1, total+currNum) + findTargetSumWaysHelper(nums, target, currIdx+1, total-currNum)
}
