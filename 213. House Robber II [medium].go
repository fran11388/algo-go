package main

import "fmt"

func main(){
	fmt.Println(rob([]int{5,6}))
}
func rob(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	return getmax(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	if len(nums)==0{
		return 0
	}
	dpTable:=make([]int,len(nums))
	for i,_:=range nums{
		houseA:=0
		if i-2>=0{
			houseA=dpTable[i-2]
		}
		houseB:=0
		if i-1>=0{
			houseB=dpTable[i-1]
		}
		if nums[i]+houseA>houseB{
			//rob
			dpTable[i]=houseA+nums[i]
		}else{
			//not rob
			dpTable[i]=houseB
		}

	}
	return dpTable[len(nums)-1]
}

func getmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
