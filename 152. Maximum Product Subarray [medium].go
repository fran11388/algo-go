package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	//fmt.Println(maxProduct([]int{-1,-2,-3}))
	//fmt.Println(maxProduct([]int{-2,0,-1}))
}
func maxProduct(nums []int) int {
	minDPTable := make([]int, len(nums))
	maxDPTable := make([]int, len(nums))

	if len(nums) < 2 {
		return nums[0]
	}
	minDPTable[0] = nums[0]
	maxDPTable[0] = nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		prevMax := maxDPTable[i-1]
		prevMin := minDPTable[i-1]
		//2, 3, -2, 4
		minDPTable[i] = getMin(n, prevMax*n, prevMin*n)
		maxDPTable[i] = getMax(n, prevMax*n, prevMin*n)

		if maxDPTable[i]>max{
			max=maxDPTable[i]
		}
	}
	//注意 這邊是要回傳dp table中最大的，因為最大的乘積可能發生在先前subarray
	return max
}
func getMin(nums ...int) int {
	min := math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func getMax(nums ...int) int {
	max := math.MinInt32
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
