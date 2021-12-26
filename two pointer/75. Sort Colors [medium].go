package main

import "fmt"

func main() {

	arr := []int{2,0,2,1,1,0}
	sortColors(arr)
	fmt.Println()
}

func sortColors(nums []int) {
	left := -1
	right := len(nums)

	for i := 0; i < len(nums); i++ {
		if i >= right {
			break
		}

		//這邊必須要持續將當前位置的數字換到正確位置
		for true {
			if right==i || left==i{
				break
			}
			n:=nums[i]
			if n==0{
				swap(nums,left+1,i)
				left++
				continue
			}else if n==2{
				swap(nums,right-1,i)
				right--
				continue
			}
			//mean n==1
			break
		}

	}
}
func swap(nums []int, idx1, idx2 int) {
	nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
}
