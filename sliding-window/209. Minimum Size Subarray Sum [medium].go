package main

import "fmt"

func main(){
	target:=11
	nums:=[]int{1,2,3,4,5}
	fmt.Println(minSubArrayLen(target,nums))
}

func minSubArrayLen(target int, nums []int) int {
	size := len(nums) + 1

	windowSum := 0
	left, right := -1,-1
	for ; left < len(nums) && right < len(nums);  {

		if windowSum >= target { //注意題目要的是大於等於
			s:=getSize(left,right)
			if s<size{
				size=s
			}

			windowSum -=nums[left]
			left++
		}else if windowSum<target{
			right++ //這邊要注意  right移動後就要先判斷有沒有超出範圍, left不用是因為for迴圈有判斷了
			if right>=len(nums){
				break
			}
			if right==0{
				left=0
			}
			windowSum +=nums[right]
		}

	}

	if size==(len(nums) + 1){//代表從來沒有組出target
		return 0
	}
	return size
}

func getSize(left, right int) int {
	return right - left + 1
}
