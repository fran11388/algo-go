package main

import (
	"fmt"
	"math"
)

func main(){
	n:=123
	fmt.Println(reverse(n))
	n=-123
	fmt.Println(reverse(n))
}
func reverse(x int) int {
	res := 0
	for x!=0{
		pop:=x%10
		x=x/10
	//int32  : -2147483648 ~ 2147483647
		if res>math.MaxInt32/10||(res==math.MaxInt32/10&&pop>7){
			return 0
		}
		if res<math.MinInt32/10||(res==math.MinInt32/10&&pop<(-8)){
			return 0
		}

		res=res*10+pop
	}
	return res
}
