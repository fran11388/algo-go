package main

import (
	"fmt"
	"strings"
)

func main(){
	str:="abcdef"
	wordBank:=[]string{"ab","abc","cd","def","abcd","ef","c"}
	fmt.Println(allConstructByTable(str,wordBank))
}

func allConstructByTable(str string, wordBank []string) [][]string {
	table:=make([][][]string,len(str)+1)
	table[0]=[][]string{{}} //保存的是每個str位置前面的所有組合  位置0就代表空字串可有的組合
	for i:=0;i<len(str);i++{
		beginStr:=str[i:]
		//依序比對從某位置起可match到的word
		for _,word:=range wordBank{
			if strings.Index(beginStr,word)==0{
				//若可以match， 就基於這個位置+上match word的長度，
				//把當前位置保存的的所有組合加上match的word  append到 i +len(word)
				combines:=table[i]
				needAppendComnines:=make([][]string,0,len(combines))
				for _,combine:=range combines{
					//可用組合append word
					a:=make([]string,0,len(combine)+1)
					a=append(a,combine...)
					a=append(a,word)
					needAppendComnines=append(needAppendComnines,a)
				}

				//將當前位置先前的所有可用組合append新word並加到
				table[i+len(word)]=append(table[i+len(word)],needAppendComnines...)
			}
		}

	}
	return table[len(str)]

}

func allConstructByRecu(str string, wordBank []string) [][]string {
	if str == "" {
		return [][]string{{}}
	}
	var result [][]string
	for _, word := range wordBank {
		if strings.Index(str, word) == 0 {
			suffix := strings.TrimLeft(str, word)
			suffixWays := allConstructByRecu(suffix, wordBank)
			//var targetWays [][]string
			for _, suffixWay := range suffixWays {
				targetWay := append([]string{word}, suffixWay...)
				result = append(result, targetWay)
			}
		}
	}
	return result
}

func canConstruct(str string, arr []string) bool {
	if str == "" {
		return true
	}
	for _, s := range arr {
		idx := strings.Index(str, s)
		if idx == 0 {
			trimStr := strings.TrimLeft(str, s)
			r := canConstruct(trimStr, arr)
			if r {
				return true
			}

		}
	}

	return false
}