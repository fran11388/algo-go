package main

import "fmt"

func main(){
	nums:=[]int{1,-1,0}
	k:=0
	fmt.Println(subarraySum(nums,k))
}

func subarraySum(nums []int, k int) int {
	return subarraySumPrefixSumHt(nums,k)
}


// time:O(n)
// space:O(n)
func subarraySumPrefixSumHt(nums []int, k int) int {
	sum:=0
	result:=0
	ht:=make(map[int]int)
	ht[0]=1
	for i:=0;i<len(nums);i++{
		sum+=nums[i]

		//耖 這邊的城市要先執行
		if _,ok:=ht[sum-k];ok{
			result+=ht[sum-k]
		}

		if _,ok:=ht[sum];ok{
			ht[sum]++
		}else{
			ht[sum]=1
		}


	}

	return result
}



// time:O(n^2)
// space:O(1)
func subarraySumBruteForce(nums []int, k int) int {

	result:=0

	for i:=0;i<len(nums);i++{
		sum:=0
		for j:=i;j<len(nums);j++{
			sum+=nums[j]
			if sum==k{
				result++
			}
		}
	}

	return result
}
