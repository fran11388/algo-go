package main

import "fmt"

func main(){
	nums:=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}


func maxSubArray(nums []int) int {
	//宣告一個dp table用來存放到每一位置為止能形成的 Maximum Subarray
	table:=make([]int,len(nums))
	table[0]=nums[0]
	maxsum:=nums[0]
	for i:=1;i<len(nums);i++{
		prevSum:=table[i-1]
		n:=nums[i]
		table[i]=max(prevSum+n,n)
		if table[i]>maxsum{
			maxsum=table[i]
		}
	}
	return maxsum
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}


