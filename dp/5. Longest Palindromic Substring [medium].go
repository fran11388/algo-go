package main

import "fmt"

func main(){
	s:="aaaaa"
	//s:="cbbd"
	fmt.Println(longestPalindrome(s))
}
func longestPalindrome(s string) string {
	res := getCharByIdx(s,0)
	dpTable :=make([][]bool,len(s))
	for i:=0;i<len(dpTable);i++{
		dpTable[i]=make([]bool,len(s))
	}

	//init dptable value
	for i:=0;i<len(dpTable);i++{
		dpTable[i][i]=true
		if (i+1)<len(dpTable){
			if getCharByIdx(s,i)==getCharByIdx(s,i+1){
				dpTable[i][i+1]=true
			}
		}
	}

	for i:=len(s)-1;i>=0;i--{  //注意這邊順序，要從右開始，確保是bottom up
		for j:=i+1;j<len(s);j++{
			if getCharByIdx(s,i)==getCharByIdx(s,j){
				if (i+1)<(j-1){
					if dpTable[i+1][j-1]{
						dpTable[i][j]=true
						if getLeng(i,j)>len(res){
							res=s[i:j+1]
						}
					}
				}else{
					dpTable[i][j]=true //若裡面沒有的確認  直接設true
					if getLeng(i,j)>len(res){
						res=s[i:j+1]
					}
				}
			}
		}
	}




	return res
}
func getCharByIdx(s string,idx int)string{
	return string(s[idx])
}

func getLeng(idx1,idx2 int)int{
	return idx2-idx1+1
}

