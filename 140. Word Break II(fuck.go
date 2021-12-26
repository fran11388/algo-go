package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "pineapplepenapple"
	wordDict := []string{"apple","pen","applepen","pine","pineapple"}

	//s := "catsanddog"
	//wordDict := []string{"cat","cats","and","sand","dog"}

	fmt.Println(wordBreak(s, wordDict),len(wordBreak(s, wordDict)))
}

func wordBreak(s string, wordDict []string) []string {
	if s == "" {
		//var fuck []string
		return make([]string,0)
	}
	var result []string
	for _, word := range wordDict {
		if strings.Index(s, word) == 0 { //if match
			trimedStr := strings.TrimPrefix(s, word)
			trimedStrWordBreak := wordBreak(trimedStr, wordDict)

			if trimedStrWordBreak!=nil{
				if len(trimedStrWordBreak)==0{
					result=append(result,word)
				}else{
					resultStr := word
					for _, token := range trimedStrWordBreak {
						resultStr += fmt.Sprintf(" %s", token)
					}
					result = append(result, resultStr)
				}
			}


		}
	}

	return result
}
