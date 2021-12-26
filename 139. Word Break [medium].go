package main

import (
	"fmt"
	"strings"
)

func main(){
	t:="applepenapple"
	wordDict:=[]string{"apple","pen"}
	fmt.Println(wordBreakHelper(t,wordDict,make(map[string]bool)))
}

func wordBreakHelper(s string, wordDict []string,memo map[string]bool) bool {
	if v,ok:=memo[s];ok{
		return v
	}

	if s==""{
		memo[s]=true
		return true
	}
	canBreak:=false
	//loop determine can match in first position
	for _,word:=range wordDict{
		if strings.Index(s,word)==0{
			trimedStr:=strings.TrimPrefix(s,word)
			trimedStrCanBreak:= wordBreakHelper(trimedStr,wordDict,memo)
			if trimedStrCanBreak{
				canBreak=true
				break
			}
		}
	}
	memo[s]=canBreak
	return canBreak
}