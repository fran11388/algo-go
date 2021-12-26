package main

import "fmt"

func main(){
	nums := []int{}
	for i := 1; i <= 20; i++ {
		nums = append(nums, i)
	}
	fmt.Println(combinationIterative(nums))
	//fmt.Println(combinationRecursive(nums))
}


//input:
//[1,2,3]
//result:
//[]
//[] [1]
//[] [1] [2],[1,2]
//[] [1] [2],[1,2]  [3],[1,3],[2,3],[1,2,3]
func combinationIterative(nums[]int)[][]int{
	//每一元素都有拿與不拿的選擇
	result:=[][]int{{}}
	for _,n:=range nums{
		var newResult [][]int
		//不拿的結果append到新result
		newResult=append(newResult,result...)
		//加入拿的結果(蝶帶先前的結果並加入新元素)
		for _,set:=range result{
			//避免shallow copy
			newSet:=append([]int{},set...)
			newSet=append(newSet,n)
			newResult=append(newResult,newSet)
		}
		result=newResult
	}

	return result
}

//input:
//[1,2,3]
//result:
//[] [1] [2],[1,2]  [3],[1,3],[2,3],[1,2,3]
func combinationRecursive(nums []int)[][]int{
	//每一次地回呼叫聲成一個新的結果及:拿與不拿
	prevResult:=[][]int{{}}//注意這邊的初始結果包含一個空集合
	return combinationRecursiveHelp(prevResult,nums)
}

func combinationRecursiveHelp(prevResult [][]int,nums []int)[][]int{
	if len(nums)==0{
		return prevResult
	}


	var newResult [][]int
	//將不拿的結果append到新結果及
	//[], [], []...
	newResult=append([][]int{},prevResult...)

	//將拿的結果append到新的結果及(迴圈蝶帶先前的結果再append)
	//取第一個元素作為要append的新元素
	takenNum :=nums[0]
	for _,set:=range prevResult{
		//避免shallow copy
		newSet:=append([]int{},set...)
		newSet=append(newSet, takenNum)
		newResult=append(newResult,newSet)
	}

	//將新生成的結果地回呼叫(這邊傳入的num去掉第一個取走的元素)
	return combinationRecursiveHelp(newResult,nums[1:])
}

