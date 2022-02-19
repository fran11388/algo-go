package main

func subsets(nums []int) [][]int {
	result:=[][]int{}
	beginIdx:=0
	currNums:=[]int{}
	backtrack(nums,beginIdx,&currNums,&result)
	return result
}

func backtrack(nums []int, beginIdx int, currNums *[]int, result *[][]int){
	if beginIdx==len(nums){
		newnums:=make([]int,len(*currNums))
		copy(newnums,*currNums)
		*result=append(*result,newnums)
		return
	}

	//對於每個element, 可以考慮選與不選
	n:=nums[beginIdx]
	*currNums=append(*currNums,n)
	backtrack(nums,beginIdx+1,currNums,result) //選

	*currNums=(*currNums)[:len(*currNums)-1]
	backtrack(nums,beginIdx+1,currNums,result) //不選
}
