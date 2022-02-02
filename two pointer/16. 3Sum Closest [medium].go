package main

import (
	"math"
	"sort"
)

func threeSumClosestByTwoPointer(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]-nums[j] < 0
	})

	diff:=math.MaxInt64 //紀錄右偏移
	for i:=0;i<len(nums)-2;i++{
		n1:=nums[i]
		left:=i+1
		right:=len(nums)-1
		for left<right{
			n2:=nums[left]
			n3:=nums[right]
			sum:=n1+n2+n3
			if sum-target==0{
				return sum
			}else if sum >target {
				right--
				if abs(sum-target)<abs(diff){
					diff=sum-target
				}
			}else{
				left++
				if abs(sum-target)<abs(diff){
					diff=sum-target
				}
			}

		}
	}
	return diff+target
}
func abs(a int)int{
	if a<0{
		return -a
	}
	return a
}