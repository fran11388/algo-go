package main

import "fmt"

func main(){
fmt.Println(combine(4,2))
}


func combine(n int, k int) [][]int {
	result:=[][]int{}
	backtrack(n,k,&[]int{},make(map[int]bool),&result)
	return result
}

func backtrack(n int, k int,currNums *[]int,used map[int]bool,result *[][]int){
	//goal
	if len(*currNums)==k{
		*result=append(*result,append([]int{},*currNums...))
		return
	}

	i:=1
	if len(*currNums)>0{
		//這個步驟很重要，choice的選擇必須要在當前num之後
		// 實作時一定要先把圖畫出來，絲路才會清楚
		c:=*currNums
		i=c[len(c)-1]
	}

	for ;i<=n;i++{

		if _,ok:=used[i];ok{
			continue
		}

		//make choice
		used[i]=true//mark used
		*currNums=append(*currNums,i)
		backtrack(n,k,currNums,used,result)
		//undo choice
		c:=*currNums
		*currNums=c[0:len(c)-1]
		delete(used,i)
	}
}