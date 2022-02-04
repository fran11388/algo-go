package main

import (
	"fmt"
	"sort"
)

func main() {
	//data:=[]int{0,0,2,2,0,0,2,2}
	//t:=4
	data := []int{1, 0, -1, 0, -2, 2}
	t := 0
	fmt.Println(fourSum(data, t))
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	if len(nums) < 4 {
		return result
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	result = kSum(nums, 4, 0, target)
	return result
}

//對於4(k size)sum 可以拆解成迭代nums 0~n-(k-1)的元素 +上遞迴 size-1的 ksum的解
func kSum(nums []int, size int, startIdx int, target int) [][]int {
	if size == 2 {
		return twoSum(nums, startIdx, target)
	}
	result := [][]int{}
	exist := map[int]bool{}
	for i := startIdx; i <= len(nums)-size; i++ {
		ele := nums[i]
		if _, ok := exist[ele]; ok { //跳過重複element
			continue
		}
		exist[ele] = true
		subResults := kSum(nums, size-1, i+1, target-ele) //注意!! 這邊遞迴傳入的startIdx參數要是當前i+1，一開始弄錯了答案一直錯
		for _, subResult := range subResults {
			newresult := []int{ele}
			newresult = append(newresult, subResult...)
			result = append(result, newresult)
		}
	}
	return result
}

func twoSum(nums []int, startIdx int, target int) [][]int {
	left := startIdx
	right := len(nums) - 1
	result := [][]int{}
	exist := map[int]bool{}
	for left < right {
		n1 := nums[left]
		n2 := nums[right]
		sum := n1 + n2
		if sum == target {
			if _, ok := exist[n1]; !ok { //去除重複
				ans := []int{n1, n2}
				result = append(result, ans)
				exist[n1] = true
				exist[n2] = true
			}

			left++
			right--
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return result
}
