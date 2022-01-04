package main

import "fmt"

func main() {
	nums := []int{1, 5, 11, 5}
	//nums:=[]int{1,2,3,5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {


	return canPartitionByDP(nums)
	//return canPartitionByBruteForce(nums)
}

func canPartitionByDP(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	list:=make([]int,0,target)//just prevent relocate
	seen:=map[int]bool{}
	list=append(list,0)
	seen[0]=true
	for _,n:=range nums{

		listLen:=len(list)
		for i:=0;i<listLen;i++{
			listN:=list[i]
			newN:=listN+n //計算新的可加總出的數
			if newN==target{ //可組出目標數直接return true
				return true
			}

			if _,ok:=seen[newN];!ok{
				seen[newN]=true
				list=append(list,newN)
			}
		}



	}

	return false
}

//Time Limit Exceeded
func canPartitionByBruteForce(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	currSum := 0
	return canPartitionHelper(nums, 0, currSum, target)
}

func canPartitionHelper(nums []int, idx int, currSum int, target int) bool {
	if currSum == target {
		return true
	}
	if idx >= len(nums) {
		return false
	}
	n := nums[idx]

	a := canPartitionHelper(nums, idx+1, currSum, target)
	if a {
		return true
	}
	b := canPartitionHelper(nums, idx+1, currSum+n, target)
	if b {
		return true
	}
	return false
}
