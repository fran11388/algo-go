package main

import "fmt"

func main(){

	fmt.Println(12<<3)
}
func singleNumber(nums []int) int {
	r:=nums[0]
	for i:=1;i<len(nums);i++{
		r=r^nums[i]
	}

	return r
}
