package main

//先計算每個位置從左, 從右的prefix sum，對於每一當前位置等於兩邊的相乘
//最後直接複用right 的prefix sum
func productExceptSelfByPrefixSum(nums []int) []int {
	prifixSumFromLeft:=make(  []int,len(nums) )
	prifixSumFromLeft[0]=1 //設1 因為最後該位置會乘上右prefix sum
	for i:=1;i<len(nums);i++{
		prifixSumFromLeft[i] = nums[i-1]*prifixSumFromLeft[i-1]
	}

	prifixSumFromRight:=make(  []int,len(nums) )
	prifixSumFromRight[len(nums)-1]=1
	for i:=len(nums)-2;i>=0;i--{
		prifixSumFromRight[i] = nums[i+1]*prifixSumFromRight[i+1]
	}

	for i:=0;i<len(nums);i++{
		prifixSumFromRight[i]*= prifixSumFromLeft[i]
	}
	return prifixSumFromRight
}


//基於prefix sum 但優化空間(先從左開始leftPrefixSum直接放output array，
//再從右邊向左trace 逐一記錄一個rightPrefixSum，
//同時做運算放回output)
func productExceptSelfByPrefixSumOptimizeSpace(nums []int) []int {
	output:=make([]int ,len(nums))
	output[0]=1
	for i:=1;i<len(nums);i++{
		output[i]=output[i-1]*nums[i-1]
	}

	rightPrefixSum:=1
	for i:=len(nums)-2;i>=0;i--{
		rightPrefixSum*=nums[i+1]
		output[i]=output[i]*  rightPrefixSum //相當於leftPrefixSum* rightPrefixSum
	}
	return output
}
