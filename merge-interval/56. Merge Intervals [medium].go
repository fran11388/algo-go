package main

import (
	"fmt"
	"sort"
)

func main(){
	d:=[][]int{{1,3},{2,6},{8,10},{15,18}}
	fmt.Println(merge(d))

}
func merge(intervals [][]int) [][]int {
	if len(intervals)<2{
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0]<intervals[j][0]
	})

	result:=[][]int{}
	currInterval:=[]int{intervals[0][0],intervals[0][1]}
	for i:=1;i<len(intervals);i++{
		interval:=intervals[i]
		if isInInterval(currInterval,interval){
			mergeInterval(currInterval,interval)
		}else{
			//currInterval add to result
			result=append(result,[]int{currInterval[0],currInterval[1]})

			//reset interval
			currInterval[0]=interval[0]
			currInterval[1]=interval[1]
		}
	}

	//記得最後還要把currInterval家道結果中
	result=append(result,[]int{currInterval[0],currInterval[1]})

	return result
}

func isInInterval(currInterval[]int,interval []int)bool{
	min:=currInterval[0]
	max:=currInterval[1]

	v:=interval[0]
	if v<=max && v>=min{
		return true
	}
	return false
}

func mergeInterval(currInterval []int,interval []int){
	currInterval[1]=getMax(currInterval[1],interval[1])
}

func getMax(a,b int)int{
	if a>b{
		return a
	}
	return b
}


