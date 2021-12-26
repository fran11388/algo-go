package main

import "fmt"

func main() {
	t := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(t))
}

func maxSubArray(nums []int) int {
	table := make([]int, len(nums))
	table[0] = nums[0]
	maxsum:= nums[0]
	for i := 1; i < len(nums); i++ {
		table[i]=max(table[i-1]+nums[i],nums[i])  //注意這邊要比較的是當前數字 與當前數字+上先前數字的最大
		if table[i]>maxsum{
			maxsum=table[i]
		}
	}
	return maxsum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
