package main

import "fmt"

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 3}))
}

func permuteUnique(nums []int) [][]int {
	result:=[][]int{}
	used:=make([]bool,len(nums))
	backtrack(nums,&[]int{},used,&result)
	return result

}

func backtrack(nums []int ,perms *[]int,used[]bool,result *[][]int){
	if len(*perms)==len(nums){
		newperms:=make([]int,len(*perms))
		copy(newperms,*perms)
		*result=append(*result,newperms)
		return
	}

	choosed:=map[int]bool{} //紀錄當前位置放過的數 確保不重複
	for i,n:=range nums{
		if _,ok:=choosed[n];ok{
			continue
		}

		if used[i]{ //已被使用的不拿
			continue
		}


		choosed[n]=true

		//make choice
		used[i]=true
		*perms=append(*perms,n)
		backtrack(nums,perms,used,result)

		//undo choice
		used[i]=false
		*perms=(*perms)[:len(*perms)-1]
	}
}