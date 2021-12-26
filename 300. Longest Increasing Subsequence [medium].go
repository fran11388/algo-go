package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))
}

func lengthOfLIS(nums []int) int {
	leng := make([]int, len(nums))
	leng[0]=1
	maxLeng:=1

	for currIdx :=1; currIdx <len(nums); currIdx++{
		currMaxLeng :=1
		for prevIdx:=0;prevIdx< currIdx;prevIdx++{

			if nums[currIdx]>nums[prevIdx]{ //注意這邊要是大於  因為是遞增
				if leng[prevIdx]+1> currMaxLeng{
					currMaxLeng=leng[prevIdx]+1
				}
			}
		}
		leng[currIdx]=currMaxLeng
		if currMaxLeng>maxLeng{
			maxLeng=currMaxLeng
		}
	}
	//注意 這邊不是直接返回最後一個  而是返回陣列中最大的
	//因為最長子序列的最後一個元素，不一定剛好是輸入元素的最後一個元素
	return maxLeng
}
