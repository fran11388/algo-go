package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(numDecodings("226"))
}

func numDecodings(s string) int {
	table:=[]string{"1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26"}
//因為題目要的只是組合結果  所以根本不用管字串應到甚麼英文

	return decodeWays(s,table,make(map[string]int))


	//return 0
}

func decodeWays(s string,table []string,ht map[string]int)int{
	if s==""{
		return 1
	}

	if w,ok:=ht[s];ok{
		return w
	}

	ways:=0
	//loop table check match char

	for _,c:=range table{
		if strings.Index(s,c)==0{

			trimed:=strings.TrimPrefix(s,c)
			d:=decodeWays(trimed,table,ht)
			ht[trimed]=d
			ways+=d
		}
	}
	ht[s]=ways
	return ways
}
