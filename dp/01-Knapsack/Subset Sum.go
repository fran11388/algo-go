package main

import "fmt"

func main() {
	nums :=[]int{1,2,7,1,5}
	s:=10
	nums2 :=[]int{1,3,4,8}
	s2:=6
	fmt.Println(canPartition(nums,s))
	fmt.Println(canPartition(nums2,s2))
}

func canPartition(nums []int, sum int) bool {
	return canPartitionByRecu(nums,0,sum,map[string]bool{})
}

func canPartitionByRecu(nums []int,idx int, remaning int,cache map[string]bool) bool {
	k:=getKey(idx,remaning)
	if v,ok:=cache[k];ok{
		return v
	}


	if remaning==0{
		return true
	}

	if idx>=len(nums){
		return false
	}

	n:=nums[idx]
	a1:=canPartitionByRecu(nums,idx+1,remaning-n,cache)
	if a1{
		cache[k]=true
		return true
	}
	a2:=canPartitionByRecu(nums,idx+1,remaning,cache)
	if a2{
		cache[k]=true
		return true
	}
	cache[k]=false
	return false
}

func getKey(idx int,remaning int)string{
	return fmt.Sprintf("%d_%d",idx,remaning)
}
