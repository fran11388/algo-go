package main

import "fmt"

func main(){
fmt.Println(combine(4,2))
}

func combine(n int, k int) [][]int {
	result:=[][]int{}
	numsIdx:=1
	currNums:=[]int{}
	backtrack(n,k,numsIdx,&currNums,&result)
	return result
}

func backtrack(n int, k int, numsIdx int,currNums *[]int, result *[][]int ){
	if len(*currNums)==k{
		newnums:=make([]int,k)
		copy(newnums,*currNums)
		*result=append(*result,newnums)
		return
	}

	for i:=numsIdx;i<=n;i++{
		*currNums=append(*currNums,i)
		backtrack(n,k,i+1,currNums,result ) //i+1代表下次從當前選的element之後考慮
		*currNums=(*currNums)[:len(*currNums)-1]//undo choice
	}
}

