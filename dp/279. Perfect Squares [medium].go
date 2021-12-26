package main

import "fmt"

func main(){
	n:=100
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	table:=make([]int,n+1)
	for i:=0;i<len(table);i++{
		table[i]=n
	}
	table[0]=0
	for i:=0;i<n;i++{ //耖  數值會從0開始啊
		for j:=1;(j*j) <=(n-i);j++{
			t:=i+(j*j)
			table[t]=min(table[t],table[i]+1)

		}

	}
	return table[n]
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
