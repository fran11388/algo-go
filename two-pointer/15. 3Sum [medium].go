package main

import (
	"fmt"
	"sort"
)

func main() {
	d := []int{34,55,79,28,46,33,2,48,31,-3,84,71,52,-3,93,15,21,-43,57,-6,86,56,94,74,83,-14,28,-66,46,-49,62,-11,43,65,77,12,47,61,26,1,13,29,55,-82,76,26,15,-29,36,-29,10,-70,69,17,49}
	fmt.Println(threeSumByTwoPointer(d))
}

func threeSumByTwoPointer(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	//sort input
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	// 2   0   0   -2   -2
	// i   lo            hi
	for i := 0; i < len(nums)-2; i++ {
		lo := i + 1
		hi := len(nums) - 1
		//第1個元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 { //代表後面的加總不可能等於0 直接跳過
			continue
		}
		target := 0 - nums[i]
		for lo < hi {
			n1 := nums[lo]
			n2 := nums[hi]

			//第2個元素去重(只需判斷lo指標的元素有無重複，因為target是固定的，只要lo指的元素不重複，就可以確保另一個也不重複)
			if lo>i+1&&nums[lo]==nums[lo-1]{
				//注意!! 這邊不可直接break 要continue探索其他可能
				lo++
				continue
			}

			if n1+n2 == target {
				ans := []int{nums[i], n1, n2}
				result = append(result, ans)
				//移動左右指標
				lo++
				hi--
			} else if n1+n2 > target { //要讓總和變小，移動右指標
				hi--
			} else { //要讓總和變大，移動左指標
				lo++
			}
		}

	}
	return result
}

//wrong answer(will have duplicate
func threeSumByTraverseAll(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	for i := 0; i < len(nums)-2; i++ {
		target := 0 - nums[i]
		for j := i + 1; j < len(nums)-1; j++ {
			n1 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				n2 := nums[k]
				if n1+n2 == target {
					ans := []int{nums[i], nums[j], nums[k]}
					result = append(result, ans)
				}
			}
		}

	}

	return result
}
