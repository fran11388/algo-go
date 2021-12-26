package main

import "fmt"

func main(){
	fmt.Println(rob([]int{2,1,1,2}))
}
func rob(nums []int) int {
	dpTable:=make([]int,len(nums))
	for i,_:=range nums{
		preprePrice:=0
		if i-2>=0{
			preprePrice=dpTable[i-2]
		}

		prevPrice:=0
		if i-1>=0{
			prevPrice=dpTable[i-1]
		}

		if nums[i]+preprePrice>prevPrice{
			dpTable[i]=nums[i]+preprePrice
		}else{
			dpTable[i]=prevPrice
		}
	}
	return dpTable[len(nums)-1]
}
